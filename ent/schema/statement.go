package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Statement holds the schema definition for the Statement entity.
type Statement struct {
	ent.Schema
}

func (Statement) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}

// Fields of the Statement.
func (Statement) Fields() []ent.Field {
	return []ent.Field{
		field.String("namespace"),
		field.JSON("statement", map[string]interface{}{}),
	}
}

// Edges of the Statement.
func (Statement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("objects", Object.Type),
		edge.To("predicates", Spredicate.Type),
		edge.To("subjects", Subject.Type),
	}

}
