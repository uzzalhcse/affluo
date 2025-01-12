// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/commissionplan"
	"affluo/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommissionPlanCreate is the builder for creating a CommissionPlan entity.
type CommissionPlanCreate struct {
	config
	mutation *CommissionPlanMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cpc *CommissionPlanCreate) SetName(s string) *CommissionPlanCreate {
	cpc.mutation.SetName(s)
	return cpc
}

// SetDescription sets the "description" field.
func (cpc *CommissionPlanCreate) SetDescription(s string) *CommissionPlanCreate {
	cpc.mutation.SetDescription(s)
	return cpc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cpc *CommissionPlanCreate) SetNillableDescription(s *string) *CommissionPlanCreate {
	if s != nil {
		cpc.SetDescription(*s)
	}
	return cpc
}

// SetType sets the "type" field.
func (cpc *CommissionPlanCreate) SetType(c commissionplan.Type) *CommissionPlanCreate {
	cpc.mutation.SetType(c)
	return cpc
}

// SetClickCommission sets the "click_commission" field.
func (cpc *CommissionPlanCreate) SetClickCommission(f float64) *CommissionPlanCreate {
	cpc.mutation.SetClickCommission(f)
	return cpc
}

// SetNillableClickCommission sets the "click_commission" field if the given value is not nil.
func (cpc *CommissionPlanCreate) SetNillableClickCommission(f *float64) *CommissionPlanCreate {
	if f != nil {
		cpc.SetClickCommission(*f)
	}
	return cpc
}

// SetImpressionCommission sets the "impression_commission" field.
func (cpc *CommissionPlanCreate) SetImpressionCommission(f float64) *CommissionPlanCreate {
	cpc.mutation.SetImpressionCommission(f)
	return cpc
}

// SetNillableImpressionCommission sets the "impression_commission" field if the given value is not nil.
func (cpc *CommissionPlanCreate) SetNillableImpressionCommission(f *float64) *CommissionPlanCreate {
	if f != nil {
		cpc.SetImpressionCommission(*f)
	}
	return cpc
}

// SetLeadCommission sets the "lead_commission" field.
func (cpc *CommissionPlanCreate) SetLeadCommission(f float64) *CommissionPlanCreate {
	cpc.mutation.SetLeadCommission(f)
	return cpc
}

// SetNillableLeadCommission sets the "lead_commission" field if the given value is not nil.
func (cpc *CommissionPlanCreate) SetNillableLeadCommission(f *float64) *CommissionPlanCreate {
	if f != nil {
		cpc.SetLeadCommission(*f)
	}
	return cpc
}

// SetMinimumPayout sets the "minimum_payout" field.
func (cpc *CommissionPlanCreate) SetMinimumPayout(f float64) *CommissionPlanCreate {
	cpc.mutation.SetMinimumPayout(f)
	return cpc
}

// SetNillableMinimumPayout sets the "minimum_payout" field if the given value is not nil.
func (cpc *CommissionPlanCreate) SetNillableMinimumPayout(f *float64) *CommissionPlanCreate {
	if f != nil {
		cpc.SetMinimumPayout(*f)
	}
	return cpc
}

// SetValidFrom sets the "valid_from" field.
func (cpc *CommissionPlanCreate) SetValidFrom(t time.Time) *CommissionPlanCreate {
	cpc.mutation.SetValidFrom(t)
	return cpc
}

// SetNillableValidFrom sets the "valid_from" field if the given value is not nil.
func (cpc *CommissionPlanCreate) SetNillableValidFrom(t *time.Time) *CommissionPlanCreate {
	if t != nil {
		cpc.SetValidFrom(*t)
	}
	return cpc
}

// SetValidUntil sets the "valid_until" field.
func (cpc *CommissionPlanCreate) SetValidUntil(t time.Time) *CommissionPlanCreate {
	cpc.mutation.SetValidUntil(t)
	return cpc
}

// SetNillableValidUntil sets the "valid_until" field if the given value is not nil.
func (cpc *CommissionPlanCreate) SetNillableValidUntil(t *time.Time) *CommissionPlanCreate {
	if t != nil {
		cpc.SetValidUntil(*t)
	}
	return cpc
}

