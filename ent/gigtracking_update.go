// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/gigtracking"
	"affluo/ent/predicate"
	"affluo/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GigTrackingUpdate is the builder for updating GigTracking entities.
type GigTrackingUpdate struct {
	config
	hooks    []Hook
	mutation *GigTrackingMutation
}

// Where appends a list predicates to the GigTrackingUpdate builder.
func (gtu *GigTrackingUpdate) Where(ps ...predicate.GigTracking) *GigTrackingUpdate {
	gtu.mutation.Where(ps...)
	return gtu
}

// SetDate sets the "date" field.
func (gtu *GigTrackingUpdate) SetDate(t time.Time) *GigTrackingUpdate {
	gtu.mutation.SetDate(t)
	return gtu
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (gtu *GigTrackingUpdate) SetNillableDate(t *time.Time) *GigTrackingUpdate {
	if t != nil {
		gtu.SetDate(*t)
	}
	return gtu
}

// SetAffiliateUserID sets the "affiliate_user_id" field.
func (gtu *GigTrackingUpdate) SetAffiliateUserID(s string) *GigTrackingUpdate {
	gtu.mutation.SetAffiliateUserID(s)
	return gtu
}

// SetNillableAffiliateUserID sets the "affiliate_user_id" field if the given value is not nil.
func (gtu *GigTrackingUpdate) SetNillableAffiliateUserID(s *string) *GigTrackingUpdate {
	if s != nil {
		gtu.SetAffiliateUserID(*s)
	}
	return gtu
}

// ClearAffiliateUserID clears the value of the "affiliate_user_id" field.
func (gtu *GigTrackingUpdate) ClearAffiliateUserID() *GigTrackingUpdate {
	gtu.mutation.ClearAffiliateUserID()
	return gtu
}

// SetType sets the "type" field.
func (gtu *GigTrackingUpdate) SetType(s string) *GigTrackingUpdate {
	gtu.mutation.SetType(s)
	return gtu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (gtu *GigTrackingUpdate) SetNillableType(s *string) *GigTrackingUpdate {
	if s != nil {
		gtu.SetType(*s)
	}
	return gtu
}

// SetUtmQuery sets the "utm_query" field.
func (gtu *GigTrackingUpdate) SetUtmQuery(s string) *GigTrackingUpdate {
	gtu.mutation.SetUtmQuery(s)
	return gtu
}

// SetNillableUtmQuery sets the "utm_query" field if the given value is not nil.
func (gtu *GigTrackingUpdate) SetNillableUtmQuery(s *string) *GigTrackingUpdate {
	if s != nil {
		gtu.SetUtmQuery(*s)
	}
	return gtu
}

// ClearUtmQuery clears the value of the "utm_query" field.
func (gtu *GigTrackingUpdate) ClearUtmQuery() *GigTrackingUpdate {
	gtu.mutation.ClearUtmQuery()
	return gtu
}

// SetLp sets the "lp" field.
func (gtu *GigTrackingUpdate) SetLp(s string) *GigTrackingUpdate {
	gtu.mutation.SetLp(s)
	return gtu
}

// SetNillableLp sets the "lp" field if the given value is not nil.
func (gtu *GigTrackingUpdate) SetNillableLp(s *string) *GigTrackingUpdate {
	if s != nil {
		gtu.SetLp(*s)
	}
	return gtu
}

// ClearLp clears the value of the "lp" field.
func (gtu *GigTrackingUpdate) ClearLp() *GigTrackingUpdate {
	gtu.mutation.ClearLp()
	return gtu
}

// SetTrackID sets the "track_id" field.
func (gtu *GigTrackingUpdate) SetTrackID(s string) *GigTrackingUpdate {
	gtu.mutation.SetTrackID(s)
	return gtu
}

// SetNillableTrackID sets the "track_id" field if the given value is not nil.
func (gtu *GigTrackingUpdate) SetNillableTrackID(s *string) *GigTrackingUpdate {
	if s != nil {
		gtu.SetTrackID(*s)
	}
	return gtu
}

// ClearTrackID clears the value of the "track_id" field.
func (gtu *GigTrackingUpdate) ClearTrackID() *GigTrackingUpdate {
	gtu.mutation.ClearTrackID()
	return gtu
}

// SetRevenue sets the "revenue" field.
func (gtu *GigTrackingUpdate) SetRevenue(f float64) *GigTrackingUpdate {
	gtu.mutation.ResetRevenue()
	gtu.mutation.SetRevenue(f)
	return gtu
}

// SetNillableRevenue sets the "revenue" field if the given value is not nil.
func (gtu *GigTrackingUpdate) SetNillableRevenue(f *float64) *GigTrackingUpdate {
	if f != nil {
		gtu.SetRevenue(*f)
	}
	return gtu
}

// AddRevenue adds f to the "revenue" field.
func (gtu *GigTrackingUpdate) AddRevenue(f float64) *GigTrackingUpdate {
	gtu.mutation.AddRevenue(f)
	return gtu
}

// SetCreatedAt sets the "created_at" field.
func (gtu *GigTrackingUpdate) SetCreatedAt(t time.Time) *GigTrackingUpdate {
	gtu.mutation.SetCreatedAt(t)
	return gtu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gtu *GigTrackingUpdate) SetNillableCreatedAt(t *time.Time) *GigTrackingUpdate {
	if t != nil {
		gtu.SetCreatedAt(*t)
	}
	return gtu
}

// SetUpdatedAt sets the "updated_at" field.
func (gtu *GigTrackingUpdate) SetUpdatedAt(t time.Time) *GigTrackingUpdate {
	gtu.mutation.SetUpdatedAt(t)
	return gtu
}

// SetPublisherID sets the "publisher" edge to the User entity by ID.
func (gtu *GigTrackingUpdate) SetPublisherID(id int64) *GigTrackingUpdate {
	gtu.mutation.SetPublisherID(id)
	return gtu
}

// SetPublisher sets the "publisher" edge to the User entity.
func (gtu *GigTrackingUpdate) SetPublisher(u *User) *GigTrackingUpdate {
	return gtu.SetPublisherID(u.ID)
}

// Mutation returns the GigTrackingMutation object of the builder.
func (gtu *GigTrackingUpdate) Mutation() *GigTrackingMutation {
	return gtu.mutation
}

// ClearPublisher clears the "publisher" edge to the User entity.
func (gtu *GigTrackingUpdate) ClearPublisher() *GigTrackingUpdate {
	gtu.mutation.ClearPublisher()
	return gtu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gtu *GigTrackingUpdate) Save(ctx context.Context) (int, error) {
	gtu.defaults()
	return withHooks(ctx, gtu.sqlSave, gtu.mutation, gtu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gtu *GigTrackingUpdate) SaveX(ctx context.Context) int {
	affected, err := gtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gtu *GigTrackingUpdate) Exec(ctx context.Context) error {
	_, err := gtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gtu *GigTrackingUpdate) ExecX(ctx context.Context) {
	if err := gtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gtu *GigTrackingUpdate) defaults() {
	if _, ok := gtu.mutation.UpdatedAt(); !ok {
		v := gigtracking.UpdateDefaultUpdatedAt()
		gtu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gtu *GigTrackingUpdate) check() error {
	if gtu.mutation.PublisherCleared() && len(gtu.mutation.PublisherIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "GigTracking.publisher"`)
	}
	return nil
}

func (gtu *GigTrackingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gtu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(gigtracking.Table, gigtracking.Columns, sqlgraph.NewFieldSpec(gigtracking.FieldID, field.TypeInt64))
	if ps := gtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gtu.mutation.Date(); ok {
		_spec.SetField(gigtracking.FieldDate, field.TypeTime, value)
	}
	if value, ok := gtu.mutation.AffiliateUserID(); ok {
		_spec.SetField(gigtracking.FieldAffiliateUserID, field.TypeString, value)
	}
	if gtu.mutation.AffiliateUserIDCleared() {
		_spec.ClearField(gigtracking.FieldAffiliateUserID, field.TypeString)
	}
	if value, ok := gtu.mutation.GetType(); ok {
		_spec.SetField(gigtracking.FieldType, field.TypeString, value)
	}
	if value, ok := gtu.mutation.UtmQuery(); ok {
		_spec.SetField(gigtracking.FieldUtmQuery, field.TypeString, value)
	}
	if gtu.mutation.UtmQueryCleared() {
		_spec.ClearField(gigtracking.FieldUtmQuery, field.TypeString)
	}
	if value, ok := gtu.mutation.Lp(); ok {
		_spec.SetField(gigtracking.FieldLp, field.TypeString, value)
	}
	if gtu.mutation.LpCleared() {
		_spec.ClearField(gigtracking.FieldLp, field.TypeString)
	}
	if value, ok := gtu.mutation.TrackID(); ok {
		_spec.SetField(gigtracking.FieldTrackID, field.TypeString, value)
	}
	if gtu.mutation.TrackIDCleared() {
		_spec.ClearField(gigtracking.FieldTrackID, field.TypeString)
	}
	if value, ok := gtu.mutation.Revenue(); ok {
		_spec.SetField(gigtracking.FieldRevenue, field.TypeFloat64, value)
	}
	if value, ok := gtu.mutation.AddedRevenue(); ok {
		_spec.AddField(gigtracking.FieldRevenue, field.TypeFloat64, value)
	}
	if value, ok := gtu.mutation.CreatedAt(); ok {
		_spec.SetField(gigtracking.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := gtu.mutation.UpdatedAt(); ok {
		_spec.SetField(gigtracking.FieldUpdatedAt, field.TypeTime, value)
	}
	if gtu.mutation.PublisherCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   gigtracking.PublisherTable,
			Columns: []string{gigtracking.PublisherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gtu.mutation.PublisherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   gigtracking.PublisherTable,
			Columns: []string{gigtracking.PublisherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gigtracking.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gtu.mutation.done = true
	return n, nil
}

// GigTrackingUpdateOne is the builder for updating a single GigTracking entity.
type GigTrackingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GigTrackingMutation
}

// SetDate sets the "date" field.
func (gtuo *GigTrackingUpdateOne) SetDate(t time.Time) *GigTrackingUpdateOne {
	gtuo.mutation.SetDate(t)
	return gtuo
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (gtuo *GigTrackingUpdateOne) SetNillableDate(t *time.Time) *GigTrackingUpdateOne {
	if t != nil {
		gtuo.SetDate(*t)
	}
	return gtuo
}

// SetAffiliateUserID sets the "affiliate_user_id" field.
func (gtuo *GigTrackingUpdateOne) SetAffiliateUserID(s string) *GigTrackingUpdateOne {
	gtuo.mutation.SetAffiliateUserID(s)
	return gtuo
}

// SetNillableAffiliateUserID sets the "affiliate_user_id" field if the given value is not nil.
func (gtuo *GigTrackingUpdateOne) SetNillableAffiliateUserID(s *string) *GigTrackingUpdateOne {
	if s != nil {
		gtuo.SetAffiliateUserID(*s)
	}
	return gtuo
}

// ClearAffiliateUserID clears the value of the "affiliate_user_id" field.
func (gtuo *GigTrackingUpdateOne) ClearAffiliateUserID() *GigTrackingUpdateOne {
	gtuo.mutation.ClearAffiliateUserID()
	return gtuo
}

// SetType sets the "type" field.
func (gtuo *GigTrackingUpdateOne) SetType(s string) *GigTrackingUpdateOne {
	gtuo.mutation.SetType(s)
	return gtuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (gtuo *GigTrackingUpdateOne) SetNillableType(s *string) *GigTrackingUpdateOne {
	if s != nil {
		gtuo.SetType(*s)
	}
	return gtuo
}

// SetUtmQuery sets the "utm_query" field.
func (gtuo *GigTrackingUpdateOne) SetUtmQuery(s string) *GigTrackingUpdateOne {
	gtuo.mutation.SetUtmQuery(s)
	return gtuo
}

// SetNillableUtmQuery sets the "utm_query" field if the given value is not nil.
func (gtuo *GigTrackingUpdateOne) SetNillableUtmQuery(s *string) *GigTrackingUpdateOne {
	if s != nil {
		gtuo.SetUtmQuery(*s)
	}
	return gtuo
}

// ClearUtmQuery clears the value of the "utm_query" field.
func (gtuo *GigTrackingUpdateOne) ClearUtmQuery() *GigTrackingUpdateOne {
	gtuo.mutation.ClearUtmQuery()
	return gtuo
}

// SetLp sets the "lp" field.
func (gtuo *GigTrackingUpdateOne) SetLp(s string) *GigTrackingUpdateOne {
	gtuo.mutation.SetLp(s)
	return gtuo
}

// SetNillableLp sets the "lp" field if the given value is not nil.
func (gtuo *GigTrackingUpdateOne) SetNillableLp(s *string) *GigTrackingUpdateOne {
	if s != nil {
		gtuo.SetLp(*s)
	}
	return gtuo
}

// ClearLp clears the value of the "lp" field.
func (gtuo *GigTrackingUpdateOne) ClearLp() *GigTrackingUpdateOne {
	gtuo.mutation.ClearLp()
	return gtuo
}

// SetTrackID sets the "track_id" field.
func (gtuo *GigTrackingUpdateOne) SetTrackID(s string) *GigTrackingUpdateOne {
	gtuo.mutation.SetTrackID(s)
	return gtuo
}

// SetNillableTrackID sets the "track_id" field if the given value is not nil.
func (gtuo *GigTrackingUpdateOne) SetNillableTrackID(s *string) *GigTrackingUpdateOne {
	if s != nil {
		gtuo.SetTrackID(*s)
	}
	return gtuo
}

// ClearTrackID clears the value of the "track_id" field.
func (gtuo *GigTrackingUpdateOne) ClearTrackID() *GigTrackingUpdateOne {
	gtuo.mutation.ClearTrackID()
	return gtuo
}

// SetRevenue sets the "revenue" field.
func (gtuo *GigTrackingUpdateOne) SetRevenue(f float64) *GigTrackingUpdateOne {
	gtuo.mutation.ResetRevenue()
	gtuo.mutation.SetRevenue(f)
	return gtuo
}

// SetNillableRevenue sets the "revenue" field if the given value is not nil.
func (gtuo *GigTrackingUpdateOne) SetNillableRevenue(f *float64) *GigTrackingUpdateOne {
	if f != nil {
		gtuo.SetRevenue(*f)
	}
	return gtuo
}

// AddRevenue adds f to the "revenue" field.
func (gtuo *GigTrackingUpdateOne) AddRevenue(f float64) *GigTrackingUpdateOne {
	gtuo.mutation.AddRevenue(f)
	return gtuo
}

// SetCreatedAt sets the "created_at" field.
func (gtuo *GigTrackingUpdateOne) SetCreatedAt(t time.Time) *GigTrackingUpdateOne {
	gtuo.mutation.SetCreatedAt(t)
	return gtuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gtuo *GigTrackingUpdateOne) SetNillableCreatedAt(t *time.Time) *GigTrackingUpdateOne {
	if t != nil {
		gtuo.SetCreatedAt(*t)
	}
	return gtuo
}

// SetUpdatedAt sets the "updated_at" field.
func (gtuo *GigTrackingUpdateOne) SetUpdatedAt(t time.Time) *GigTrackingUpdateOne {
	gtuo.mutation.SetUpdatedAt(t)
	return gtuo
}

// SetPublisherID sets the "publisher" edge to the User entity by ID.
func (gtuo *GigTrackingUpdateOne) SetPublisherID(id int64) *GigTrackingUpdateOne {
	gtuo.mutation.SetPublisherID(id)
	return gtuo
}

// SetPublisher sets the "publisher" edge to the User entity.
func (gtuo *GigTrackingUpdateOne) SetPublisher(u *User) *GigTrackingUpdateOne {
	return gtuo.SetPublisherID(u.ID)
}

// Mutation returns the GigTrackingMutation object of the builder.
func (gtuo *GigTrackingUpdateOne) Mutation() *GigTrackingMutation {
	return gtuo.mutation
}

// ClearPublisher clears the "publisher" edge to the User entity.
func (gtuo *GigTrackingUpdateOne) ClearPublisher() *GigTrackingUpdateOne {
	gtuo.mutation.ClearPublisher()
	return gtuo
}

// Where appends a list predicates to the GigTrackingUpdate builder.
func (gtuo *GigTrackingUpdateOne) Where(ps ...predicate.GigTracking) *GigTrackingUpdateOne {
	gtuo.mutation.Where(ps...)
	return gtuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gtuo *GigTrackingUpdateOne) Select(field string, fields ...string) *GigTrackingUpdateOne {
	gtuo.fields = append([]string{field}, fields...)
	return gtuo
}

// Save executes the query and returns the updated GigTracking entity.
func (gtuo *GigTrackingUpdateOne) Save(ctx context.Context) (*GigTracking, error) {
	gtuo.defaults()
	return withHooks(ctx, gtuo.sqlSave, gtuo.mutation, gtuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gtuo *GigTrackingUpdateOne) SaveX(ctx context.Context) *GigTracking {
	node, err := gtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gtuo *GigTrackingUpdateOne) Exec(ctx context.Context) error {
	_, err := gtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gtuo *GigTrackingUpdateOne) ExecX(ctx context.Context) {
	if err := gtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gtuo *GigTrackingUpdateOne) defaults() {
	if _, ok := gtuo.mutation.UpdatedAt(); !ok {
		v := gigtracking.UpdateDefaultUpdatedAt()
		gtuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gtuo *GigTrackingUpdateOne) check() error {
	if gtuo.mutation.PublisherCleared() && len(gtuo.mutation.PublisherIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "GigTracking.publisher"`)
	}
	return nil
}

func (gtuo *GigTrackingUpdateOne) sqlSave(ctx context.Context) (_node *GigTracking, err error) {
	if err := gtuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(gigtracking.Table, gigtracking.Columns, sqlgraph.NewFieldSpec(gigtracking.FieldID, field.TypeInt64))
	id, ok := gtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GigTracking.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gigtracking.FieldID)
		for _, f := range fields {
			if !gigtracking.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != gigtracking.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gtuo.mutation.Date(); ok {
		_spec.SetField(gigtracking.FieldDate, field.TypeTime, value)
	}
	if value, ok := gtuo.mutation.AffiliateUserID(); ok {
		_spec.SetField(gigtracking.FieldAffiliateUserID, field.TypeString, value)
	}
	if gtuo.mutation.AffiliateUserIDCleared() {
		_spec.ClearField(gigtracking.FieldAffiliateUserID, field.TypeString)
	}
	if value, ok := gtuo.mutation.GetType(); ok {
		_spec.SetField(gigtracking.FieldType, field.TypeString, value)
	}
	if value, ok := gtuo.mutation.UtmQuery(); ok {
		_spec.SetField(gigtracking.FieldUtmQuery, field.TypeString, value)
	}
	if gtuo.mutation.UtmQueryCleared() {
		_spec.ClearField(gigtracking.FieldUtmQuery, field.TypeString)
	}
	if value, ok := gtuo.mutation.Lp(); ok {
		_spec.SetField(gigtracking.FieldLp, field.TypeString, value)
	}
	if gtuo.mutation.LpCleared() {
		_spec.ClearField(gigtracking.FieldLp, field.TypeString)
	}
	if value, ok := gtuo.mutation.TrackID(); ok {
		_spec.SetField(gigtracking.FieldTrackID, field.TypeString, value)
	}
	if gtuo.mutation.TrackIDCleared() {
		_spec.ClearField(gigtracking.FieldTrackID, field.TypeString)
	}
	if value, ok := gtuo.mutation.Revenue(); ok {
		_spec.SetField(gigtracking.FieldRevenue, field.TypeFloat64, value)
	}
	if value, ok := gtuo.mutation.AddedRevenue(); ok {
		_spec.AddField(gigtracking.FieldRevenue, field.TypeFloat64, value)
	}
	if value, ok := gtuo.mutation.CreatedAt(); ok {
		_spec.SetField(gigtracking.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := gtuo.mutation.UpdatedAt(); ok {
		_spec.SetField(gigtracking.FieldUpdatedAt, field.TypeTime, value)
	}
	if gtuo.mutation.PublisherCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   gigtracking.PublisherTable,
			Columns: []string{gigtracking.PublisherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gtuo.mutation.PublisherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   gigtracking.PublisherTable,
			Columns: []string{gigtracking.PublisherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &GigTracking{config: gtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gigtracking.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gtuo.mutation.done = true
	return _node, nil
}
