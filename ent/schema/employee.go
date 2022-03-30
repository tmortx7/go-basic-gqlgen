package schema

import (
	"go-basic-gqlgen/ent/schema/ulid"
	"go-basic-gqlgen/pkg/const/globalid"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Employee holds the schema definition for the Employee entity.
type Employee struct {
	ent.Schema
}

// Fields of the Employee.
func (Employee) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(ulid.ID("")).
			DefaultFunc(func() ulid.ID {
				return ulid.MustNew(globalid.New().Employee.Prefix)
			}),
		field.String("firstName"),
		field.String("lastName"),
		field.Time("created_at").Default(time.Now()).Immutable(),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Edges of the Employee.
func (Employee) Edges() []ent.Edge {
	return nil
}
