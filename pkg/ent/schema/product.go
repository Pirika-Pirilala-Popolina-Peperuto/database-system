package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Money float64

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Annotations(&entsql.Annotation{
				Default: "gen_random_uuid()",
			},
				entproto.Field(1),
			),
		field.String("name").
			Annotations(entproto.Field(2)),
		field.String("description").Optional().
			Annotations(entproto.Field(3)),
		field.Float("price").SchemaType(map[string]string{
			dialect.Postgres: "decimal(12,2)",
		}).
			Annotations(entproto.Field(4)),
		field.Int("quantity").
			Annotations(entproto.Field(5)),
		field.UUID("picture_id", uuid.UUID{}).
			Annotations(entproto.Field(6)),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("orders", Order.Type).
			Ref("products").
			Annotations(entproto.Field(7)),
		edge.To("categories", Category.Type).
			Annotations(entproto.Field(8)),
		edge.From("shopping_cart_owners", User.Type).
			Ref("shopping_cart_products").
			Annotations(entproto.Field(9)),
		edge.From("picture", Picture.Type).
			Ref("product").
			Unique().
			Required().
			Field("picture_id").
			Annotations(entproto.Field(10)),
	}
}

func (Product) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
