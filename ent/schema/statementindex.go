package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"zotregistry.io/zot/pkg/search/schema"
)

// StatementIndex holds the schema definition for the StatementIndex entity.
type StatementIndex struct {
	ent.Schema
}

// Fields of the StatementIndex.
func (StatementIndex) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("object", schema.Object{}),
		field.JSON("predicate", schema.Predicate{}),
		field.JSON("subject", schema.Subject{}),
		field.JSON("statement", schema.Location{}),
	}
}

// Edges of the StatementIndex.
func (StatementIndex) Edges() []ent.Edge {
	return nil
}
