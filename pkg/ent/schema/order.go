package schema

import (
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Annotations(&entsql.Annotation{
				Default: "gen_random_uuid()",
			},
				entproto.Field(1),
			),
		field.UUID("customer_id", uuid.UUID{}).
			Annotations(entproto.Field(2)),
		field.Time("created_at").
			Default(time.Now).
			Annotations(entproto.Field(3)),
		field.Enum("process_status").
			Values("UNPAID", "PROCESSING", "COMPLETE").
			Annotations(
				entproto.Field(4),
				entproto.Enum(map[string]int32{
					"UNPAID":     1,
					"PROCESSING": 2,
					"COMPLETE":   3,
				}),
			),
		field.String("remark").Optional().
			Annotations(entproto.Field(5)),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("orders").
			Field("customer_id").
			Unique().
			Required().
			Annotations(entproto.Field(6)),
		edge.To("products", Product.Type).
			Annotations(entproto.Field(7)),
		edge.To("discounts", Discount.Type).
			Annotations(entproto.Field(8)),
	}
}

func (Order) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
