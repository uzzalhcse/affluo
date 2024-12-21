// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/banner"
	"affluo/ent/lead"
	"affluo/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LeadUpdate is the builder for updating Lead entities.
type LeadUpdate struct {
	config
	hooks    []Hook
	mutation *LeadMutation
}

// Where appends a list predicates to the LeadUpdate builder.
func (lu *LeadUpdate) Where(ps ...predicate.Lead) *LeadUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetReferenceID sets the "reference_id" field.
func (lu *LeadUpdate) SetReferenceID(s string) *LeadUpdate {
	lu.mutation.SetReferenceID(s)
	return lu
}

// SetNillableReferenceID sets the "reference_id" field if the given value is not nil.
func (lu *LeadUpdate) SetNillableReferenceID(s *string) *LeadUpdate {
	if s != nil {
		lu.SetReferenceID(*s)
	}
	return lu
}

// ClearReferenceID clears the value of the "reference_id" field.
func (lu *LeadUpdate) ClearReferenceID() *LeadUpdate {
	lu.mutation.ClearReferenceID()
	return lu
}

// SetType sets the "type" field.
func (lu *LeadUpdate) SetType(l lead.Type) *LeadUpdate {
	lu.mutation.SetType(l)
	return lu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (lu *LeadUpdate) SetNillableType(l *lead.Type) *LeadUpdate {
	if l != nil {
		lu.SetType(*l)
	}
	return lu
}

// SetAmount sets the "amount" field.
func (lu *LeadUpdate) SetAmount(f float64) *LeadUpdate {
	lu.mutation.ResetAmount()
	lu.mutation.SetAmount(f)
	return lu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (lu *LeadUpdate) SetNillableAmount(f *float64) *LeadUpdate {
	if f != nil {
		lu.SetAmount(*f)
	}
	return lu
}

// AddAmount adds f to the "amount" field.
func (lu *LeadUpdate) AddAmount(f float64) *LeadUpdate {
	lu.mutation.AddAmount(f)
	return lu
}

// ClearAmount clears the value of the "amount" field.
func (lu *LeadUpdate) ClearAmount() *LeadUpdate {
	lu.mutation.ClearAmount()
	return lu
}

// SetCurrency sets the "currency" field.
func (lu *LeadUpdate) SetCurrency(s string) *LeadUpdate {
	lu.mutation.SetCurrency(s)
	return lu
}

// SetNillableCurrency sets the "currency" field if the given value is not nil.
func (lu *LeadUpdate) SetNillableCurrency(s *string) *LeadUpdate {
	if s != nil {
		lu.SetCurrency(*s)
	}
	return lu
}

// SetIPAddress sets the "ip_address" field.
func (lu *LeadUpdate) SetIPAddress(s string) *LeadUpdate {
	lu.mutation.SetIPAddress(s)
	return lu
}

// SetNillableIPAddress sets the "ip_address" field if the given value is not nil.
func (lu *LeadUpdate) SetNillableIPAddress(s *string) *LeadUpdate {
	if s != nil {
		lu.SetIPAddress(*s)
	}
	return lu
}

// ClearIPAddress clears the value of the "ip_address" field.
func (lu *LeadUpdate) ClearIPAddress() *LeadUpdate {
	lu.mutation.ClearIPAddress()
	return lu
}

// SetUserAgent sets the "user_agent" field.
func (lu *LeadUpdate) SetUserAgent(s string) *LeadUpdate {
	lu.mutation.SetUserAgent(s)
	return lu
}

// SetNillableUserAgent sets the "user_agent" field if the given value is not nil.
func (lu *LeadUpdate) SetNillableUserAgent(s *string) *LeadUpdate {
	if s != nil {
		lu.SetUserAgent(*s)
	}
	return lu
}

// ClearUserAgent clears the value of the "user_agent" field.
func (lu *LeadUpdate) ClearUserAgent() *LeadUpdate {
	lu.mutation.ClearUserAgent()
	return lu
}

// SetMetadata sets the "metadata" field.
func (lu *LeadUpdate) SetMetadata(m map[string]interface{}) *LeadUpdate {
	lu.mutation.SetMetadata(m)
	return lu
}

// ClearMetadata clears the value of the "metadata" field.
func (lu *LeadUpdate) ClearMetadata() *LeadUpdate {
	lu.mutation.ClearMetadata()
	return lu
}

// SetCreatedAt sets the "created_at" field.
func (lu *LeadUpdate) SetCreatedAt(t time.Time) *LeadUpdate {
	lu.mutation.SetCreatedAt(t)
	return lu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lu *LeadUpdate) SetNillableCreatedAt(t *time.Time) *LeadUpdate {
	if t != nil {
		lu.SetCreatedAt(*t)
	}
	return lu
}

// SetBannerID sets the "banner" edge to the Banner entity by ID.
func (lu *LeadUpdate) SetBannerID(id int64) *LeadUpdate {
	lu.mutation.SetBannerID(id)
	return lu
}

// SetNillableBannerID sets the "banner" edge to the Banner entity by ID if the given value is not nil.
func (lu *LeadUpdate) SetNillableBannerID(id *int64) *LeadUpdate {
	if id != nil {
		lu = lu.SetBannerID(*id)
	}
	return lu
}

// SetBanner sets the "banner" edge to the Banner entity.
func (lu *LeadUpdate) SetBanner(b *Banner) *LeadUpdate {
	return lu.SetBannerID(b.ID)
}

// Mutation returns the LeadMutation object of the builder.
func (lu *LeadUpdate) Mutation() *LeadMutation {
	return lu.mutation
}

// ClearBanner clears the "banner" edge to the Banner entity.
func (lu *LeadUpdate) ClearBanner() *LeadUpdate {
	lu.mutation.ClearBanner()
	return lu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LeadUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LeadUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LeadUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LeadUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *LeadUpdate) check() error {
	if v, ok := lu.mutation.GetType(); ok {
		if err := lead.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Lead.type": %w`, err)}
		}
	}
	return nil
}

func (lu *LeadUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(lead.Table, lead.Columns, sqlgraph.NewFieldSpec(lead.FieldID, field.TypeInt64))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.ReferenceID(); ok {
		_spec.SetField(lead.FieldReferenceID, field.TypeString, value)
	}
	if lu.mutation.ReferenceIDCleared() {
		_spec.ClearField(lead.FieldReferenceID, field.TypeString)
	}
	if value, ok := lu.mutation.GetType(); ok {
		_spec.SetField(lead.FieldType, field.TypeEnum, value)
	}
	if value, ok := lu.mutation.Amount(); ok {
		_spec.SetField(lead.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := lu.mutation.AddedAmount(); ok {
		_spec.AddField(lead.FieldAmount, field.TypeFloat64, value)
	}
	if lu.mutation.AmountCleared() {
		_spec.ClearField(lead.FieldAmount, field.TypeFloat64)
	}
	if value, ok := lu.mutation.Currency(); ok {
		_spec.SetField(lead.FieldCurrency, field.TypeString, value)
	}
	if value, ok := lu.mutation.IPAddress(); ok {
		_spec.SetField(lead.FieldIPAddress, field.TypeString, value)
	}
	if lu.mutation.IPAddressCleared() {
		_spec.ClearField(lead.FieldIPAddress, field.TypeString)
	}
	if value, ok := lu.mutation.UserAgent(); ok {
		_spec.SetField(lead.FieldUserAgent, field.TypeString, value)
	}
	if lu.mutation.UserAgentCleared() {
		_spec.ClearField(lead.FieldUserAgent, field.TypeString)
	}
	if value, ok := lu.mutation.Metadata(); ok {
		_spec.SetField(lead.FieldMetadata, field.TypeJSON, value)
	}
	if lu.mutation.MetadataCleared() {
		_spec.ClearField(lead.FieldMetadata, field.TypeJSON)
	}
	if value, ok := lu.mutation.CreatedAt(); ok {
		_spec.SetField(lead.FieldCreatedAt, field.TypeTime, value)
	}
	if lu.mutation.BannerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lead.BannerTable,
			Columns: []string{lead.BannerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(banner.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.BannerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lead.BannerTable,
			Columns: []string{lead.BannerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(banner.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lead.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LeadUpdateOne is the builder for updating a single Lead entity.
type LeadUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LeadMutation
}

// SetReferenceID sets the "reference_id" field.
func (luo *LeadUpdateOne) SetReferenceID(s string) *LeadUpdateOne {
	luo.mutation.SetReferenceID(s)
	return luo
}

// SetNillableReferenceID sets the "reference_id" field if the given value is not nil.
func (luo *LeadUpdateOne) SetNillableReferenceID(s *string) *LeadUpdateOne {
	if s != nil {
		luo.SetReferenceID(*s)
	}
	return luo
}

// ClearReferenceID clears the value of the "reference_id" field.
func (luo *LeadUpdateOne) ClearReferenceID() *LeadUpdateOne {
	luo.mutation.ClearReferenceID()
	return luo
}

// SetType sets the "type" field.
func (luo *LeadUpdateOne) SetType(l lead.Type) *LeadUpdateOne {
	luo.mutation.SetType(l)
	return luo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (luo *LeadUpdateOne) SetNillableType(l *lead.Type) *LeadUpdateOne {
	if l != nil {
		luo.SetType(*l)
	}
	return luo
}

// SetAmount sets the "amount" field.
func (luo *LeadUpdateOne) SetAmount(f float64) *LeadUpdateOne {
	luo.mutation.ResetAmount()
	luo.mutation.SetAmount(f)
	return luo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (luo *LeadUpdateOne) SetNillableAmount(f *float64) *LeadUpdateOne {
	if f != nil {
		luo.SetAmount(*f)
	}
	return luo
}

// AddAmount adds f to the "amount" field.
func (luo *LeadUpdateOne) AddAmount(f float64) *LeadUpdateOne {
	luo.mutation.AddAmount(f)
	return luo
}

// ClearAmount clears the value of the "amount" field.
func (luo *LeadUpdateOne) ClearAmount() *LeadUpdateOne {
	luo.mutation.ClearAmount()
	return luo
}

// SetCurrency sets the "currency" field.
func (luo *LeadUpdateOne) SetCurrency(s string) *LeadUpdateOne {
	luo.mutation.SetCurrency(s)
	return luo
}

// SetNillableCurrency sets the "currency" field if the given value is not nil.
func (luo *LeadUpdateOne) SetNillableCurrency(s *string) *LeadUpdateOne {
	if s != nil {
		luo.SetCurrency(*s)
	}
	return luo
}

// SetIPAddress sets the "ip_address" field.
func (luo *LeadUpdateOne) SetIPAddress(s string) *LeadUpdateOne {
	luo.mutation.SetIPAddress(s)
	return luo
}

// SetNillableIPAddress sets the "ip_address" field if the given value is not nil.
func (luo *LeadUpdateOne) SetNillableIPAddress(s *string) *LeadUpdateOne {
	if s != nil {
		luo.SetIPAddress(*s)
	}
	return luo
}

// ClearIPAddress clears the value of the "ip_address" field.
func (luo *LeadUpdateOne) ClearIPAddress() *LeadUpdateOne {
	luo.mutation.ClearIPAddress()
	return luo
}

// SetUserAgent sets the "user_agent" field.
func (luo *LeadUpdateOne) SetUserAgent(s string) *LeadUpdateOne {
	luo.mutation.SetUserAgent(s)
	return luo
}

// SetNillableUserAgent sets the "user_agent" field if the given value is not nil.
func (luo *LeadUpdateOne) SetNillableUserAgent(s *string) *LeadUpdateOne {
	if s != nil {
		luo.SetUserAgent(*s)
	}
	return luo
}

// ClearUserAgent clears the value of the "user_agent" field.
func (luo *LeadUpdateOne) ClearUserAgent() *LeadUpdateOne {
	luo.mutation.ClearUserAgent()
	return luo
}

// SetMetadata sets the "metadata" field.
func (luo *LeadUpdateOne) SetMetadata(m map[string]interface{}) *LeadUpdateOne {
	luo.mutation.SetMetadata(m)
	return luo
}

// ClearMetadata clears the value of the "metadata" field.
func (luo *LeadUpdateOne) ClearMetadata() *LeadUpdateOne {
	luo.mutation.ClearMetadata()
	return luo
}

// SetCreatedAt sets the "created_at" field.
func (luo *LeadUpdateOne) SetCreatedAt(t time.Time) *LeadUpdateOne {
	luo.mutation.SetCreatedAt(t)
	return luo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (luo *LeadUpdateOne) SetNillableCreatedAt(t *time.Time) *LeadUpdateOne {
	if t != nil {
		luo.SetCreatedAt(*t)
	}
	return luo
}

// SetBannerID sets the "banner" edge to the Banner entity by ID.
func (luo *LeadUpdateOne) SetBannerID(id int64) *LeadUpdateOne {
	luo.mutation.SetBannerID(id)
	return luo
}

// SetNillableBannerID sets the "banner" edge to the Banner entity by ID if the given value is not nil.
func (luo *LeadUpdateOne) SetNillableBannerID(id *int64) *LeadUpdateOne {
	if id != nil {
		luo = luo.SetBannerID(*id)
	}
	return luo
}

// SetBanner sets the "banner" edge to the Banner entity.
func (luo *LeadUpdateOne) SetBanner(b *Banner) *LeadUpdateOne {
	return luo.SetBannerID(b.ID)
}

// Mutation returns the LeadMutation object of the builder.
func (luo *LeadUpdateOne) Mutation() *LeadMutation {
	return luo.mutation
}

// ClearBanner clears the "banner" edge to the Banner entity.
func (luo *LeadUpdateOne) ClearBanner() *LeadUpdateOne {
	luo.mutation.ClearBanner()
	return luo
}

// Where appends a list predicates to the LeadUpdate builder.
func (luo *LeadUpdateOne) Where(ps ...predicate.Lead) *LeadUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LeadUpdateOne) Select(field string, fields ...string) *LeadUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Lead entity.
func (luo *LeadUpdateOne) Save(ctx context.Context) (*Lead, error) {
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LeadUpdateOne) SaveX(ctx context.Context) *Lead {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LeadUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LeadUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *LeadUpdateOne) check() error {
	if v, ok := luo.mutation.GetType(); ok {
		if err := lead.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Lead.type": %w`, err)}
		}
	}
	return nil
}

func (luo *LeadUpdateOne) sqlSave(ctx context.Context) (_node *Lead, err error) {
	if err := luo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(lead.Table, lead.Columns, sqlgraph.NewFieldSpec(lead.FieldID, field.TypeInt64))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Lead.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lead.FieldID)
		for _, f := range fields {
			if !lead.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != lead.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.ReferenceID(); ok {
		_spec.SetField(lead.FieldReferenceID, field.TypeString, value)
	}
	if luo.mutation.ReferenceIDCleared() {
		_spec.ClearField(lead.FieldReferenceID, field.TypeString)
	}
	if value, ok := luo.mutation.GetType(); ok {
		_spec.SetField(lead.FieldType, field.TypeEnum, value)
	}
	if value, ok := luo.mutation.Amount(); ok {
		_spec.SetField(lead.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := luo.mutation.AddedAmount(); ok {
		_spec.AddField(lead.FieldAmount, field.TypeFloat64, value)
	}
	if luo.mutation.AmountCleared() {
		_spec.ClearField(lead.FieldAmount, field.TypeFloat64)
	}
	if value, ok := luo.mutation.Currency(); ok {
		_spec.SetField(lead.FieldCurrency, field.TypeString, value)
	}
	if value, ok := luo.mutation.IPAddress(); ok {
		_spec.SetField(lead.FieldIPAddress, field.TypeString, value)
	}
	if luo.mutation.IPAddressCleared() {
		_spec.ClearField(lead.FieldIPAddress, field.TypeString)
	}
	if value, ok := luo.mutation.UserAgent(); ok {
		_spec.SetField(lead.FieldUserAgent, field.TypeString, value)
	}
	if luo.mutation.UserAgentCleared() {
		_spec.ClearField(lead.FieldUserAgent, field.TypeString)
	}
	if value, ok := luo.mutation.Metadata(); ok {
		_spec.SetField(lead.FieldMetadata, field.TypeJSON, value)
	}
	if luo.mutation.MetadataCleared() {
		_spec.ClearField(lead.FieldMetadata, field.TypeJSON)
	}
	if value, ok := luo.mutation.CreatedAt(); ok {
		_spec.SetField(lead.FieldCreatedAt, field.TypeTime, value)
	}
	if luo.mutation.BannerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lead.BannerTable,
			Columns: []string{lead.BannerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(banner.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.BannerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lead.BannerTable,
			Columns: []string{lead.BannerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(banner.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Lead{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lead.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}