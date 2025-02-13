// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"zotregistry.io/zot/ent/object"
	"zotregistry.io/zot/ent/predicate"
	"zotregistry.io/zot/ent/statement"
)

// ObjectQuery is the builder for querying Object entities.
type ObjectQuery struct {
	config
	ctx                *QueryContext
	order              []object.OrderOption
	inters             []Interceptor
	predicates         []predicate.Object
	withStatement      *StatementQuery
	modifiers          []func(*sql.Selector)
	loadTotal          []func(context.Context, []*Object) error
	withNamedStatement map[string]*StatementQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ObjectQuery builder.
func (oq *ObjectQuery) Where(ps ...predicate.Object) *ObjectQuery {
	oq.predicates = append(oq.predicates, ps...)
	return oq
}

// Limit the number of records to be returned by this query.
func (oq *ObjectQuery) Limit(limit int) *ObjectQuery {
	oq.ctx.Limit = &limit
	return oq
}

// Offset to start from.
func (oq *ObjectQuery) Offset(offset int) *ObjectQuery {
	oq.ctx.Offset = &offset
	return oq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oq *ObjectQuery) Unique(unique bool) *ObjectQuery {
	oq.ctx.Unique = &unique
	return oq
}

// Order specifies how the records should be ordered.
func (oq *ObjectQuery) Order(o ...object.OrderOption) *ObjectQuery {
	oq.order = append(oq.order, o...)
	return oq
}

// QueryStatement chains the current query on the "statement" edge.
func (oq *ObjectQuery) QueryStatement() *StatementQuery {
	query := (&StatementClient{config: oq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(object.Table, object.FieldID, selector),
			sqlgraph.To(statement.Table, statement.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, object.StatementTable, object.StatementPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Object entity from the query.
// Returns a *NotFoundError when no Object was found.
func (oq *ObjectQuery) First(ctx context.Context) (*Object, error) {
	nodes, err := oq.Limit(1).All(setContextOp(ctx, oq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{object.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oq *ObjectQuery) FirstX(ctx context.Context) *Object {
	node, err := oq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Object ID from the query.
// Returns a *NotFoundError when no Object ID was found.
func (oq *ObjectQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(1).IDs(setContextOp(ctx, oq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{object.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oq *ObjectQuery) FirstIDX(ctx context.Context) int {
	id, err := oq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Object entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Object entity is found.
// Returns a *NotFoundError when no Object entities are found.
func (oq *ObjectQuery) Only(ctx context.Context) (*Object, error) {
	nodes, err := oq.Limit(2).All(setContextOp(ctx, oq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{object.Label}
	default:
		return nil, &NotSingularError{object.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oq *ObjectQuery) OnlyX(ctx context.Context) *Object {
	node, err := oq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Object ID in the query.
// Returns a *NotSingularError when more than one Object ID is found.
// Returns a *NotFoundError when no entities are found.
func (oq *ObjectQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(2).IDs(setContextOp(ctx, oq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{object.Label}
	default:
		err = &NotSingularError{object.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oq *ObjectQuery) OnlyIDX(ctx context.Context) int {
	id, err := oq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Objects.
func (oq *ObjectQuery) All(ctx context.Context) ([]*Object, error) {
	ctx = setContextOp(ctx, oq.ctx, "All")
	if err := oq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Object, *ObjectQuery]()
	return withInterceptors[[]*Object](ctx, oq, qr, oq.inters)
}

// AllX is like All, but panics if an error occurs.
func (oq *ObjectQuery) AllX(ctx context.Context) []*Object {
	nodes, err := oq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Object IDs.
func (oq *ObjectQuery) IDs(ctx context.Context) (ids []int, err error) {
	if oq.ctx.Unique == nil && oq.path != nil {
		oq.Unique(true)
	}
	ctx = setContextOp(ctx, oq.ctx, "IDs")
	if err = oq.Select(object.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oq *ObjectQuery) IDsX(ctx context.Context) []int {
	ids, err := oq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oq *ObjectQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, oq.ctx, "Count")
	if err := oq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, oq, querierCount[*ObjectQuery](), oq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (oq *ObjectQuery) CountX(ctx context.Context) int {
	count, err := oq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oq *ObjectQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, oq.ctx, "Exist")
	switch _, err := oq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oq *ObjectQuery) ExistX(ctx context.Context) bool {
	exist, err := oq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ObjectQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oq *ObjectQuery) Clone() *ObjectQuery {
	if oq == nil {
		return nil
	}
	return &ObjectQuery{
		config:        oq.config,
		ctx:           oq.ctx.Clone(),
		order:         append([]object.OrderOption{}, oq.order...),
		inters:        append([]Interceptor{}, oq.inters...),
		predicates:    append([]predicate.Object{}, oq.predicates...),
		withStatement: oq.withStatement.Clone(),
		// clone intermediate query.
		sql:  oq.sql.Clone(),
		path: oq.path,
	}
}

// WithStatement tells the query-builder to eager-load the nodes that are connected to
// the "statement" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *ObjectQuery) WithStatement(opts ...func(*StatementQuery)) *ObjectQuery {
	query := (&StatementClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oq.withStatement = query
	return oq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ObjectType string `json:"objectType,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Object.Query().
//		GroupBy(object.FieldObjectType).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (oq *ObjectQuery) GroupBy(field string, fields ...string) *ObjectGroupBy {
	oq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ObjectGroupBy{build: oq}
	grbuild.flds = &oq.ctx.Fields
	grbuild.label = object.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ObjectType string `json:"objectType,omitempty"`
//	}
//
//	client.Object.Query().
//		Select(object.FieldObjectType).
//		Scan(ctx, &v)
func (oq *ObjectQuery) Select(fields ...string) *ObjectSelect {
	oq.ctx.Fields = append(oq.ctx.Fields, fields...)
	sbuild := &ObjectSelect{ObjectQuery: oq}
	sbuild.label = object.Label
	sbuild.flds, sbuild.scan = &oq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ObjectSelect configured with the given aggregations.
func (oq *ObjectQuery) Aggregate(fns ...AggregateFunc) *ObjectSelect {
	return oq.Select().Aggregate(fns...)
}

func (oq *ObjectQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range oq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, oq); err != nil {
				return err
			}
		}
	}
	for _, f := range oq.ctx.Fields {
		if !object.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oq.path != nil {
		prev, err := oq.path(ctx)
		if err != nil {
			return err
		}
		oq.sql = prev
	}
	return nil
}

func (oq *ObjectQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Object, error) {
	var (
		nodes       = []*Object{}
		_spec       = oq.querySpec()
		loadedTypes = [1]bool{
			oq.withStatement != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Object).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Object{config: oq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(oq.modifiers) > 0 {
		_spec.Modifiers = oq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := oq.withStatement; query != nil {
		if err := oq.loadStatement(ctx, query, nodes,
			func(n *Object) { n.Edges.Statement = []*Statement{} },
			func(n *Object, e *Statement) { n.Edges.Statement = append(n.Edges.Statement, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range oq.withNamedStatement {
		if err := oq.loadStatement(ctx, query, nodes,
			func(n *Object) { n.appendNamedStatement(name) },
			func(n *Object, e *Statement) { n.appendNamedStatement(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range oq.loadTotal {
		if err := oq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (oq *ObjectQuery) loadStatement(ctx context.Context, query *StatementQuery, nodes []*Object, init func(*Object), assign func(*Object, *Statement)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Object)
	nids := make(map[int]map[*Object]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(object.StatementTable)
		s.Join(joinT).On(s.C(statement.FieldID), joinT.C(object.StatementPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(object.StatementPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(object.StatementPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Object]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Statement](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "statement" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (oq *ObjectQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oq.querySpec()
	if len(oq.modifiers) > 0 {
		_spec.Modifiers = oq.modifiers
	}
	_spec.Node.Columns = oq.ctx.Fields
	if len(oq.ctx.Fields) > 0 {
		_spec.Unique = oq.ctx.Unique != nil && *oq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, oq.driver, _spec)
}

func (oq *ObjectQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(object.Table, object.Columns, sqlgraph.NewFieldSpec(object.FieldID, field.TypeInt))
	_spec.From = oq.sql
	if unique := oq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if oq.path != nil {
		_spec.Unique = true
	}
	if fields := oq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, object.FieldID)
		for i := range fields {
			if fields[i] != object.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oq *ObjectQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oq.driver.Dialect())
	t1 := builder.Table(object.Table)
	columns := oq.ctx.Fields
	if len(columns) == 0 {
		columns = object.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oq.sql != nil {
		selector = oq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oq.ctx.Unique != nil && *oq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range oq.predicates {
		p(selector)
	}
	for _, p := range oq.order {
		p(selector)
	}
	if offset := oq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedStatement tells the query-builder to eager-load the nodes that are connected to the "statement"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (oq *ObjectQuery) WithNamedStatement(name string, opts ...func(*StatementQuery)) *ObjectQuery {
	query := (&StatementClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if oq.withNamedStatement == nil {
		oq.withNamedStatement = make(map[string]*StatementQuery)
	}
	oq.withNamedStatement[name] = query
	return oq
}

// ObjectGroupBy is the group-by builder for Object entities.
type ObjectGroupBy struct {
	selector
	build *ObjectQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ogb *ObjectGroupBy) Aggregate(fns ...AggregateFunc) *ObjectGroupBy {
	ogb.fns = append(ogb.fns, fns...)
	return ogb
}

// Scan applies the selector query and scans the result into the given value.
func (ogb *ObjectGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ogb.build.ctx, "GroupBy")
	if err := ogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ObjectQuery, *ObjectGroupBy](ctx, ogb.build, ogb, ogb.build.inters, v)
}

func (ogb *ObjectGroupBy) sqlScan(ctx context.Context, root *ObjectQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ogb.fns))
	for _, fn := range ogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ogb.flds)+len(ogb.fns))
		for _, f := range *ogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ObjectSelect is the builder for selecting fields of Object entities.
type ObjectSelect struct {
	*ObjectQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (os *ObjectSelect) Aggregate(fns ...AggregateFunc) *ObjectSelect {
	os.fns = append(os.fns, fns...)
	return os
}

// Scan applies the selector query and scans the result into the given value.
func (os *ObjectSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, os.ctx, "Select")
	if err := os.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ObjectQuery, *ObjectSelect](ctx, os.ObjectQuery, os, os.inters, v)
}

func (os *ObjectSelect) sqlScan(ctx context.Context, root *ObjectQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(os.fns))
	for _, fn := range os.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*os.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := os.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