// SetIsActive sets the "is_active" field.
func (cpc *CommissionPlanCreate) SetIsActive(b bool) *CommissionPlanCreate {
	cpc.mutation.SetIsActive(b)
	return cpc
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (cpc *CommissionPlanCreate) SetNillableIsActive(b *bool) *CommissionPlanCreate {
	if b != nil {
		cpc.SetIsActive(*b)
	}
	return cpc
}

// SetIsDefault sets the "is_default" field.
func (cpc *CommissionPlanCreate) SetIsDefault(b bool) *CommissionPlanCreate {
	cpc.mutation.SetIsDefault(b)
	return cpc
}

// SetNillableIsDefault sets the "is_default" field if the given value is not nil.
func (cpc *CommissionPlanCreate) SetNillableIsDefault(b *bool) *CommissionPlanCreate {
	if b != nil {
		cpc.SetIsDefault(*b)
	}
	return cpc
}

// AddPublisherIDs adds the "publishers" edge to the User entity by IDs.
func (cpc *CommissionPlanCreate) AddPublisherIDs(ids ...int64) *CommissionPlanCreate {
	cpc.mutation.AddPublisherIDs(ids...)
	return cpc
}

// AddPublishers adds the "publishers" edges to the User entity.
func (cpc *CommissionPlanCreate) AddPublishers(u ...*User) *CommissionPlanCreate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cpc.AddPublisherIDs(ids...)
}

// Mutation returns the CommissionPlanMutation object of the builder.
func (cpc *CommissionPlanCreate) Mutation() *CommissionPlanMutation {
	return cpc.mutation
}

