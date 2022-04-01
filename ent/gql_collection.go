// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (e *EmployeeQuery) CollectFields(ctx context.Context, satisfies ...string) *EmployeeQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		e = e.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return e
}

func (e *EmployeeQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *EmployeeQuery {
	return e
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (gr *GroupQuery) CollectFields(ctx context.Context, satisfies ...string) *GroupQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		gr = gr.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return gr
}

func (gr *GroupQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *GroupQuery {
	return gr
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (l *LinkQuery) CollectFields(ctx context.Context, satisfies ...string) *LinkQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		l = l.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return l
}

func (l *LinkQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *LinkQuery {
	return l
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) *UserQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		u = u.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return u
}

func (u *UserQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *UserQuery {
	return u
}
