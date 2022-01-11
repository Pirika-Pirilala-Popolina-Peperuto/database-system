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

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Annotations(&entsql.Annotation{
				Default: "gen_random_uuid()",
			},
				entproto.Field(1),
			),
		field.String("name").
			Annotations(entproto.Field(2)),
		field.String("password").
			Sensitive().
			Annotations(entproto.Field(3)),
		field.String("email").
			Annotations(entproto.Field(4)),
		field.String("address").
			Annotations(entproto.Field(5)),
		field.String("user_type").
			Annotations(&entsql.Annotation{
				Default: "customer",
			}, entproto.Field(6)),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("orders", Order.Type).
			Annotations(entproto.Field(7)),
		edge.To("shopping_cart_products", Product.Type).
			StorageKey(edge.Table("shopping_cart")).
			Annotations(entproto.Field(8)),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
