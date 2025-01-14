// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/commissionplan"
	"affluo/ent/predicate"
	"affluo/ent/user"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommissionPlanUpdate is the builder for updating CommissionPlan entities.
type CommissionPlanUpdate struct {
	config
	hooks    []Hook
	mutation *CommissionPlanMutation
}

// Where appends a list predicates to the CommissionPlanUpdate builder.
func (cpu *CommissionPlanUpdate) Where(ps ...predicate.CommissionPlan) *CommissionPlanUpdate {
	cpu.mutation.Where(ps...)
	return cpu
}

// SetName sets the "name" field.
func (cpu *CommissionPlanUpdate) SetName(s string) *CommissionPlanUpdate {
	cpu.mutation.SetName(s)
	return cpu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cpu *CommissionPlanUpdate) SetNillableName(s *string) *CommissionPlanUpdate {
	if s != nil {
		cpu.SetName(*s)
	}
	return cpu
}

// SetDescription sets the "description" field.
func (cpu *CommissionPlanUpdate) SetDescription(s string) *CommissionPlanUpdate {
	cpu.mutation.SetDescription(s)
	return cpu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cpu *CommissionPlanUpdate) SetNillableDescription(s *string) *CommissionPlanUpdate {
	if s != nil {
		cpu.SetDescription(*s)
	}
	return cpu
}

// ClearDescription clears the value of the "description" field.
func (cpu *CommissionPlanUpdate) ClearDescription() *CommissionPlanUpdate {
	cpu.mutation.ClearDescription()
	return cpu
}

// SetType sets the "type" field.
func (cpu *CommissionPlanUpdate) SetType(c commissionplan.Type) *CommissionPlanUpdate {
	cpu.mutation.SetType(c)
	return cpu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cpu *CommissionPlanUpdate) SetNillableType(c *commissionplan.Type) *CommissionPlanUpdate {
	if c != nil {
		cpu.SetType(*c)
	}
	return cpu
}

// SetClickCommission sets the "click_commission" field.
func (cpu *CommissionPlanUpdate) SetClickCommission(f float64) *CommissionPlanUpdate {
	cpu.mutation.ResetClickCommission()
	cpu.mutation.SetClickCommission(f)
	return cpu
}

// SetNillableClickCommission sets the "click_commission" field if the given value is not nil.
func (cpu *CommissionPlanUpdate) SetNillableClickCommission(f *float64) *CommissionPlanUpdate {
	if f != nil {
		cpu.SetClickCommission(*f)
	}
	return cpu
}

// AddClickCommission adds f to the "click_commission" field.
func (cpu *CommissionPlanUpdate) AddClickCommission(f float64) *CommissionPlanUpdate {
	cpu.mutation.AddClickCommission(f)
	return cpu
}

// SetImpressionCommission sets the "impression_commission" field.
func (cpu *CommissionPlanUpdate) SetImpressionCommission(f float64) *CommissionPlanUpdate {
	cpu.mutation.ResetImpressionCommission()
	cpu.mutation.SetImpressionCommission(f)
	return cpu
}

// SetNillableImpressionCommission sets the "impression_commission" field if the given value is not nil.
func (cpu *CommissionPlanUpdate) SetNillableImpressionCommission(f *float64) *CommissionPlanUpdate {
	if f != nil {
		cpu.SetImpressionCommission(*f)
	}
	return cpu
}

// AddImpressionCommission adds f to the "impression_commission" field.
func (cpu *CommissionPlanUpdate) AddImpressionCommission(f float64) *CommissionPlanUpdate {
	cpu.mutation.AddImpressionCommission(f)
	return cpu
}

// SetFirstLeadCommission sets the "first_lead_commission" field.
func (cpu *CommissionPlanUpdate) SetFirstLeadCommission(f float64) *CommissionPlanUpdate {
	cpu.mutation.ResetFirstLeadCommission()
	cpu.mutation.SetFirstLeadCommission(f)
	return cpu
}

// SetNillableFirstLeadCommission sets the "first_lead_commission" field if the given value is not nil.
func (cpu *CommissionPlanUpdate) SetNillableFirstLeadCommission(f *float64) *CommissionPlanUpdate {
	if f != nil {
		cpu.SetFirstLeadCommission(*f)
	}
	return cpu
}

