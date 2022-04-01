package schema

import (
	"go-basic-gqlgen/ent/schema/ulid"
	"go-basic-gqlgen/pkg/const/globalid"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(ulid.ID("")).
			DefaultFunc(func() ulid.ID {
				return ulid.MustNew(globalid.New().User.Prefix)
			}),
		field.String("name").
			NotEmpty().
			MaxLen(255),
		field.Time("created_at").Default(time.Now()).Immutable(),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("links", Link.Type).
			StorageKey(
				edge.Column("user_id"),
				edge.Symbol("user_id"),
			),
	}
}
