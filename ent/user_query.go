// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/bannerstats"
	"affluo/ent/campaign"
	"affluo/ent/gigtracking"
	"affluo/ent/payout"
	"affluo/ent/post"
	"affluo/ent/predicate"
	"affluo/ent/referral"
	"affluo/ent/track"
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

// UserQuery is the builder for querying User entities.
type UserQuery struct {
	config
	ctx              *QueryContext
	order            []user.OrderOption
	inters           []Interceptor
	predicates       []predicate.User
	withCampaigns    *CampaignQuery
	withReferrals    *ReferralQuery
	withTracks       *TrackQuery
	withPayouts      *PayoutQuery
	withPosts        *PostQuery
	withStats        *BannerStatsQuery
	withGigTrackings *GigTrackingQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserQuery builder.
func (uq *UserQuery) Where(ps ...predicate.User) *UserQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit the number of records to be returned by this query.
func (uq *UserQuery) Limit(limit int) *UserQuery {
	uq.ctx.Limit = &limit
	return uq
}

// Offset to start from.
func (uq *UserQuery) Offset(offset int) *UserQuery {
	uq.ctx.Offset = &offset
	return uq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uq *UserQuery) Unique(unique bool) *UserQuery {
	uq.ctx.Unique = &unique
	return uq
}

// Order specifies how the records should be ordered.
func (uq *UserQuery) Order(o ...user.OrderOption) *UserQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// QueryCampaigns chains the current query on the "campaigns" edge.
func (uq *UserQuery) QueryCampaigns() *CampaignQuery {
	query := (&CampaignClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(campaign.Table, campaign.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.CampaignsTable, user.CampaignsColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReferrals chains the current query on the "referrals" edge.
func (uq *UserQuery) QueryReferrals() *ReferralQuery {
	query := (&ReferralClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(referral.Table, referral.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.ReferralsTable, user.ReferralsColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTracks chains the current query on the "tracks" edge.
func (uq *UserQuery) QueryTracks() *TrackQuery {
	query := (&TrackClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(track.Table, track.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.TracksTable, user.TracksColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPayouts chains the current query on the "payouts" edge.
func (uq *UserQuery) QueryPayouts() *PayoutQuery {
	query := (&PayoutClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(payout.Table, payout.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.PayoutsTable, user.PayoutsColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPosts chains the current query on the "posts" edge.
func (uq *UserQuery) QueryPosts() *PostQuery {
	query := (&PostClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.PostsTable, user.PostsColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStats chains the current query on the "stats" edge.
func (uq *UserQuery) QueryStats() *BannerStatsQuery {
	query := (&BannerStatsClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(bannerstats.Table, bannerstats.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.StatsTable, user.StatsColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGigTrackings chains the current query on the "gig_trackings" edge.
func (uq *UserQuery) QueryGigTrackings() *GigTrackingQuery {
	query := (&GigTrackingClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(gigtracking.Table, gigtracking.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, user.GigTrackingsTable, user.GigTrackingsColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first User entity from the query.
// Returns a *NotFoundError when no User was found.
func (uq *UserQuery) First(ctx context.Context) (*User, error) {
	nodes, err := uq.Limit(1).All(setContextOp(ctx, uq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{user.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UserQuery) FirstX(ctx context.Context) *User {
	node, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first User ID from the query.
// Returns a *NotFoundError when no User ID was found.
func (uq *UserQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = uq.Limit(1).IDs(setContextOp(ctx, uq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{user.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uq *UserQuery) FirstIDX(ctx context.Context) int64 {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single User entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one User entity is found.
// Returns a *NotFoundError when no User entities are found.
func (uq *UserQuery) Only(ctx context.Context) (*User, error) {
	nodes, err := uq.Limit(2).All(setContextOp(ctx, uq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{user.Label}
	default:
		return nil, &NotSingularError{user.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UserQuery) OnlyX(ctx context.Context) *User {
	node, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only User ID in the query.
// Returns a *NotSingularError when more than one User ID is found.
// Returns a *NotFoundError when no entities are found.
func (uq *UserQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = uq.Limit(2).IDs(setContextOp(ctx, uq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{user.Label}
	default:
		err = &NotSingularError{user.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uq *UserQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Users.
func (uq *UserQuery) All(ctx context.Context) ([]*User, error) {
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryAll)
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*User, *UserQuery]()
	return withInterceptors[[]*User](ctx, uq, qr, uq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uq *UserQuery) AllX(ctx context.Context) []*User {
	nodes, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of User IDs.
func (uq *UserQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if uq.ctx.Unique == nil && uq.path != nil {
		uq.Unique(true)
	}
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryIDs)
	if err = uq.Select(user.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UserQuery) IDsX(ctx context.Context) []int64 {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryCount)
	if err := uq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uq, querierCount[*UserQuery](), uq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UserQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryExist)
	switch _, err := uq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UserQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UserQuery) Clone() *UserQuery {
	if uq == nil {
		return nil
	}
	return &UserQuery{
		config:           uq.config,
		ctx:              uq.ctx.Clone(),
		order:            append([]user.OrderOption{}, uq.order...),
		inters:           append([]Interceptor{}, uq.inters...),
		predicates:       append([]predicate.User{}, uq.predicates...),
		withCampaigns:    uq.withCampaigns.Clone(),
		withReferrals:    uq.withReferrals.Clone(),
		withTracks:       uq.withTracks.Clone(),
		withPayouts:      uq.withPayouts.Clone(),
		withPosts:        uq.withPosts.Clone(),
		withStats:        uq.withStats.Clone(),
		withGigTrackings: uq.withGigTrackings.Clone(),
		// clone intermediate query.
		sql:  uq.sql.Clone(),
		path: uq.path,
	}
}

// WithCampaigns tells the query-builder to eager-load the nodes that are connected to
// the "campaigns" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithCampaigns(opts ...func(*CampaignQuery)) *UserQuery {
	query := (&CampaignClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withCampaigns = query
	return uq
}

// WithReferrals tells the query-builder to eager-load the nodes that are connected to
// the "referrals" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithReferrals(opts ...func(*ReferralQuery)) *UserQuery {
	query := (&ReferralClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withReferrals = query
	return uq
}

// WithTracks tells the query-builder to eager-load the nodes that are connected to
// the "tracks" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithTracks(opts ...func(*TrackQuery)) *UserQuery {
	query := (&TrackClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withTracks = query
	return uq
}

// WithPayouts tells the query-builder to eager-load the nodes that are connected to
// the "payouts" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithPayouts(opts ...func(*PayoutQuery)) *UserQuery {
	query := (&PayoutClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withPayouts = query
	return uq
}

// WithPosts tells the query-builder to eager-load the nodes that are connected to
// the "posts" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithPosts(opts ...func(*PostQuery)) *UserQuery {
	query := (&PostClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withPosts = query
	return uq
}

// WithStats tells the query-builder to eager-load the nodes that are connected to
// the "stats" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithStats(opts ...func(*BannerStatsQuery)) *UserQuery {
	query := (&BannerStatsClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withStats = query
	return uq
}

// WithGigTrackings tells the query-builder to eager-load the nodes that are connected to
// the "gig_trackings" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithGigTrackings(opts ...func(*GigTrackingQuery)) *UserQuery {
	query := (&GigTrackingClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withGigTrackings = query
	return uq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.User.Query().
//		GroupBy(user.FieldUsername).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uq *UserQuery) GroupBy(field string, fields ...string) *UserGroupBy {
	uq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserGroupBy{build: uq}
	grbuild.flds = &uq.ctx.Fields
	grbuild.label = user.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//	}
//
//	client.User.Query().
//		Select(user.FieldUsername).
//		Scan(ctx, &v)
func (uq *UserQuery) Select(fields ...string) *UserSelect {
	uq.ctx.Fields = append(uq.ctx.Fields, fields...)
	sbuild := &UserSelect{UserQuery: uq}
	sbuild.label = user.Label
	sbuild.flds, sbuild.scan = &uq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserSelect configured with the given aggregations.
func (uq *UserQuery) Aggregate(fns ...AggregateFunc) *UserSelect {
	return uq.Select().Aggregate(fns...)
}

func (uq *UserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uq); err != nil {
				return err
			}
		}
	}
	for _, f := range uq.ctx.Fields {
		if !user.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uq.path != nil {
		prev, err := uq.path(ctx)
		if err != nil {
			return err
		}
		uq.sql = prev
	}
	return nil
}

func (uq *UserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*User, error) {
	var (
		nodes       = []*User{}
		_spec       = uq.querySpec()
		loadedTypes = [7]bool{
			uq.withCampaigns != nil,
			uq.withReferrals != nil,
			uq.withTracks != nil,
			uq.withPayouts != nil,
			uq.withPosts != nil,
			uq.withStats != nil,
			uq.withGigTrackings != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*User).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &User{config: uq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uq.withCampaigns; query != nil {
		if err := uq.loadCampaigns(ctx, query, nodes,
			func(n *User) { n.Edges.Campaigns = []*Campaign{} },
			func(n *User, e *Campaign) { n.Edges.Campaigns = append(n.Edges.Campaigns, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withReferrals; query != nil {
		if err := uq.loadReferrals(ctx, query, nodes,
			func(n *User) { n.Edges.Referrals = []*Referral{} },
			func(n *User, e *Referral) { n.Edges.Referrals = append(n.Edges.Referrals, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withTracks; query != nil {
		if err := uq.loadTracks(ctx, query, nodes,
			func(n *User) { n.Edges.Tracks = []*Track{} },
			func(n *User, e *Track) { n.Edges.Tracks = append(n.Edges.Tracks, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withPayouts; query != nil {
		if err := uq.loadPayouts(ctx, query, nodes,
			func(n *User) { n.Edges.Payouts = []*Payout{} },
			func(n *User, e *Payout) { n.Edges.Payouts = append(n.Edges.Payouts, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withPosts; query != nil {
		if err := uq.loadPosts(ctx, query, nodes,
			func(n *User) { n.Edges.Posts = []*Post{} },
			func(n *User, e *Post) { n.Edges.Posts = append(n.Edges.Posts, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withStats; query != nil {
		if err := uq.loadStats(ctx, query, nodes,
			func(n *User) { n.Edges.Stats = []*BannerStats{} },
			func(n *User, e *BannerStats) { n.Edges.Stats = append(n.Edges.Stats, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withGigTrackings; query != nil {
		if err := uq.loadGigTrackings(ctx, query, nodes,
			func(n *User) { n.Edges.GigTrackings = []*GigTracking{} },
			func(n *User, e *GigTracking) { n.Edges.GigTrackings = append(n.Edges.GigTrackings, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uq *UserQuery) loadCampaigns(ctx context.Context, query *CampaignQuery, nodes []*User, init func(*User), assign func(*User, *Campaign)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Campaign(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.CampaignsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_campaigns
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_campaigns" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_campaigns" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadReferrals(ctx context.Context, query *ReferralQuery, nodes []*User, init func(*User), assign func(*User, *Referral)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Referral(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.ReferralsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_referrals
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_referrals" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_referrals" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadTracks(ctx context.Context, query *TrackQuery, nodes []*User, init func(*User), assign func(*User, *Track)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Track(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.TracksColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_tracks
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_tracks" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_tracks" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadPayouts(ctx context.Context, query *PayoutQuery, nodes []*User, init func(*User), assign func(*User, *Payout)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Payout(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.PayoutsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_payouts
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_payouts" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_payouts" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadPosts(ctx context.Context, query *PostQuery, nodes []*User, init func(*User), assign func(*User, *Post)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Post(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.PostsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_posts
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_posts" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_posts" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadStats(ctx context.Context, query *BannerStatsQuery, nodes []*User, init func(*User), assign func(*User, *BannerStats)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.BannerStats(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.StatsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_stats
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_stats" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_stats" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadGigTrackings(ctx context.Context, query *GigTrackingQuery, nodes []*User, init func(*User), assign func(*User, *GigTracking)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.GigTracking(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.GigTrackingsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.gig_tracking_publisher
		if fk == nil {
			return fmt.Errorf(`foreign-key "gig_tracking_publisher" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "gig_tracking_publisher" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (uq *UserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uq.querySpec()
	_spec.Node.Columns = uq.ctx.Fields
	if len(uq.ctx.Fields) > 0 {
		_spec.Unique = uq.ctx.Unique != nil && *uq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, uq.driver, _spec)
}

func (uq *UserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64))
	_spec.From = uq.sql
	if unique := uq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if uq.path != nil {
		_spec.Unique = true
	}
	if fields := uq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for i := range fields {
			if fields[i] != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uq *UserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uq.driver.Dialect())
	t1 := builder.Table(user.Table)
	columns := uq.ctx.Fields
	if len(columns) == 0 {
		columns = user.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uq.sql != nil {
		selector = uq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uq.ctx.Unique != nil && *uq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range uq.predicates {
		p(selector)
	}
	for _, p := range uq.order {
		p(selector)
	}
	if offset := uq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserGroupBy is the group-by builder for User entities.
type UserGroupBy struct {
	selector
	build *UserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UserGroupBy) Aggregate(fns ...AggregateFunc) *UserGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the selector query and scans the result into the given value.
func (ugb *UserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ugb.build.ctx, ent.OpQueryGroupBy)
	if err := ugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserQuery, *UserGroupBy](ctx, ugb.build, ugb, ugb.build.inters, v)
}

func (ugb *UserGroupBy) sqlScan(ctx context.Context, root *UserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ugb.fns))
	for _, fn := range ugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ugb.flds)+len(ugb.fns))
		for _, f := range *ugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserSelect is the builder for selecting fields of User entities.
type UserSelect struct {
	*UserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (us *UserSelect) Aggregate(fns ...AggregateFunc) *UserSelect {
	us.fns = append(us.fns, fns...)
	return us
}

// Scan applies the selector query and scans the result into the given value.
func (us *UserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, us.ctx, ent.OpQuerySelect)
	if err := us.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserQuery, *UserSelect](ctx, us.UserQuery, us, us.inters, v)
}

func (us *UserSelect) sqlScan(ctx context.Context, root *UserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(us.fns))
	for _, fn := range us.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*us.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := us.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
