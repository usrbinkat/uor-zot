// Code generated by ent, DO NOT EDIT.

package object

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"zotregistry.io/zot/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Object {
	return predicate.Object(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Object {
	return predicate.Object(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Object {
	return predicate.Object(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Object {
	return predicate.Object(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Object {
	return predicate.Object(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Object {
	return predicate.Object(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Object {
	return predicate.Object(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Object {
	return predicate.Object(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Object {
	return predicate.Object(sql.FieldLTE(FieldID, id))
}

// ObjectType applies equality check predicate on the "objectType" field. It's identical to ObjectTypeEQ.
func ObjectType(v string) predicate.Object {
	return predicate.Object(sql.FieldEQ(FieldObjectType, v))
}

// ObjectTypeEQ applies the EQ predicate on the "objectType" field.
func ObjectTypeEQ(v string) predicate.Object {
	return predicate.Object(sql.FieldEQ(FieldObjectType, v))
}

// ObjectTypeNEQ applies the NEQ predicate on the "objectType" field.
func ObjectTypeNEQ(v string) predicate.Object {
	return predicate.Object(sql.FieldNEQ(FieldObjectType, v))
}

// ObjectTypeIn applies the In predicate on the "objectType" field.
func ObjectTypeIn(vs ...string) predicate.Object {
	return predicate.Object(sql.FieldIn(FieldObjectType, vs...))
}

// ObjectTypeNotIn applies the NotIn predicate on the "objectType" field.
func ObjectTypeNotIn(vs ...string) predicate.Object {
	return predicate.Object(sql.FieldNotIn(FieldObjectType, vs...))
}

// ObjectTypeGT applies the GT predicate on the "objectType" field.
func ObjectTypeGT(v string) predicate.Object {
	return predicate.Object(sql.FieldGT(FieldObjectType, v))
}

// ObjectTypeGTE applies the GTE predicate on the "objectType" field.
func ObjectTypeGTE(v string) predicate.Object {
	return predicate.Object(sql.FieldGTE(FieldObjectType, v))
}

// ObjectTypeLT applies the LT predicate on the "objectType" field.
func ObjectTypeLT(v string) predicate.Object {
	return predicate.Object(sql.FieldLT(FieldObjectType, v))
}

// ObjectTypeLTE applies the LTE predicate on the "objectType" field.
func ObjectTypeLTE(v string) predicate.Object {
	return predicate.Object(sql.FieldLTE(FieldObjectType, v))
}

// ObjectTypeContains applies the Contains predicate on the "objectType" field.
func ObjectTypeContains(v string) predicate.Object {
	return predicate.Object(sql.FieldContains(FieldObjectType, v))
}

// ObjectTypeHasPrefix applies the HasPrefix predicate on the "objectType" field.
func ObjectTypeHasPrefix(v string) predicate.Object {
	return predicate.Object(sql.FieldHasPrefix(FieldObjectType, v))
}

// ObjectTypeHasSuffix applies the HasSuffix predicate on the "objectType" field.
func ObjectTypeHasSuffix(v string) predicate.Object {
	return predicate.Object(sql.FieldHasSuffix(FieldObjectType, v))
}

// ObjectTypeEqualFold applies the EqualFold predicate on the "objectType" field.
func ObjectTypeEqualFold(v string) predicate.Object {
	return predicate.Object(sql.FieldEqualFold(FieldObjectType, v))
}

// ObjectTypeContainsFold applies the ContainsFold predicate on the "objectType" field.
func ObjectTypeContainsFold(v string) predicate.Object {
	return predicate.Object(sql.FieldContainsFold(FieldObjectType, v))
}

// HasStatement applies the HasEdge predicate on the "statement" edge.
func HasStatement() predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, StatementTable, StatementPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatementWith applies the HasEdge predicate on the "statement" edge with a given conditions (other predicates).
func HasStatementWith(preds ...predicate.Statement) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		step := newStatementStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Object) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Object) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Object) predicate.Object {
	return predicate.Object(func(s *sql.Selector) {
		p(s.Not())
	})
}