// AddFirstLeadCommission adds f to the "first_lead_commission" field.
func (cpu *CommissionPlanUpdate) AddFirstLeadCommission(f float64) *CommissionPlanUpdate {
	cpu.mutation.AddFirstLeadCommission(f)
	return cpu
}

// SetRepeatLeadCommission sets the "repeat_lead_commission" field.
func (cpu *CommissionPlanUpdate) SetRepeatLeadCommission(f float64) *CommissionPlanUpdate {
	cpu.mutation.ResetRepeatLeadCommission()
	cpu.mutation.SetRepeatLeadCommission(f)
	return cpu
}

// SetNillableRepeatLeadCommission sets the "repeat_lead_commission" field if the given value is not nil.
func (cpu *CommissionPlanUpdate) SetNillableRepeatLeadCommission(f *float64) *CommissionPlanUpdate {
	if f != nil {
		cpu.SetRepeatLeadCommission(*f)
	}
	return cpu
}

// AddRepeatLeadCommission adds f to the "repeat_lead_commission" field.
func (cpu *CommissionPlanUpdate) AddRepeatLeadCommission(f float64) *CommissionPlanUpdate {
	cpu.mutation.AddRepeatLeadCommission(f)
	return cpu
}

// SetValidMonths sets the "valid_months" field.
func (cpu *CommissionPlanUpdate) SetValidMonths(i int) *CommissionPlanUpdate {
	cpu.mutation.ResetValidMonths()
	cpu.mutation.SetValidMonths(i)
	return cpu
}

// SetNillableValidMonths sets the "valid_months" field if the given value is not nil.
func (cpu *CommissionPlanUpdate) SetNillableValidMonths(i *int) *CommissionPlanUpdate {
	if i != nil {
		cpu.SetValidMonths(*i)
	}
	return cpu
}

// AddValidMonths adds i to the "valid_months" field.
func (cpu *CommissionPlanUpdate) AddValidMonths(i int) *CommissionPlanUpdate {
	cpu.mutation.AddValidMonths(i)
	return cpu
}

// SetMinimumPayout sets the "minimum_payout" field.
func (cpu *CommissionPlanUpdate) SetMinimumPayout(f float64) *CommissionPlanUpdate {
	cpu.mutation.ResetMinimumPayout()
	cpu.mutation.SetMinimumPayout(f)
	return cpu
}

// SetNillableMinimumPayout sets the "minimum_payout" field if the given value is not nil.
func (cpu *CommissionPlanUpdate) SetNillableMinimumPayout(f *float64) *CommissionPlanUpdate {
	if f != nil {
		cpu.SetMinimumPayout(*f)
	}
	return cpu
}

// AddMinimumPayout adds f to the "minimum_payout" field.
func (cpu *CommissionPlanUpdate) AddMinimumPayout(f float64) *CommissionPlanUpdate {
	cpu.mutation.AddMinimumPayout(f)
	return cpu
}

// SetIsActive sets the "is_active" field.
func (cpu *CommissionPlanUpdate) SetIsActive(b bool) *CommissionPlanUpdate {
	cpu.mutation.SetIsActive(b)
	return cpu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (cpu *CommissionPlanUpdate) SetNillableIsActive(b *bool) *CommissionPlanUpdate {
	if b != nil {
		cpu.SetIsActive(*b)
	}
	return cpu
}

// SetIsDefault sets the "is_default" field.
func (cpu *CommissionPlanUpdate) SetIsDefault(b bool) *CommissionPlanUpdate {
	cpu.mutation.SetIsDefault(b)
	return cpu
}

// SetNillableIsDefault sets the "is_default" field if the given value is not nil.
func (cpu *CommissionPlanUpdate) SetNillableIsDefault(b *bool) *CommissionPlanUpdate {
	if b != nil {
		cpu.SetIsDefault(*b)
	}
	return cpu
}

// AddPublisherIDs adds the "publishers" edge to the User entity by IDs.
func (cpu *CommissionPlanUpdate) AddPublisherIDs(ids ...int64) *CommissionPlanUpdate {
	cpu.mutation.AddPublisherIDs(ids...)
	return cpu
}

// AddPublishers adds the "publishers" edges to the User entity.
func (cpu *CommissionPlanUpdate) AddPublishers(u ...*User) *CommissionPlanUpdate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cpu.AddPublisherIDs(ids...)
}

