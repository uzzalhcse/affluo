// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/commissionhistory"
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

// CommissionHistoryQuery is the builder for querying CommissionHistory entities.
type CommissionHistoryQuery struct {
	config
	ctx        *QueryContext
	order      []commissionhistory.OrderOption
	inters     []Interceptor
	predicates []predicate.CommissionHistory
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CommissionHistoryQuery builder.
func (chq *CommissionHistoryQuery) Where(ps ...predicate.CommissionHistory) *CommissionHistoryQuery {
	chq.predicates = append(chq.predicates, ps...)
	return chq
}

// Limit the number of records to be returned by this query.
func (chq *CommissionHistoryQuery) Limit(limit int) *CommissionHistoryQuery {
	chq.ctx.Limit = &limit
	return chq
}

// Offset to start from.
func (chq *CommissionHistoryQuery) Offset(offset int) *CommissionHistoryQuery {
	chq.ctx.Offset = &offset
	return chq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (chq *CommissionHistoryQuery) Unique(unique bool) *CommissionHistoryQuery {
	chq.ctx.Unique = &unique
	return chq
}

// Order specifies how the records should be ordered.
func (chq *CommissionHistoryQuery) Order(o ...commissionhistory.OrderOption) *CommissionHistoryQuery {
	chq.order = append(chq.order, o...)
	return chq
}

// QueryUser chains the current query on the "user" edge.
func (chq *CommissionHistoryQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: chq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := chq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := chq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(commissionhistory.Table, commissionhistory.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, commissionhistory.UserTable, commissionhistory.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(chq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CommissionHistory entity from the query.
// Returns a *NotFoundError when no CommissionHistory was found.
func (chq *CommissionHistoryQuery) First(ctx context.Context) (*CommissionHistory, error) {
	nodes, err := chq.Limit(1).All(setContextOp(ctx, chq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{commissionhistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (chq *CommissionHistoryQuery) FirstX(ctx context.Context) *CommissionHistory {
	node, err := chq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CommissionHistory ID from the query.
// Returns a *NotFoundError when no CommissionHistory ID was found.
func (chq *CommissionHistoryQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = chq.Limit(1).IDs(setContextOp(ctx, chq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{commissionhistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (chq *CommissionHistoryQuery) FirstIDX(ctx context.Context) int64 {
	id, err := chq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CommissionHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CommissionHistory entity is found.
// Returns a *NotFoundError when no CommissionHistory entities are found.
func (chq *CommissionHistoryQuery) Only(ctx context.Context) (*CommissionHistory, error) {
	nodes, err := chq.Limit(2).All(setContextOp(ctx, chq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{commissionhistory.Label}
	default:
		return nil, &NotSingularError{commissionhistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (chq *CommissionHistoryQuery) OnlyX(ctx context.Context) *CommissionHistory {
	node, err := chq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CommissionHistory ID in the query.
// Returns a *NotSingularError when more than one CommissionHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (chq *CommissionHistoryQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = chq.Limit(2).IDs(setContextOp(ctx, chq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{commissionhistory.Label}
	default:
		err = &NotSingularError{commissionhistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (chq *CommissionHistoryQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := chq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CommissionHistories.
func (chq *CommissionHistoryQuery) All(ctx context.Context) ([]*CommissionHistory, error) {
	ctx = setContextOp(ctx, chq.ctx, ent.OpQueryAll)
	if err := chq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CommissionHistory, *CommissionHistoryQuery]()
	return withInterceptors[[]*CommissionHistory](ctx, chq, qr, chq.inters)
}

// AllX is like All, but panics if an error occurs.
func (chq *CommissionHistoryQuery) AllX(ctx context.Context) []*CommissionHistory {
	nodes, err := chq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CommissionHistory IDs.
func (chq *CommissionHistoryQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if chq.ctx.Unique == nil && chq.path != nil {
		chq.Unique(true)
	}
	ctx = setContextOp(ctx, chq.ctx, ent.OpQueryIDs)
	if err = chq.Select(commissionhistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (chq *CommissionHistoryQuery) IDsX(ctx context.Context) []int64 {
	ids, err := chq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (chq *CommissionHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, chq.ctx, ent.OpQueryCount)
	if err := chq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, chq, querierCount[*CommissionHistoryQuery](), chq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (chq *CommissionHistoryQuery) CountX(ctx context.Context) int {
	count, err := chq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (chq *CommissionHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, chq.ctx, ent.OpQueryExist)
	switch _, err := chq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (chq *CommissionHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := chq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CommissionHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (chq *CommissionHistoryQuery) Clone() *CommissionHistoryQuery {
	if chq == nil {
		return nil
	}
	return &CommissionHistoryQuery{
		config:     chq.config,
		ctx:        chq.ctx.Clone(),
		order:      append([]commissionhistory.OrderOption{}, chq.order...),
		inters:     append([]Interceptor{}, chq.inters...),
		predicates: append([]predicate.CommissionHistory{}, chq.predicates...),
		withUser:   chq.withUser.Clone(),
		// clone intermediate query.
		sql:  chq.sql.Clone(),
		path: chq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (chq *CommissionHistoryQuery) WithUser(opts ...func(*UserQuery)) *CommissionHistoryQuery {
	query := (&UserClient{config: chq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	chq.withUser = query
	return chq
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
//	client.CommissionHistory.Query().
//		GroupBy(commissionhistory.FieldAmount).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (chq *CommissionHistoryQuery) GroupBy(field string, fields ...string) *CommissionHistoryGroupBy {
	chq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CommissionHistoryGroupBy{build: chq}
	grbuild.flds = &chq.ctx.Fields
	grbuild.label = commissionhistory.Label
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
//	client.CommissionHistory.Query().
//		Select(commissionhistory.FieldAmount).
//		Scan(ctx, &v)
func (chq *CommissionHistoryQuery) Select(fields ...string) *CommissionHistorySelect {
	chq.ctx.Fields = append(chq.ctx.Fields, fields...)
	sbuild := &CommissionHistorySelect{CommissionHistoryQuery: chq}
	sbuild.label = commissionhistory.Label
	sbuild.flds, sbuild.scan = &chq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CommissionHistorySelect configured with the given aggregations.
func (chq *CommissionHistoryQuery) Aggregate(fns ...AggregateFunc) *CommissionHistorySelect {
	return chq.Select().Aggregate(fns...)
}

func (chq *CommissionHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range chq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, chq); err != nil {
				return err
			}
		}
	}
	for _, f := range chq.ctx.Fields {
		if !commissionhistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if chq.path != nil {
		prev, err := chq.path(ctx)
		if err != nil {
			return err
		}
		chq.sql = prev
	}
	return nil
}

func (chq *CommissionHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CommissionHistory, error) {
	var (
		nodes       = []*CommissionHistory{}
		withFKs     = chq.withFKs
		_spec       = chq.querySpec()
		loadedTypes = [1]bool{
			chq.withUser != nil,
		}
	)
	if chq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, commissionhistory.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CommissionHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CommissionHistory{config: chq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, chq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := chq.withUser; query != nil {
		if err := chq.loadUser(ctx, query, nodes, nil,
			func(n *CommissionHistory, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (chq *CommissionHistoryQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*CommissionHistory, init func(*CommissionHistory), assign func(*CommissionHistory, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*CommissionHistory)
	for i := range nodes {
		if nodes[i].user_commission_histories == nil {
			continue
		}
		fk := *nodes[i].user_commission_histories
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
			return fmt.Errorf(`unexpected foreign-key "user_commission_histories" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (chq *CommissionHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := chq.querySpec()
	_spec.Node.Columns = chq.ctx.Fields
	if len(chq.ctx.Fields) > 0 {
		_spec.Unique = chq.ctx.Unique != nil && *chq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, chq.driver, _spec)
}

func (chq *CommissionHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(commissionhistory.Table, commissionhistory.Columns, sqlgraph.NewFieldSpec(commissionhistory.FieldID, field.TypeInt64))
	_spec.From = chq.sql
	if unique := chq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if chq.path != nil {
		_spec.Unique = true
	}
	if fields := chq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, commissionhistory.FieldID)
		for i := range fields {
			if fields[i] != commissionhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := chq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := chq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := chq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := chq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (chq *CommissionHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(chq.driver.Dialect())
	t1 := builder.Table(commissionhistory.Table)
	columns := chq.ctx.Fields
	if len(columns) == 0 {
		columns = commissionhistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if chq.sql != nil {
		selector = chq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if chq.ctx.Unique != nil && *chq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range chq.predicates {
		p(selector)
	}
	for _, p := range chq.order {
		p(selector)
	}
	if offset := chq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := chq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CommissionHistoryGroupBy is the group-by builder for CommissionHistory entities.
type CommissionHistoryGroupBy struct {
	selector
	build *CommissionHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (chgb *CommissionHistoryGroupBy) Aggregate(fns ...AggregateFunc) *CommissionHistoryGroupBy {
	chgb.fns = append(chgb.fns, fns...)
	return chgb
}

// Scan applies the selector query and scans the result into the given value.
func (chgb *CommissionHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, chgb.build.ctx, ent.OpQueryGroupBy)
	if err := chgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommissionHistoryQuery, *CommissionHistoryGroupBy](ctx, chgb.build, chgb, chgb.build.inters, v)
}

func (chgb *CommissionHistoryGroupBy) sqlScan(ctx context.Context, root *CommissionHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(chgb.fns))
	for _, fn := range chgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*chgb.flds)+len(chgb.fns))
		for _, f := range *chgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*chgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := chgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CommissionHistorySelect is the builder for selecting fields of CommissionHistory entities.
type CommissionHistorySelect struct {
	*CommissionHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (chs *CommissionHistorySelect) Aggregate(fns ...AggregateFunc) *CommissionHistorySelect {
	chs.fns = append(chs.fns, fns...)
	return chs
}

// Scan applies the selector query and scans the result into the given value.
func (chs *CommissionHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, chs.ctx, ent.OpQuerySelect)
	if err := chs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommissionHistoryQuery, *CommissionHistorySelect](ctx, chs.CommissionHistoryQuery, chs, chs.inters, v)
}

func (chs *CommissionHistorySelect) sqlScan(ctx context.Context, root *CommissionHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(chs.fns))
	for _, fn := range chs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*chs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := chs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
