// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/campaign"
	"affluo/ent/predicate"
	"affluo/ent/referral"
	"affluo/ent/user"
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReferralQuery is the builder for querying Referral entities.
type ReferralQuery struct {
	config
	ctx          *QueryContext
	order        []referral.OrderOption
	inters       []Interceptor
	predicates   []predicate.Referral
	withReferrer *UserQuery
	withCampaign *CampaignQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ReferralQuery builder.
func (rq *ReferralQuery) Where(ps ...predicate.Referral) *ReferralQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *ReferralQuery) Limit(limit int) *ReferralQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *ReferralQuery) Offset(offset int) *ReferralQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *ReferralQuery) Unique(unique bool) *ReferralQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *ReferralQuery) Order(o ...referral.OrderOption) *ReferralQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryReferrer chains the current query on the "referrer" edge.
func (rq *ReferralQuery) QueryReferrer() *UserQuery {
	query := (&UserClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(referral.Table, referral.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, referral.ReferrerTable, referral.ReferrerColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCampaign chains the current query on the "campaign" edge.
func (rq *ReferralQuery) QueryCampaign() *CampaignQuery {
	query := (&CampaignClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(referral.Table, referral.FieldID, selector),
			sqlgraph.To(campaign.Table, campaign.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, referral.CampaignTable, referral.CampaignColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Referral entity from the query.
// Returns a *NotFoundError when no Referral was found.
func (rq *ReferralQuery) First(ctx context.Context) (*Referral, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{referral.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *ReferralQuery) FirstX(ctx context.Context) *Referral {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Referral ID from the query.
// Returns a *NotFoundError when no Referral ID was found.
func (rq *ReferralQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{referral.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *ReferralQuery) FirstIDX(ctx context.Context) int64 {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Referral entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Referral entity is found.
// Returns a *NotFoundError when no Referral entities are found.
func (rq *ReferralQuery) Only(ctx context.Context) (*Referral, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{referral.Label}
	default:
		return nil, &NotSingularError{referral.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *ReferralQuery) OnlyX(ctx context.Context) *Referral {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Referral ID in the query.
// Returns a *NotSingularError when more than one Referral ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *ReferralQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{referral.Label}
	default:
		err = &NotSingularError{referral.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *ReferralQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Referrals.
func (rq *ReferralQuery) All(ctx context.Context) ([]*Referral, error) {
	ctx = setContextOp(ctx, rq.ctx, ent.OpQueryAll)
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Referral, *ReferralQuery]()
	return withInterceptors[[]*Referral](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *ReferralQuery) AllX(ctx context.Context) []*Referral {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Referral IDs.
func (rq *ReferralQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, ent.OpQueryIDs)
	if err = rq.Select(referral.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *ReferralQuery) IDsX(ctx context.Context) []int64 {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *ReferralQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, ent.OpQueryCount)
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*ReferralQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *ReferralQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *ReferralQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, ent.OpQueryExist)
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *ReferralQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ReferralQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *ReferralQuery) Clone() *ReferralQuery {
	if rq == nil {
		return nil
	}
	return &ReferralQuery{
		config:       rq.config,
		ctx:          rq.ctx.Clone(),
		order:        append([]referral.OrderOption{}, rq.order...),
		inters:       append([]Interceptor{}, rq.inters...),
		predicates:   append([]predicate.Referral{}, rq.predicates...),
		withReferrer: rq.withReferrer.Clone(),
		withCampaign: rq.withCampaign.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithReferrer tells the query-builder to eager-load the nodes that are connected to
// the "referrer" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *ReferralQuery) WithReferrer(opts ...func(*UserQuery)) *ReferralQuery {
	query := (&UserClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withReferrer = query
	return rq
}

// WithCampaign tells the query-builder to eager-load the nodes that are connected to
// the "campaign" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *ReferralQuery) WithCampaign(opts ...func(*CampaignQuery)) *ReferralQuery {
	query := (&CampaignClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withCampaign = query
	return rq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Status referral.Status `json:"status,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Referral.Query().
//		GroupBy(referral.FieldStatus).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *ReferralQuery) GroupBy(field string, fields ...string) *ReferralGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ReferralGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = referral.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Status referral.Status `json:"status,omitempty"`
//	}
//
//	client.Referral.Query().
//		Select(referral.FieldStatus).
//		Scan(ctx, &v)
func (rq *ReferralQuery) Select(fields ...string) *ReferralSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &ReferralSelect{ReferralQuery: rq}
	sbuild.label = referral.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ReferralSelect configured with the given aggregations.
func (rq *ReferralQuery) Aggregate(fns ...AggregateFunc) *ReferralSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *ReferralQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !referral.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *ReferralQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Referral, error) {
	var (
		nodes       = []*Referral{}
		withFKs     = rq.withFKs
		_spec       = rq.querySpec()
		loadedTypes = [2]bool{
			rq.withReferrer != nil,
			rq.withCampaign != nil,
		}
	)
	if rq.withReferrer != nil || rq.withCampaign != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, referral.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Referral).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Referral{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withReferrer; query != nil {
		if err := rq.loadReferrer(ctx, query, nodes, nil,
			func(n *Referral, e *User) { n.Edges.Referrer = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withCampaign; query != nil {
		if err := rq.loadCampaign(ctx, query, nodes, nil,
			func(n *Referral, e *Campaign) { n.Edges.Campaign = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *ReferralQuery) loadReferrer(ctx context.Context, query *UserQuery, nodes []*Referral, init func(*Referral), assign func(*Referral, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Referral)
	for i := range nodes {
		if nodes[i].user_referrals == nil {
			continue
		}
		fk := *nodes[i].user_referrals
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
			return fmt.Errorf(`unexpected foreign-key "user_referrals" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (rq *ReferralQuery) loadCampaign(ctx context.Context, query *CampaignQuery, nodes []*Referral, init func(*Referral), assign func(*Referral, *Campaign)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Referral)
	for i := range nodes {
		if nodes[i].campaign_referrals == nil {
			continue
		}
		fk := *nodes[i].campaign_referrals
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(campaign.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "campaign_referrals" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (rq *ReferralQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *ReferralQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(referral.Table, referral.Columns, sqlgraph.NewFieldSpec(referral.FieldID, field.TypeInt64))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, referral.FieldID)
		for i := range fields {
			if fields[i] != referral.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *ReferralQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(referral.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = referral.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ReferralGroupBy is the group-by builder for Referral entities.
type ReferralGroupBy struct {
	selector
	build *ReferralQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *ReferralGroupBy) Aggregate(fns ...AggregateFunc) *ReferralGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *ReferralGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, ent.OpQueryGroupBy)
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReferralQuery, *ReferralGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *ReferralGroupBy) sqlScan(ctx context.Context, root *ReferralQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ReferralSelect is the builder for selecting fields of Referral entities.
type ReferralSelect struct {
	*ReferralQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *ReferralSelect) Aggregate(fns ...AggregateFunc) *ReferralSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *ReferralSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, ent.OpQuerySelect)
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReferralQuery, *ReferralSelect](ctx, rs.ReferralQuery, rs, rs.inters, v)
}

func (rs *ReferralSelect) sqlScan(ctx context.Context, root *ReferralQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}