package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-chi/chi/v5"
)

type apiHandler struct {
	db *sql.DB
}

func RegisterApiHandler(router *chi.Mux, db *entsql.Driver) {
	handler := &apiHandler{
		db: db.DB(),
	}

	router.Route("/", func(apiRouter chi.Router) {
		apiRouter.Get("/query", handler.query)
	})

}

func (handler apiHandler) query(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("sql")
	result, err := handler.db.Query(query)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	columnTypes, err := result.ColumnTypes()
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	count := len(columnTypes)
	finalRows := []interface{}{}

	for result.Next() {

		scanArgs := make([]interface{}, count)

		for i, v := range columnTypes {

			switch v.DatabaseTypeName() {
			case "VARCHAR", "TEXT", "UUID", "TIMESTAMP":
				scanArgs[i] = new(sql.NullString)
				break
			case "BOOL":
				scanArgs[i] = new(sql.NullBool)
				break
			case "INT4":
				scanArgs[i] = new(sql.NullInt64)
				break
			default:
				scanArgs[i] = new(sql.NullString)
			}
		}

		err := result.Scan(scanArgs...)

		if err != nil {
			w.Write([]byte(err.Error()))
		}

		masterData := map[string]interface{}{}

		for i, v := range columnTypes {

			if z, ok := (scanArgs[i]).(*sql.NullBool); ok {
				masterData[v.Name()] = z.Bool
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullString); ok {
				masterData[v.Name()] = z.String
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullInt64); ok {
				masterData[v.Name()] = z.Int64
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullFloat64); ok {
				masterData[v.Name()] = z.Float64
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullInt32); ok {
				masterData[v.Name()] = z.Int32
				continue
			}

			masterData[v.Name()] = scanArgs[i]
		}

		finalRows = append(finalRows, masterData)
	}

	resultJson, err := json.Marshal(finalRows)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Write(resultJson)
}
