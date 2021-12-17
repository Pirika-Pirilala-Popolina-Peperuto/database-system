package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Discount holds the schema definition for the Discount entity.
type Discount struct {
	ent.Schema
}

// Fields of the Discount.
func (Discount) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Annotations(&entsql.Annotation{
				Default: "gen_random_uuid()",
			},
				entproto.Field(1),
			).
			Annotations(),
		field.Time("period_start").
			Optional().
			Annotations(entproto.Field(2)),
		field.Time("period_end").
			Optional().
			Annotations(entproto.Field(3)),
		field.Enum("method").
			Values("PERCENTAGE", "PRICE").
			Annotations(entproto.Field(4), entproto.Enum(map[string]int32{
				"PERCENTAGE": 1,
				"PRICE":      2,
			})),
		field.Int("discount_price").
			Annotations(entproto.Field(5)),
		field.Float("discount_percentage").
			Annotations(entproto.Field(6)),
	}

}

// Edges of the Discount.
func (Discount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("orders", Order.Type).
			Ref("discounts").
			Annotations(entproto.Field(7)),
	}
}

func (Discount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
