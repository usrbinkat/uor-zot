package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Spredicate holds the schema definition for the Spredicate entity.
type Spredicate struct {
	ent.Schema
}

func (Spredicate) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}

// Fields of the Spredicate.
func (Spredicate) Fields() []ent.Field {
	return []ent.Field{
		field.String("predicateType"),
		field.JSON("predicate", map[string]interface{}{}),
	}
}

// Edges of the Spredicate.
func (Spredicate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("statement", Statement.Type).
			Ref("predicates"),
	}
}