// Save creates the CommissionPlan in the database.
func (cpc *CommissionPlanCreate) Save(ctx context.Context) (*CommissionPlan, error) {
	cpc.defaults()
	return withHooks(ctx, cpc.sqlSave, cpc.mutation, cpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cpc *CommissionPlanCreate) SaveX(ctx context.Context) *CommissionPlan {
	v, err := cpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cpc *CommissionPlanCreate) Exec(ctx context.Context) error {
	_, err := cpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpc *CommissionPlanCreate) ExecX(ctx context.Context) {
	if err := cpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cpc *CommissionPlanCreate) defaults() {
	if _, ok := cpc.mutation.ClickCommission(); !ok {
		v := commissionplan.DefaultClickCommission
		cpc.mutation.SetClickCommission(v)
	}
	if _, ok := cpc.mutation.ImpressionCommission(); !ok {
		v := commissionplan.DefaultImpressionCommission
		cpc.mutation.SetImpressionCommission(v)
	}
	if _, ok := cpc.mutation.LeadCommission(); !ok {
		v := commissionplan.DefaultLeadCommission
		cpc.mutation.SetLeadCommission(v)
	}
	if _, ok := cpc.mutation.MinimumPayout(); !ok {
		v := commissionplan.DefaultMinimumPayout
		cpc.mutation.SetMinimumPayout(v)
	}
	if _, ok := cpc.mutation.IsActive(); !ok {
		v := commissionplan.DefaultIsActive
		cpc.mutation.SetIsActive(v)
	}
	if _, ok := cpc.mutation.IsDefault(); !ok {
		v := commissionplan.DefaultIsDefault
		cpc.mutation.SetIsDefault(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cpc *CommissionPlanCreate) check() error {
	if _, ok := cpc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "CommissionPlan.name"`)}
	}
	if v, ok := cpc.mutation.Name(); ok {
		if err := commissionplan.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "CommissionPlan.name": %w`, err)}
		}
	}
	if _, ok := cpc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "CommissionPlan.type"`)}
	}
	if v, ok := cpc.mutation.GetType(); ok {
		if err := commissionplan.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "CommissionPlan.type": %w`, err)}
		}
	}
	if _, ok := cpc.mutation.ClickCommission(); !ok {
		return &ValidationError{Name: "click_commission", err: errors.New(`ent: missing required field "CommissionPlan.click_commission"`)}
	}
	if _, ok := cpc.mutation.ImpressionCommission(); !ok {
		return &ValidationError{Name: "impression_commission", err: errors.New(`ent: missing required field "CommissionPlan.impression_commission"`)}
	}
	if _, ok := cpc.mutation.LeadCommission(); !ok {
		return &ValidationError{Name: "lead_commission", err: errors.New(`ent: missing required field "CommissionPlan.lead_commission"`)}
	}
	if _, ok := cpc.mutation.MinimumPayout(); !ok {
		return &ValidationError{Name: "minimum_payout", err: errors.New(`ent: missing required field "CommissionPlan.minimum_payout"`)}
	}
	if _, ok := cpc.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "CommissionPlan.is_active"`)}
	}
	if _, ok := cpc.mutation.IsDefault(); !ok {
		return &ValidationError{Name: "is_default", err: errors.New(`ent: missing required field "CommissionPlan.is_default"`)}
	}
	return nil
}

func (cpc *CommissionPlanCreate) sqlSave(ctx context.Context) (*CommissionPlan, error) {
	if err := cpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cpc.mutation.id = &_node.ID
	cpc.mutation.done = true
	return _node, nil
}

func (cpc *CommissionPlanCreate) createSpec() (*CommissionPlan, *sqlgraph.CreateSpec) {
	var (
		_node = &CommissionPlan{config: cpc.config}
		_spec = sqlgraph.NewCreateSpec(commissionplan.Table, sqlgraph.NewFieldSpec(commissionplan.FieldID, field.TypeInt))
	)
	if value, ok := cpc.mutation.Name(); ok {
		_spec.SetField(commissionplan.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cpc.mutation.Description(); ok {
		_spec.SetField(commissionplan.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := cpc.mutation.GetType(); ok {
		_spec.SetField(commissionplan.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := cpc.mutation.ClickCommission(); ok {
		_spec.SetField(commissionplan.FieldClickCommission, field.TypeFloat64, value)
		_node.ClickCommission = value
	}
	if value, ok := cpc.mutation.ImpressionCommission(); ok {
		_spec.SetField(commissionplan.FieldImpressionCommission, field.TypeFloat64, value)
		_node.ImpressionCommission = value
	}
	if value, ok := cpc.mutation.LeadCommission(); ok {
		_spec.SetField(commissionplan.FieldLeadCommission, field.TypeFloat64, value)
		_node.LeadCommission = value
	}
	if value, ok := cpc.mutation.MinimumPayout(); ok {
		_spec.SetField(commissionplan.FieldMinimumPayout, field.TypeFloat64, value)
		_node.MinimumPayout = value
	}
	if value, ok := cpc.mutation.ValidFrom(); ok {
		_spec.SetField(commissionplan.FieldValidFrom, field.TypeTime, value)
		_node.ValidFrom = value
	}
	if value, ok := cpc.mutation.ValidUntil(); ok {
		_spec.SetField(commissionplan.FieldValidUntil, field.TypeTime, value)
		_node.ValidUntil = value
	}
	if value, ok := cpc.mutation.IsActive(); ok {
		_spec.SetField(commissionplan.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if value, ok := cpc.mutation.IsDefault(); ok {
		_spec.SetField(commissionplan.FieldIsDefault, field.TypeBool, value)
		_node.IsDefault = value
	}
	if nodes := cpc.mutation.PublishersIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CommissionPlanCreateBulk is the builder for creating many CommissionPlan entities in bulk.
type CommissionPlanCreateBulk struct {
	config
	err      error
	builders []*CommissionPlanCreate
}

// Save creates the CommissionPlan entities in the database.
func (cpcb *CommissionPlanCreateBulk) Save(ctx context.Context) ([]*CommissionPlan, error) {
	if cpcb.err != nil {
		return nil, cpcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cpcb.builders))
	nodes := make([]*CommissionPlan, len(cpcb.builders))
	mutators := make([]Mutator, len(cpcb.builders))
	for i := range cpcb.builders {
		func(i int, root context.Context) {
			builder := cpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommissionPlanMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, cpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cpcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, cpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cpcb *CommissionPlanCreateBulk) SaveX(ctx context.Context) []*CommissionPlan {
	v, err := cpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cpcb *CommissionPlanCreateBulk) Exec(ctx context.Context) error {
	_, err := cpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpcb *CommissionPlanCreateBulk) ExecX(ctx context.Context) {
	if err := cpcb.Exec(ctx); err != nil {
		panic(err)
	}
}