// Mutation returns the CommissionPlanMutation object of the builder.
func (cpu *CommissionPlanUpdate) Mutation() *CommissionPlanMutation {
	return cpu.mutation
}

// ClearPublishers clears all "publishers" edges to the User entity.
func (cpu *CommissionPlanUpdate) ClearPublishers() *CommissionPlanUpdate {
	cpu.mutation.ClearPublishers()
	return cpu
}

// RemovePublisherIDs removes the "publishers" edge to User entities by IDs.
func (cpu *CommissionPlanUpdate) RemovePublisherIDs(ids ...int64) *CommissionPlanUpdate {
	cpu.mutation.RemovePublisherIDs(ids...)
	return cpu
}

// RemovePublishers removes "publishers" edges to User entities.
func (cpu *CommissionPlanUpdate) RemovePublishers(u ...*User) *CommissionPlanUpdate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cpu.RemovePublisherIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cpu *CommissionPlanUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cpu.sqlSave, cpu.mutation, cpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cpu *CommissionPlanUpdate) SaveX(ctx context.Context) int {
	affected, err := cpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cpu *CommissionPlanUpdate) Exec(ctx context.Context) error {
	_, err := cpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpu *CommissionPlanUpdate) ExecX(ctx context.Context) {
	if err := cpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cpu *CommissionPlanUpdate) check() error {
	if v, ok := cpu.mutation.Name(); ok {
		if err := commissionplan.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "CommissionPlan.name": %w`, err)}
		}
	}
	if v, ok := cpu.mutation.GetType(); ok {
		if err := commissionplan.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "CommissionPlan.type": %w`, err)}
		}
	}
	return nil
}

