package schema

import (
	"go-basic-gqlgen/ent/schema/ulid"
	"go-basic-gqlgen/pkg/const/globalid"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Link holds the schema definition for the Link entity.
type Link struct {
	ent.Schema
}

// Fields of the Link.
func (Link) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(ulid.ID("")).
			DefaultFunc(func() ulid.ID {
				return ulid.MustNew(globalid.New().Link.Prefix)
			}),
		field.String("title").
			NotEmpty().
			MaxLen(255),
		field.String("address"),
		field.Time("created_at").Default(time.Now()).Immutable(),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Edges of the Link.
func (Link) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("links").
			Unique(),
	}
}
