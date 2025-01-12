// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/earninghistory"
	"affluo/ent/predicate"
	"affluo/ent/user"
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EarningHistoryQuery is the builder for querying EarningHistory entities.
type EarningHistoryQuery struct {
	config
	ctx        *QueryContext
	order      []earninghistory.OrderOption
	inters     []Interceptor
	predicates []predicate.EarningHistory
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EarningHistoryQuery builder.
func (ehq *EarningHistoryQuery) Where(ps ...predicate.EarningHistory) *EarningHistoryQuery {
	ehq.predicates = append(ehq.predicates, ps...)
	return ehq
}

// Limit the number of records to be returned by this query.
func (ehq *EarningHistoryQuery) Limit(limit int) *EarningHistoryQuery {
	ehq.ctx.Limit = &limit
	return ehq
}

// Offset to start from.
func (ehq *EarningHistoryQuery) Offset(offset int) *EarningHistoryQuery {
	ehq.ctx.Offset = &offset
	return ehq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ehq *EarningHistoryQuery) Unique(unique bool) *EarningHistoryQuery {
	ehq.ctx.Unique = &unique
	return ehq
}

// Order specifies how the records should be ordered.
func (ehq *EarningHistoryQuery) Order(o ...earninghistory.OrderOption) *EarningHistoryQuery {
	ehq.order = append(ehq.order, o...)
	return ehq
}

// QueryUser chains the current query on the "user" edge.
func (ehq *EarningHistoryQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: ehq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ehq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ehq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(earninghistory.Table, earninghistory.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, earninghistory.UserTable, earninghistory.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ehq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EarningHistory entity from the query.
// Returns a *NotFoundError when no EarningHistory was found.
func (ehq *EarningHistoryQuery) First(ctx context.Context) (*EarningHistory, error) {
	nodes, err := ehq.Limit(1).All(setContextOp(ctx, ehq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{earninghistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ehq *EarningHistoryQuery) FirstX(ctx context.Context) *EarningHistory {
	node, err := ehq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EarningHistory ID from the query.
// Returns a *NotFoundError when no EarningHistory ID was found.
func (ehq *EarningHistoryQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ehq.Limit(1).IDs(setContextOp(ctx, ehq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{earninghistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ehq *EarningHistoryQuery) FirstIDX(ctx context.Context) int64 {
	id, err := ehq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EarningHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EarningHistory entity is found.
// Returns a *NotFoundError when no EarningHistory entities are found.
func (ehq *EarningHistoryQuery) Only(ctx context.Context) (*EarningHistory, error) {
	nodes, err := ehq.Limit(2).All(setContextOp(ctx, ehq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{earninghistory.Label}
	default:
		return nil, &NotSingularError{earninghistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ehq *EarningHistoryQuery) OnlyX(ctx context.Context) *EarningHistory {
	node, err := ehq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EarningHistory ID in the query.
// Returns a *NotSingularError when more than one EarningHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (ehq *EarningHistoryQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ehq.Limit(2).IDs(setContextOp(ctx, ehq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{earninghistory.Label}
	default:
		err = &NotSingularError{earninghistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ehq *EarningHistoryQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := ehq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EarningHistories.
func (ehq *EarningHistoryQuery) All(ctx context.Context) ([]*EarningHistory, error) {
	ctx = setContextOp(ctx, ehq.ctx, ent.OpQueryAll)
	if err := ehq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EarningHistory, *EarningHistoryQuery]()
	return withInterceptors[[]*EarningHistory](ctx, ehq, qr, ehq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ehq *EarningHistoryQuery) AllX(ctx context.Context) []*EarningHistory {
	nodes, err := ehq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EarningHistory IDs.
func (ehq *EarningHistoryQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if ehq.ctx.Unique == nil && ehq.path != nil {
		ehq.Unique(true)
	}
	ctx = setContextOp(ctx, ehq.ctx, ent.OpQueryIDs)
	if err = ehq.Select(earninghistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ehq *EarningHistoryQuery) IDsX(ctx context.Context) []int64 {
	ids, err := ehq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ehq *EarningHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ehq.ctx, ent.OpQueryCount)
	if err := ehq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ehq, querierCount[*EarningHistoryQuery](), ehq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ehq *EarningHistoryQuery) CountX(ctx context.Context) int {
	count, err := ehq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ehq *EarningHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ehq.ctx, ent.OpQueryExist)
	switch _, err := ehq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ehq *EarningHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := ehq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EarningHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ehq *EarningHistoryQuery) Clone() *EarningHistoryQuery {
	if ehq == nil {
		return nil
	}
	return &EarningHistoryQuery{
		config:     ehq.config,
		ctx:        ehq.ctx.Clone(),
		order:      append([]earninghistory.OrderOption{}, ehq.order...),
		inters:     append([]Interceptor{}, ehq.inters...),
		predicates: append([]predicate.EarningHistory{}, ehq.predicates...),
		withUser:   ehq.withUser.Clone(),
		// clone intermediate query.
		sql:  ehq.sql.Clone(),
		path: ehq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ehq *EarningHistoryQuery) WithUser(opts ...func(*UserQuery)) *EarningHistoryQuery {
	query := (&UserClient{config: ehq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ehq.withUser = query
	return ehq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Amount float64 `json:"amount,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EarningHistory.Query().
//		GroupBy(earninghistory.FieldAmount).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ehq *EarningHistoryQuery) GroupBy(field string, fields ...string) *EarningHistoryGroupBy {
	ehq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EarningHistoryGroupBy{build: ehq}
	grbuild.flds = &ehq.ctx.Fields
	grbuild.label = earninghistory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Amount float64 `json:"amount,omitempty"`
//	}
//
//	client.EarningHistory.Query().
//		Select(earninghistory.FieldAmount).
//		Scan(ctx, &v)
func (ehq *EarningHistoryQuery) Select(fields ...string) *EarningHistorySelect {
	ehq.ctx.Fields = append(ehq.ctx.Fields, fields...)
	sbuild := &EarningHistorySelect{EarningHistoryQuery: ehq}
	sbuild.label = earninghistory.Label
	sbuild.flds, sbuild.scan = &ehq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EarningHistorySelect configured with the given aggregations.
func (ehq *EarningHistoryQuery) Aggregate(fns ...AggregateFunc) *EarningHistorySelect {
	return ehq.Select().Aggregate(fns...)
}

func (ehq *EarningHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ehq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ehq); err != nil {
				return err
			}
		}
	}
	for _, f := range ehq.ctx.Fields {
		if !earninghistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ehq.path != nil {
		prev, err := ehq.path(ctx)
		if err != nil {
			return err
		}
		ehq.sql = prev
	}
	return nil
}

func (ehq *EarningHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EarningHistory, error) {
	var (
		nodes       = []*EarningHistory{}
		withFKs     = ehq.withFKs
		_spec       = ehq.querySpec()
		loadedTypes = [1]bool{
			ehq.withUser != nil,
		}
	)
	if ehq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, earninghistory.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EarningHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EarningHistory{config: ehq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ehq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ehq.withUser; query != nil {
		if err := ehq.loadUser(ctx, query, nodes, nil,
			func(n *EarningHistory, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ehq *EarningHistoryQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*EarningHistory, init func(*EarningHistory), assign func(*EarningHistory, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*EarningHistory)
	for i := range nodes {
		if nodes[i].user_earning_histories == nil {
			continue
		}
		fk := *nodes[i].user_earning_histories
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_earning_histories" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ehq *EarningHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ehq.querySpec()
	_spec.Node.Columns = ehq.ctx.Fields
	if len(ehq.ctx.Fields) > 0 {
		_spec.Unique = ehq.ctx.Unique != nil && *ehq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ehq.driver, _spec)
}

func (ehq *EarningHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(earninghistory.Table, earninghistory.Columns, sqlgraph.NewFieldSpec(earninghistory.FieldID, field.TypeInt64))
	_spec.From = ehq.sql
	if unique := ehq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ehq.path != nil {
		_spec.Unique = true
	}
	if fields := ehq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, earninghistory.FieldID)
		for i := range fields {
			if fields[i] != earninghistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ehq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ehq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ehq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ehq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ehq *EarningHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ehq.driver.Dialect())
	t1 := builder.Table(earninghistory.Table)
	columns := ehq.ctx.Fields
	if len(columns) == 0 {
		columns = earninghistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ehq.sql != nil {
		selector = ehq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ehq.ctx.Unique != nil && *ehq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ehq.predicates {
		p(selector)
	}
	for _, p := range ehq.order {
		p(selector)
	}
	if offset := ehq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ehq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EarningHistoryGroupBy is the group-by builder for EarningHistory entities.
type EarningHistoryGroupBy struct {
	selector
	build *EarningHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ehgb *EarningHistoryGroupBy) Aggregate(fns ...AggregateFunc) *EarningHistoryGroupBy {
	ehgb.fns = append(ehgb.fns, fns...)
	return ehgb
}

// Scan applies the selector query and scans the result into the given value.
func (ehgb *EarningHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ehgb.build.ctx, ent.OpQueryGroupBy)
	if err := ehgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EarningHistoryQuery, *EarningHistoryGroupBy](ctx, ehgb.build, ehgb, ehgb.build.inters, v)
}

func (ehgb *EarningHistoryGroupBy) sqlScan(ctx context.Context, root *EarningHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ehgb.fns))
	for _, fn := range ehgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ehgb.flds)+len(ehgb.fns))
		for _, f := range *ehgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ehgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ehgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EarningHistorySelect is the builder for selecting fields of EarningHistory entities.
type EarningHistorySelect struct {
	*EarningHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ehs *EarningHistorySelect) Aggregate(fns ...AggregateFunc) *EarningHistorySelect {
	ehs.fns = append(ehs.fns, fns...)
	return ehs
}

// Scan applies the selector query and scans the result into the given value.
func (ehs *EarningHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ehs.ctx, ent.OpQuerySelect)
	if err := ehs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EarningHistoryQuery, *EarningHistorySelect](ctx, ehs.EarningHistoryQuery, ehs, ehs.inters, v)
}

func (ehs *EarningHistorySelect) sqlScan(ctx context.Context, root *EarningHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ehs.fns))
	for _, fn := range ehs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ehs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ehs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
