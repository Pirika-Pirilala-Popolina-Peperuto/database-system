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

// Picture holds the schema definition for the Picture entity.
type Picture struct {
	ent.Schema
}

// Fields of the Picture.
func (Picture) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Annotations(&entsql.Annotation{
				Default: "gen_random_uuid()",
			},
				entproto.Field(1),
			),
		field.String("picture_url").
			Annotations(entproto.Field(2)),
	}
}

// Edges of the Picture.
func (Picture) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product", Product.Type).
			Unique().
			Annotations(entproto.Field(3)),
	}
}

func (Picture) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
