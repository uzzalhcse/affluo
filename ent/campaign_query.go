// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/banner"
	"affluo/ent/campaign"
	"affluo/ent/campaignlink"
	"affluo/ent/predicate"
	"affluo/ent/user"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CampaignQuery is the builder for querying Campaign entities.
type CampaignQuery struct {
	config
	ctx         *QueryContext
	order       []campaign.OrderOption
	inters      []Interceptor
	predicates  []predicate.Campaign
	withOwner   *UserQuery
	withLinks   *CampaignLinkQuery
	withBanners *BannerQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CampaignQuery builder.
func (cq *CampaignQuery) Where(ps ...predicate.Campaign) *CampaignQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CampaignQuery) Limit(limit int) *CampaignQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CampaignQuery) Offset(offset int) *CampaignQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CampaignQuery) Unique(unique bool) *CampaignQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CampaignQuery) Order(o ...campaign.OrderOption) *CampaignQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryOwner chains the current query on the "owner" edge.
func (cq *CampaignQuery) QueryOwner() *UserQuery {
	query := (&UserClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(campaign.Table, campaign.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, campaign.OwnerTable, campaign.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLinks chains the current query on the "links" edge.
func (cq *CampaignQuery) QueryLinks() *CampaignLinkQuery {
	query := (&CampaignLinkClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(campaign.Table, campaign.FieldID, selector),
			sqlgraph.To(campaignlink.Table, campaignlink.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, campaign.LinksTable, campaign.LinksColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBanners chains the current query on the "banners" edge.
func (cq *CampaignQuery) QueryBanners() *BannerQuery {
	query := (&BannerClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(campaign.Table, campaign.FieldID, selector),
			sqlgraph.To(banner.Table, banner.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, campaign.BannersTable, campaign.BannersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Campaign entity from the query.
// Returns a *NotFoundError when no Campaign was found.
func (cq *CampaignQuery) First(ctx context.Context) (*Campaign, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{campaign.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CampaignQuery) FirstX(ctx context.Context) *Campaign {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Campaign ID from the query.
// Returns a *NotFoundError when no Campaign ID was found.
func (cq *CampaignQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{campaign.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CampaignQuery) FirstIDX(ctx context.Context) int64 {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Campaign entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Campaign entity is found.
// Returns a *NotFoundError when no Campaign entities are found.
func (cq *CampaignQuery) Only(ctx context.Context) (*Campaign, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{campaign.Label}
	default:
		return nil, &NotSingularError{campaign.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CampaignQuery) OnlyX(ctx context.Context) *Campaign {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Campaign ID in the query.
// Returns a *NotSingularError when more than one Campaign ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CampaignQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{campaign.Label}
	default:
		err = &NotSingularError{campaign.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CampaignQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Campaigns.
func (cq *CampaignQuery) All(ctx context.Context) ([]*Campaign, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryAll)
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Campaign, *CampaignQuery]()
	return withInterceptors[[]*Campaign](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CampaignQuery) AllX(ctx context.Context) []*Campaign {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Campaign IDs.
func (cq *CampaignQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryIDs)
	if err = cq.Select(campaign.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CampaignQuery) IDsX(ctx context.Context) []int64 {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CampaignQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryCount)
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CampaignQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CampaignQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CampaignQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryExist)
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CampaignQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CampaignQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CampaignQuery) Clone() *CampaignQuery {
	if cq == nil {
		return nil
	}
	return &CampaignQuery{
		config:      cq.config,
		ctx:         cq.ctx.Clone(),
		order:       append([]campaign.OrderOption{}, cq.order...),
		inters:      append([]Interceptor{}, cq.inters...),
		predicates:  append([]predicate.Campaign{}, cq.predicates...),
		withOwner:   cq.withOwner.Clone(),
		withLinks:   cq.withLinks.Clone(),
		withBanners: cq.withBanners.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CampaignQuery) WithOwner(opts ...func(*UserQuery)) *CampaignQuery {
	query := (&UserClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withOwner = query
	return cq
}

// WithLinks tells the query-builder to eager-load the nodes that are connected to
// the "links" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CampaignQuery) WithLinks(opts ...func(*CampaignLinkQuery)) *CampaignQuery {
	query := (&CampaignLinkClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withLinks = query
	return cq
}

// WithBanners tells the query-builder to eager-load the nodes that are connected to
// the "banners" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CampaignQuery) WithBanners(opts ...func(*BannerQuery)) *CampaignQuery {
	query := (&BannerClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withBanners = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Campaign.Query().
//		GroupBy(campaign.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CampaignQuery) GroupBy(field string, fields ...string) *CampaignGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CampaignGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = campaign.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Campaign.Query().
//		Select(campaign.FieldName).
//		Scan(ctx, &v)
func (cq *CampaignQuery) Select(fields ...string) *CampaignSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CampaignSelect{CampaignQuery: cq}
	sbuild.label = campaign.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CampaignSelect configured with the given aggregations.
func (cq *CampaignQuery) Aggregate(fns ...AggregateFunc) *CampaignSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CampaignQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !campaign.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CampaignQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Campaign, error) {
	var (
		nodes       = []*Campaign{}
		withFKs     = cq.withFKs
		_spec       = cq.querySpec()
		loadedTypes = [3]bool{
			cq.withOwner != nil,
			cq.withLinks != nil,
			cq.withBanners != nil,
		}
	)
	if cq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, campaign.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Campaign).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Campaign{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withOwner; query != nil {
		if err := cq.loadOwner(ctx, query, nodes, nil,
			func(n *Campaign, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withLinks; query != nil {
		if err := cq.loadLinks(ctx, query, nodes,
			func(n *Campaign) { n.Edges.Links = []*CampaignLink{} },
			func(n *Campaign, e *CampaignLink) { n.Edges.Links = append(n.Edges.Links, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withBanners; query != nil {
		if err := cq.loadBanners(ctx, query, nodes,
			func(n *Campaign) { n.Edges.Banners = []*Banner{} },
			func(n *Campaign, e *Banner) { n.Edges.Banners = append(n.Edges.Banners, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CampaignQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*Campaign, init func(*Campaign), assign func(*Campaign, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Campaign)
	for i := range nodes {
		if nodes[i].user_campaigns == nil {
			continue
		}
		fk := *nodes[i].user_campaigns
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
			return fmt.Errorf(`unexpected foreign-key "user_campaigns" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *CampaignQuery) loadLinks(ctx context.Context, query *CampaignLinkQuery, nodes []*Campaign, init func(*Campaign), assign func(*Campaign, *CampaignLink)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Campaign)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.CampaignLink(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(campaign.LinksColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.campaign_links
		if fk == nil {
			return fmt.Errorf(`foreign-key "campaign_links" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "campaign_links" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *CampaignQuery) loadBanners(ctx context.Context, query *BannerQuery, nodes []*Campaign, init func(*Campaign), assign func(*Campaign, *Banner)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int64]*Campaign)
	nids := make(map[int64]map[*Campaign]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(campaign.BannersTable)
		s.Join(joinT).On(s.C(banner.FieldID), joinT.C(campaign.BannersPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(campaign.BannersPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(campaign.BannersPrimaryKey[0]))
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
				outValue := values[0].(*sql.NullInt64).Int64
				inValue := values[1].(*sql.NullInt64).Int64
				if nids[inValue] == nil {
					nids[inValue] = map[*Campaign]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Banner](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "banners" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (cq *CampaignQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CampaignQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(campaign.Table, campaign.Columns, sqlgraph.NewFieldSpec(campaign.FieldID, field.TypeInt64))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, campaign.FieldID)
		for i := range fields {
			if fields[i] != campaign.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CampaignQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(campaign.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = campaign.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CampaignGroupBy is the group-by builder for Campaign entities.
type CampaignGroupBy struct {
	selector
	build *CampaignQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CampaignGroupBy) Aggregate(fns ...AggregateFunc) *CampaignGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CampaignGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, ent.OpQueryGroupBy)
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CampaignQuery, *CampaignGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CampaignGroupBy) sqlScan(ctx context.Context, root *CampaignQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CampaignSelect is the builder for selecting fields of Campaign entities.
type CampaignSelect struct {
	*CampaignQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CampaignSelect) Aggregate(fns ...AggregateFunc) *CampaignSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CampaignSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, ent.OpQuerySelect)
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CampaignQuery, *CampaignSelect](ctx, cs.CampaignQuery, cs, cs.inters, v)
}

func (cs *CampaignSelect) sqlScan(ctx context.Context, root *CampaignQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
