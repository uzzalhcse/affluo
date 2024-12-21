// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/banner"
	"affluo/ent/lead"
	"affluo/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LeadQuery is the builder for querying Lead entities.
type LeadQuery struct {
	config
	ctx        *QueryContext
	order      []lead.OrderOption
	inters     []Interceptor
	predicates []predicate.Lead
	withBanner *BannerQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LeadQuery builder.
func (lq *LeadQuery) Where(ps ...predicate.Lead) *LeadQuery {
	lq.predicates = append(lq.predicates, ps...)
	return lq
}

// Limit the number of records to be returned by this query.
func (lq *LeadQuery) Limit(limit int) *LeadQuery {
	lq.ctx.Limit = &limit
	return lq
}

// Offset to start from.
func (lq *LeadQuery) Offset(offset int) *LeadQuery {
	lq.ctx.Offset = &offset
	return lq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lq *LeadQuery) Unique(unique bool) *LeadQuery {
	lq.ctx.Unique = &unique
	return lq
}

// Order specifies how the records should be ordered.
func (lq *LeadQuery) Order(o ...lead.OrderOption) *LeadQuery {
	lq.order = append(lq.order, o...)
	return lq
}

// QueryBanner chains the current query on the "banner" edge.
func (lq *LeadQuery) QueryBanner() *BannerQuery {
	query := (&BannerClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(lead.Table, lead.FieldID, selector),
			sqlgraph.To(banner.Table, banner.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, lead.BannerTable, lead.BannerColumn),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Lead entity from the query.
// Returns a *NotFoundError when no Lead was found.
func (lq *LeadQuery) First(ctx context.Context) (*Lead, error) {
	nodes, err := lq.Limit(1).All(setContextOp(ctx, lq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{lead.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lq *LeadQuery) FirstX(ctx context.Context) *Lead {
	node, err := lq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Lead ID from the query.
// Returns a *NotFoundError when no Lead ID was found.
func (lq *LeadQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = lq.Limit(1).IDs(setContextOp(ctx, lq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{lead.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lq *LeadQuery) FirstIDX(ctx context.Context) int64 {
	id, err := lq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Lead entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Lead entity is found.
// Returns a *NotFoundError when no Lead entities are found.
func (lq *LeadQuery) Only(ctx context.Context) (*Lead, error) {
	nodes, err := lq.Limit(2).All(setContextOp(ctx, lq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{lead.Label}
	default:
		return nil, &NotSingularError{lead.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lq *LeadQuery) OnlyX(ctx context.Context) *Lead {
	node, err := lq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Lead ID in the query.
// Returns a *NotSingularError when more than one Lead ID is found.
// Returns a *NotFoundError when no entities are found.
func (lq *LeadQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = lq.Limit(2).IDs(setContextOp(ctx, lq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{lead.Label}
	default:
		err = &NotSingularError{lead.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lq *LeadQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := lq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Leads.
func (lq *LeadQuery) All(ctx context.Context) ([]*Lead, error) {
	ctx = setContextOp(ctx, lq.ctx, ent.OpQueryAll)
	if err := lq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Lead, *LeadQuery]()
	return withInterceptors[[]*Lead](ctx, lq, qr, lq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lq *LeadQuery) AllX(ctx context.Context) []*Lead {
	nodes, err := lq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Lead IDs.
func (lq *LeadQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if lq.ctx.Unique == nil && lq.path != nil {
		lq.Unique(true)
	}
	ctx = setContextOp(ctx, lq.ctx, ent.OpQueryIDs)
	if err = lq.Select(lead.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lq *LeadQuery) IDsX(ctx context.Context) []int64 {
	ids, err := lq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lq *LeadQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lq.ctx, ent.OpQueryCount)
	if err := lq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lq, querierCount[*LeadQuery](), lq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lq *LeadQuery) CountX(ctx context.Context) int {
	count, err := lq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lq *LeadQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lq.ctx, ent.OpQueryExist)
	switch _, err := lq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lq *LeadQuery) ExistX(ctx context.Context) bool {
	exist, err := lq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LeadQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lq *LeadQuery) Clone() *LeadQuery {
	if lq == nil {
		return nil
	}
	return &LeadQuery{
		config:     lq.config,
		ctx:        lq.ctx.Clone(),
		order:      append([]lead.OrderOption{}, lq.order...),
		inters:     append([]Interceptor{}, lq.inters...),
		predicates: append([]predicate.Lead{}, lq.predicates...),
		withBanner: lq.withBanner.Clone(),
		// clone intermediate query.
		sql:  lq.sql.Clone(),
		path: lq.path,
	}
}

// WithBanner tells the query-builder to eager-load the nodes that are connected to
// the "banner" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LeadQuery) WithBanner(opts ...func(*BannerQuery)) *LeadQuery {
	query := (&BannerClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withBanner = query
	return lq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ReferenceID string `json:"reference_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Lead.Query().
//		GroupBy(lead.FieldReferenceID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lq *LeadQuery) GroupBy(field string, fields ...string) *LeadGroupBy {
	lq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LeadGroupBy{build: lq}
	grbuild.flds = &lq.ctx.Fields
	grbuild.label = lead.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ReferenceID string `json:"reference_id,omitempty"`
//	}
//
//	client.Lead.Query().
//		Select(lead.FieldReferenceID).
//		Scan(ctx, &v)
func (lq *LeadQuery) Select(fields ...string) *LeadSelect {
	lq.ctx.Fields = append(lq.ctx.Fields, fields...)
	sbuild := &LeadSelect{LeadQuery: lq}
	sbuild.label = lead.Label
	sbuild.flds, sbuild.scan = &lq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LeadSelect configured with the given aggregations.
func (lq *LeadQuery) Aggregate(fns ...AggregateFunc) *LeadSelect {
	return lq.Select().Aggregate(fns...)
}

func (lq *LeadQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lq); err != nil {
				return err
			}
		}
	}
	for _, f := range lq.ctx.Fields {
		if !lead.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lq.path != nil {
		prev, err := lq.path(ctx)
		if err != nil {
			return err
		}
		lq.sql = prev
	}
	return nil
}

func (lq *LeadQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Lead, error) {
	var (
		nodes       = []*Lead{}
		withFKs     = lq.withFKs
		_spec       = lq.querySpec()
		loadedTypes = [1]bool{
			lq.withBanner != nil,
		}
	)
	if lq.withBanner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, lead.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Lead).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Lead{config: lq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lq.withBanner; query != nil {
		if err := lq.loadBanner(ctx, query, nodes, nil,
			func(n *Lead, e *Banner) { n.Edges.Banner = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lq *LeadQuery) loadBanner(ctx context.Context, query *BannerQuery, nodes []*Lead, init func(*Lead), assign func(*Lead, *Banner)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Lead)
	for i := range nodes {
		if nodes[i].banner_leads == nil {
			continue
		}
		fk := *nodes[i].banner_leads
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(banner.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "banner_leads" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (lq *LeadQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lq.querySpec()
	_spec.Node.Columns = lq.ctx.Fields
	if len(lq.ctx.Fields) > 0 {
		_spec.Unique = lq.ctx.Unique != nil && *lq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lq.driver, _spec)
}

func (lq *LeadQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(lead.Table, lead.Columns, sqlgraph.NewFieldSpec(lead.FieldID, field.TypeInt64))
	_spec.From = lq.sql
	if unique := lq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lq.path != nil {
		_spec.Unique = true
	}
	if fields := lq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lead.FieldID)
		for i := range fields {
			if fields[i] != lead.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lq *LeadQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lq.driver.Dialect())
	t1 := builder.Table(lead.Table)
	columns := lq.ctx.Fields
	if len(columns) == 0 {
		columns = lead.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lq.sql != nil {
		selector = lq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lq.ctx.Unique != nil && *lq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lq.predicates {
		p(selector)
	}
	for _, p := range lq.order {
		p(selector)
	}
	if offset := lq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LeadGroupBy is the group-by builder for Lead entities.
type LeadGroupBy struct {
	selector
	build *LeadQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lgb *LeadGroupBy) Aggregate(fns ...AggregateFunc) *LeadGroupBy {
	lgb.fns = append(lgb.fns, fns...)
	return lgb
}

// Scan applies the selector query and scans the result into the given value.
func (lgb *LeadGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lgb.build.ctx, ent.OpQueryGroupBy)
	if err := lgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LeadQuery, *LeadGroupBy](ctx, lgb.build, lgb, lgb.build.inters, v)
}

func (lgb *LeadGroupBy) sqlScan(ctx context.Context, root *LeadQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lgb.fns))
	for _, fn := range lgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lgb.flds)+len(lgb.fns))
		for _, f := range *lgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LeadSelect is the builder for selecting fields of Lead entities.
type LeadSelect struct {
	*LeadQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ls *LeadSelect) Aggregate(fns ...AggregateFunc) *LeadSelect {
	ls.fns = append(ls.fns, fns...)
	return ls
}

// Scan applies the selector query and scans the result into the given value.
func (ls *LeadSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ls.ctx, ent.OpQuerySelect)
	if err := ls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LeadQuery, *LeadSelect](ctx, ls.LeadQuery, ls, ls.inters, v)
}

func (ls *LeadSelect) sqlScan(ctx context.Context, root *LeadQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ls.fns))
	for _, fn := range ls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