func (cpu *CommissionPlanUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cpu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(commissionplan.Table, commissionplan.Columns, sqlgraph.NewFieldSpec(commissionplan.FieldID, field.TypeInt))
	if ps := cpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cpu.mutation.Name(); ok {
		_spec.SetField(commissionplan.FieldName, field.TypeString, value)
	}
	if value, ok := cpu.mutation.Description(); ok {
		_spec.SetField(commissionplan.FieldDescription, field.TypeString, value)
	}
	if cpu.mutation.DescriptionCleared() {
		_spec.ClearField(commissionplan.FieldDescription, field.TypeString)
	}
	if value, ok := cpu.mutation.GetType(); ok {
		_spec.SetField(commissionplan.FieldType, field.TypeEnum, value)
	}
	if value, ok := cpu.mutation.ClickCommission(); ok {
		_spec.SetField(commissionplan.FieldClickCommission, field.TypeFloat64, value)
	}
	if value, ok := cpu.mutation.AddedClickCommission(); ok {
		_spec.AddField(commissionplan.FieldClickCommission, field.TypeFloat64, value)
	}
	if value, ok := cpu.mutation.ImpressionCommission(); ok {
		_spec.SetField(commissionplan.FieldImpressionCommission, field.TypeFloat64, value)
	}
	if value, ok := cpu.mutation.AddedImpressionCommission(); ok {
		_spec.AddField(commissionplan.FieldImpressionCommission, field.TypeFloat64, value)
	}
	if value, ok := cpu.mutation.FirstLeadCommission(); ok {
		_spec.SetField(commissionplan.FieldFirstLeadCommission, field.TypeFloat64, value)
	}
	if value, ok := cpu.mutation.AddedFirstLeadCommission(); ok {
		_spec.AddField(commissionplan.FieldFirstLeadCommission, field.TypeFloat64, value)
	}
	if value, ok := cpu.mutation.RepeatLeadCommission(); ok {
		_spec.SetField(commissionplan.FieldRepeatLeadCommission, field.TypeFloat64, value)
	}
	if value, ok := cpu.mutation.AddedRepeatLeadCommission(); ok {
		_spec.AddField(commissionplan.FieldRepeatLeadCommission, field.TypeFloat64, value)
	}
	if value, ok := cpu.mutation.ValidMonths(); ok {
		_spec.SetField(commissionplan.FieldValidMonths, field.TypeInt, value)
	}
	if value, ok := cpu.mutation.AddedValidMonths(); ok {
		_spec.AddField(commissionplan.FieldValidMonths, field.TypeInt, value)
	}
	if value, ok := cpu.mutation.MinimumPayout(); ok {
		_spec.SetField(commissionplan.FieldMinimumPayout, field.TypeFloat64, value)
	}
	if value, ok := cpu.mutation.AddedMinimumPayout(); ok {
		_spec.AddField(commissionplan.FieldMinimumPayout, field.TypeFloat64, value)
	}
	if value, ok := cpu.mutation.IsActive(); ok {
		_spec.SetField(commissionplan.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := cpu.mutation.IsDefault(); ok {
		_spec.SetField(commissionplan.FieldIsDefault, field.TypeBool, value)
	}
	if cpu.mutation.PublishersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   commissionplan.PublishersTable,
			Columns: []string{commissionplan.PublishersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpu.mutation.RemovedPublishersIDs(); len(nodes) > 0 && !cpu.mutation.PublishersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   commissionplan.PublishersTable,
			Columns: []string{commissionplan.PublishersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpu.mutation.PublishersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   commissionplan.PublishersTable,
			Columns: []string{commissionplan.PublishersColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, cpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{commissionplan.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cpu.mutation.done = true
	return n, nil
}

// CommissionPlanUpdateOne is the builder for updating a single CommissionPlan entity.
type CommissionPlanUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommissionPlanMutation
}

// SetName sets the "name" field.
func (cpuo *CommissionPlanUpdateOne) SetName(s string) *CommissionPlanUpdateOne {
	cpuo.mutation.SetName(s)
	return cpuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cpuo *CommissionPlanUpdateOne) SetNillableName(s *string) *CommissionPlanUpdateOne {
	if s != nil {
		cpuo.SetName(*s)
	}
	return cpuo
}

// SetDescription sets the "description" field.
func (cpuo *CommissionPlanUpdateOne) SetDescription(s string) *CommissionPlanUpdateOne {
	cpuo.mutation.SetDescription(s)
	return cpuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cpuo *CommissionPlanUpdateOne) SetNillableDescription(s *string) *CommissionPlanUpdateOne {
	if s != nil {
		cpuo.SetDescription(*s)
	}
	return cpuo
}

// ClearDescription clears the value of the "description" field.
func (cpuo *CommissionPlanUpdateOne) ClearDescription() *CommissionPlanUpdateOne {
	cpuo.mutation.ClearDescription()
	return cpuo
}

// SetType sets the "type" field.
func (cpuo *CommissionPlanUpdateOne) SetType(c commissionplan.Type) *CommissionPlanUpdateOne {
	cpuo.mutation.SetType(c)
	return cpuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cpuo *CommissionPlanUpdateOne) SetNillableType(c *commissionplan.Type) *CommissionPlanUpdateOne {
	if c != nil {
		cpuo.SetType(*c)
	}
	return cpuo
}

// SetClickCommission sets the "click_commission" field.
func (cpuo *CommissionPlanUpdateOne) SetClickCommission(f float64) *CommissionPlanUpdateOne {
	cpuo.mutation.ResetClickCommission()
	cpuo.mutation.SetClickCommission(f)
	return cpuo
}

// SetNillableClickCommission sets the "click_commission" field if the given value is not nil.
func (cpuo *CommissionPlanUpdateOne) SetNillableClickCommission(f *float64) *CommissionPlanUpdateOne {
	if f != nil {
		cpuo.SetClickCommission(*f)
	}
	return cpuo
}

// AddClickCommission adds f to the "click_commission" field.
func (cpuo *CommissionPlanUpdateOne) AddClickCommission(f float64) *CommissionPlanUpdateOne {
	cpuo.mutation.AddClickCommission(f)
	return cpuo
}

// SetImpressionCommission sets the "impression_commission" field.
func (cpuo *CommissionPlanUpdateOne) SetImpressionCommission(f float64) *CommissionPlanUpdateOne {
	cpuo.mutation.ResetImpressionCommission()
	cpuo.mutation.SetImpressionCommission(f)
	return cpuo
}

// SetNillableImpressionCommission sets the "impression_commission" field if the given value is not nil.
func (cpuo *CommissionPlanUpdateOne) SetNillableImpressionCommission(f *float64) *CommissionPlanUpdateOne {
	if f != nil {
		cpuo.SetImpressionCommission(*f)
	}
	return cpuo
}

// AddImpressionCommission adds f to the "impression_commission" field.
func (cpuo *CommissionPlanUpdateOne) AddImpressionCommission(f float64) *CommissionPlanUpdateOne {
	cpuo.mutation.AddImpressionCommission(f)
	return cpuo
}

// SetFirstLeadCommission sets the "first_lead_commission" field.
func (cpuo *CommissionPlanUpdateOne) SetFirstLeadCommission(f float64) *CommissionPlanUpdateOne {
	cpuo.mutation.ResetFirstLeadCommission()
	cpuo.mutation.SetFirstLeadCommission(f)
	return cpuo
}

// SetNillableFirstLeadCommission sets the "first_lead_commission" field if the given value is not nil.
func (cpuo *CommissionPlanUpdateOne) SetNillableFirstLeadCommission(f *float64) *CommissionPlanUpdateOne {
	if f != nil {
		cpuo.SetFirstLeadCommission(*f)
	}
	return cpuo
}

// AddFirstLeadCommission adds f to the "first_lead_commission" field.
func (cpuo *CommissionPlanUpdateOne) AddFirstLeadCommission(f float64) *CommissionPlanUpdateOne {
	cpuo.mutation.AddFirstLeadCommission(f)
	return cpuo
}

// SetRepeatLeadCommission sets the "repeat_lead_commission" field.
func (cpuo *CommissionPlanUpdateOne) SetRepeatLeadCommission(f float64) *CommissionPlanUpdateOne {
	cpuo.mutation.ResetRepeatLeadCommission()
	cpuo.mutation.SetRepeatLeadCommission(f)
	return cpuo
}

// SetNillableRepeatLeadCommission sets the "repeat_lead_commission" field if the given value is not nil.
func (cpuo *CommissionPlanUpdateOne) SetNillableRepeatLeadCommission(f *float64) *CommissionPlanUpdateOne {
	if f != nil {
		cpuo.SetRepeatLeadCommission(*f)
	}
	return cpuo
}

// AddRepeatLeadCommission adds f to the "repeat_lead_commission" field.
func (cpuo *CommissionPlanUpdateOne) AddRepeatLeadCommission(f float64) *CommissionPlanUpdateOne {
	cpuo.mutation.AddRepeatLeadCommission(f)
	return cpuo
}

// SetValidMonths sets the "valid_months" field.
func (cpuo *CommissionPlanUpdateOne) SetValidMonths(i int) *CommissionPlanUpdateOne {
	cpuo.mutation.ResetValidMonths()
	cpuo.mutation.SetValidMonths(i)
	return cpuo
}

// SetNillableValidMonths sets the "valid_months" field if the given value is not nil.
func (cpuo *CommissionPlanUpdateOne) SetNillableValidMonths(i *int) *CommissionPlanUpdateOne {
	if i != nil {
		cpuo.SetValidMonths(*i)
	}
	return cpuo
}

// AddValidMonths adds i to the "valid_months" field.
func (cpuo *CommissionPlanUpdateOne) AddValidMonths(i int) *CommissionPlanUpdateOne {
	cpuo.mutation.AddValidMonths(i)
	return cpuo
}

// SetMinimumPayout sets the "minimum_payout" field.
func (cpuo *CommissionPlanUpdateOne) SetMinimumPayout(f float64) *CommissionPlanUpdateOne {
	cpuo.mutation.ResetMinimumPayout()
	cpuo.mutation.SetMinimumPayout(f)
	return cpuo
}

// SetNillableMinimumPayout sets the "minimum_payout" field if the given value is not nil.
func (cpuo *CommissionPlanUpdateOne) SetNillableMinimumPayout(f *float64) *CommissionPlanUpdateOne {
	if f != nil {
		cpuo.SetMinimumPayout(*f)
	}
	return cpuo
}

// AddMinimumPayout adds f to the "minimum_payout" field.
func (cpuo *CommissionPlanUpdateOne) AddMinimumPayout(f float64) *CommissionPlanUpdateOne {
	cpuo.mutation.AddMinimumPayout(f)
	return cpuo
}

// SetIsActive sets the "is_active" field.
func (cpuo *CommissionPlanUpdateOne) SetIsActive(b bool) *CommissionPlanUpdateOne {
	cpuo.mutation.SetIsActive(b)
	return cpuo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (cpuo *CommissionPlanUpdateOne) SetNillableIsActive(b *bool) *CommissionPlanUpdateOne {
	if b != nil {
		cpuo.SetIsActive(*b)
	}
	return cpuo
}

// SetIsDefault sets the "is_default" field.
func (cpuo *CommissionPlanUpdateOne) SetIsDefault(b bool) *CommissionPlanUpdateOne {
	cpuo.mutation.SetIsDefault(b)
	return cpuo
}

// SetNillableIsDefault sets the "is_default" field if the given value is not nil.
func (cpuo *CommissionPlanUpdateOne) SetNillableIsDefault(b *bool) *CommissionPlanUpdateOne {
	if b != nil {
		cpuo.SetIsDefault(*b)
	}
	return cpuo
}

// AddPublisherIDs adds the "publishers" edge to the User entity by IDs.
func (cpuo *CommissionPlanUpdateOne) AddPublisherIDs(ids ...int64) *CommissionPlanUpdateOne {
	cpuo.mutation.AddPublisherIDs(ids...)
	return cpuo
}

// AddPublishers adds the "publishers" edges to the User entity.
func (cpuo *CommissionPlanUpdateOne) AddPublishers(u ...*User) *CommissionPlanUpdateOne {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cpuo.AddPublisherIDs(ids...)
}

// Mutation returns the CommissionPlanMutation object of the builder.
func (cpuo *CommissionPlanUpdateOne) Mutation() *CommissionPlanMutation {
	return cpuo.mutation
}

// ClearPublishers clears all "publishers" edges to the User entity.
func (cpuo *CommissionPlanUpdateOne) ClearPublishers() *CommissionPlanUpdateOne {
	cpuo.mutation.ClearPublishers()
	return cpuo
}

// RemovePublisherIDs removes the "publishers" edge to User entities by IDs.
func (cpuo *CommissionPlanUpdateOne) RemovePublisherIDs(ids ...int64) *CommissionPlanUpdateOne {
	cpuo.mutation.RemovePublisherIDs(ids...)
	return cpuo
}

// RemovePublishers removes "publishers" edges to User entities.
func (cpuo *CommissionPlanUpdateOne) RemovePublishers(u ...*User) *CommissionPlanUpdateOne {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cpuo.RemovePublisherIDs(ids...)
}

// Where appends a list predicates to the CommissionPlanUpdate builder.
func (cpuo *CommissionPlanUpdateOne) Where(ps ...predicate.CommissionPlan) *CommissionPlanUpdateOne {
	cpuo.mutation.Where(ps...)
	return cpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cpuo *CommissionPlanUpdateOne) Select(field string, fields ...string) *CommissionPlanUpdateOne {
	cpuo.fields = append([]string{field}, fields...)
	return cpuo
}

// Save executes the query and returns the updated CommissionPlan entity.
func (cpuo *CommissionPlanUpdateOne) Save(ctx context.Context) (*CommissionPlan, error) {
	return withHooks(ctx, cpuo.sqlSave, cpuo.mutation, cpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cpuo *CommissionPlanUpdateOne) SaveX(ctx context.Context) *CommissionPlan {
	node, err := cpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cpuo *CommissionPlanUpdateOne) Exec(ctx context.Context) error {
	_, err := cpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpuo *CommissionPlanUpdateOne) ExecX(ctx context.Context) {
	if err := cpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cpuo *CommissionPlanUpdateOne) check() error {
	if v, ok := cpuo.mutation.Name(); ok {
		if err := commissionplan.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "CommissionPlan.name": %w`, err)}
		}
	}
	if v, ok := cpuo.mutation.GetType(); ok {
		if err := commissionplan.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "CommissionPlan.type": %w`, err)}
		}
	}
	return nil
}

func (cpuo *CommissionPlanUpdateOne) sqlSave(ctx context.Context) (_node *CommissionPlan, err error) {
	if err := cpuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(commissionplan.Table, commissionplan.Columns, sqlgraph.NewFieldSpec(commissionplan.FieldID, field.TypeInt))
	id, ok := cpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CommissionPlan.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, commissionplan.FieldID)
		for _, f := range fields {
			if !commissionplan.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != commissionplan.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cpuo.mutation.Name(); ok {
		_spec.SetField(commissionplan.FieldName, field.TypeString, value)
	}
	if value, ok := cpuo.mutation.Description(); ok {
		_spec.SetField(commissionplan.FieldDescription, field.TypeString, value)
	}
	if cpuo.mutation.DescriptionCleared() {
		_spec.ClearField(commissionplan.FieldDescription, field.TypeString)
	}
	if value, ok := cpuo.mutation.GetType(); ok {
		_spec.SetField(commissionplan.FieldType, field.TypeEnum, value)
	}
	if value, ok := cpuo.mutation.ClickCommission(); ok {
		_spec.SetField(commissionplan.FieldClickCommission, field.TypeFloat64, value)
	}
	if value, ok := cpuo.mutation.AddedClickCommission(); ok {
		_spec.AddField(commissionplan.FieldClickCommission, field.TypeFloat64, value)
	}
	if value, ok := cpuo.mutation.ImpressionCommission(); ok {
		_spec.SetField(commissionplan.FieldImpressionCommission, field.TypeFloat64, value)
	}
	if value, ok := cpuo.mutation.AddedImpressionCommission(); ok {
		_spec.AddField(commissionplan.FieldImpressionCommission, field.TypeFloat64, value)
	}
	if value, ok := cpuo.mutation.FirstLeadCommission(); ok {
		_spec.SetField(commissionplan.FieldFirstLeadCommission, field.TypeFloat64, value)
	}
	if value, ok := cpuo.mutation.AddedFirstLeadCommission(); ok {
		_spec.AddField(commissionplan.FieldFirstLeadCommission, field.TypeFloat64, value)
	}
	if value, ok := cpuo.mutation.RepeatLeadCommission(); ok {
		_spec.SetField(commissionplan.FieldRepeatLeadCommission, field.TypeFloat64, value)
	}
	if value, ok := cpuo.mutation.AddedRepeatLeadCommission(); ok {
		_spec.AddField(commissionplan.FieldRepeatLeadCommission, field.TypeFloat64, value)
	}
	if value, ok := cpuo.mutation.ValidMonths(); ok {
		_spec.SetField(commissionplan.FieldValidMonths, field.TypeInt, value)
	}
	if value, ok := cpuo.mutation.AddedValidMonths(); ok {
		_spec.AddField(commissionplan.FieldValidMonths, field.TypeInt, value)
	}
	if value, ok := cpuo.mutation.MinimumPayout(); ok {
		_spec.SetField(commissionplan.FieldMinimumPayout, field.TypeFloat64, value)
	}
	if value, ok := cpuo.mutation.AddedMinimumPayout(); ok {
		_spec.AddField(commissionplan.FieldMinimumPayout, field.TypeFloat64, value)
	}
	if value, ok := cpuo.mutation.IsActive(); ok {
		_spec.SetField(commissionplan.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := cpuo.mutation.IsDefault(); ok {
		_spec.SetField(commissionplan.FieldIsDefault, field.TypeBool, value)
	}
	if cpuo.mutation.PublishersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   commissionplan.PublishersTable,
			Columns: []string{commissionplan.PublishersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpuo.mutation.RemovedPublishersIDs(); len(nodes) > 0 && !cpuo.mutation.PublishersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   commissionplan.PublishersTable,
			Columns: []string{commissionplan.PublishersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpuo.mutation.PublishersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   commissionplan.PublishersTable,
			Columns: []string{commissionplan.PublishersColumn},
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
	_node = &CommissionPlan{config: cpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{commissionplan.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cpuo.mutation.done = true
	return _node, nil
}
