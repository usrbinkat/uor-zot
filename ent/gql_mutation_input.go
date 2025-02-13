// Code generated by ent, DO NOT EDIT.

package ent

// CreateObjectInput represents a mutation input for creating objects.
type CreateObjectInput struct {
	ObjectType   string
	Object       map[string]interface{}
	StatementIDs []int
}

// Mutate applies the CreateObjectInput on the ObjectMutation builder.
func (i *CreateObjectInput) Mutate(m *ObjectMutation) {
	m.SetObjectType(i.ObjectType)
	if v := i.Object; v != nil {
		m.SetObject(v)
	}
	if v := i.StatementIDs; len(v) > 0 {
		m.AddStatementIDs(v...)
	}
}

// SetInput applies the change-set in the CreateObjectInput on the ObjectCreate builder.
func (c *ObjectCreate) SetInput(i CreateObjectInput) *ObjectCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateSpredicateInput represents a mutation input for creating spredicates.
type CreateSpredicateInput struct {
	PredicateType string
	Predicate     map[string]interface{}
	StatementIDs  []int
}

// Mutate applies the CreateSpredicateInput on the SpredicateMutation builder.
func (i *CreateSpredicateInput) Mutate(m *SpredicateMutation) {
	m.SetPredicateType(i.PredicateType)
	if v := i.Predicate; v != nil {
		m.SetPredicate(v)
	}
	if v := i.StatementIDs; len(v) > 0 {
		m.AddStatementIDs(v...)
	}
}

// SetInput applies the change-set in the CreateSpredicateInput on the SpredicateCreate builder.
func (c *SpredicateCreate) SetInput(i CreateSpredicateInput) *SpredicateCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateStatementInput represents a mutation input for creating statements.
type CreateStatementInput struct {
	Namespace    string
	Statement    map[string]interface{}
	ObjectIDs    []int
	PredicateIDs []int
	SubjectIDs   []int
}

// Mutate applies the CreateStatementInput on the StatementMutation builder.
func (i *CreateStatementInput) Mutate(m *StatementMutation) {
	m.SetNamespace(i.Namespace)
	if v := i.Statement; v != nil {
		m.SetStatement(v)
	}
	if v := i.ObjectIDs; len(v) > 0 {
		m.AddObjectIDs(v...)
	}
	if v := i.PredicateIDs; len(v) > 0 {
		m.AddPredicateIDs(v...)
	}
	if v := i.SubjectIDs; len(v) > 0 {
		m.AddSubjectIDs(v...)
	}
}

// SetInput applies the change-set in the CreateStatementInput on the StatementCreate builder.
func (c *StatementCreate) SetInput(i CreateStatementInput) *StatementCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateSubjectInput represents a mutation input for creating subjects.
type CreateSubjectInput struct {
	SubjectType  string
	Subject      map[string]interface{}
	StatementIDs []int
}

// Mutate applies the CreateSubjectInput on the SubjectMutation builder.
func (i *CreateSubjectInput) Mutate(m *SubjectMutation) {
	m.SetSubjectType(i.SubjectType)
	if v := i.Subject; v != nil {
		m.SetSubject(v)
	}
	if v := i.StatementIDs; len(v) > 0 {
		m.AddStatementIDs(v...)
	}
}

// SetInput applies the change-set in the CreateSubjectInput on the SubjectCreate builder.
func (c *SubjectCreate) SetInput(i CreateSubjectInput) *SubjectCreate {
	i.Mutate(c.Mutation())
	return c
}
