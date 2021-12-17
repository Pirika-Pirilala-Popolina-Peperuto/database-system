package ent

import (
	"context"

	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func init() {
	viper.SetDefault("postgres.host", "localhost")
	viper.SetDefault("postgres.port", 5432)
	viper.SetDefault("postgres.user", "admin")
	viper.SetDefault("postgres.dbname", "postgres")
	viper.SetDefault("postgres.password", "12345678")
}

func NewPostgresClientWithLC(lc fx.Lifecycle) (*Client, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		viper.GetString("postgres.host"),
		viper.GetInt("postgres.port"),
		viper.GetString("postgres.user"),
		viper.GetString("postgres.dbname"),
		viper.GetString("postgres.password"),
	)
	db, err := Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if viper.GetBool("debug") {
		db = db.Debug()
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return db.Schema.Create(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})

	return db, nil
}
