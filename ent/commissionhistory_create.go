// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/commissionhistory"
	"affluo/ent/user"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommissionHistoryCreate is the builder for creating a CommissionHistory entity.
type CommissionHistoryCreate struct {
	config
	mutation *CommissionHistoryMutation
	hooks    []Hook
}

// SetAmount sets the "amount" field.
func (chc *CommissionHistoryCreate) SetAmount(f float64) *CommissionHistoryCreate {
	chc.mutation.SetAmount(f)
	return chc
}

// SetAffiliateUserID sets the "affiliate_user_id" field.
func (chc *CommissionHistoryCreate) SetAffiliateUserID(s string) *CommissionHistoryCreate {
	chc.mutation.SetAffiliateUserID(s)
	return chc
}

// SetTrxID sets the "trx_id" field.
func (chc *CommissionHistoryCreate) SetTrxID(s string) *CommissionHistoryCreate {
	chc.mutation.SetTrxID(s)
	return chc
}

// SetNillableTrxID sets the "trx_id" field if the given value is not nil.
func (chc *CommissionHistoryCreate) SetNillableTrxID(s *string) *CommissionHistoryCreate {
	if s != nil {
		chc.SetTrxID(*s)
	}
	return chc
}

// SetTrackID sets the "track_id" field.
func (chc *CommissionHistoryCreate) SetTrackID(s string) *CommissionHistoryCreate {
	chc.mutation.SetTrackID(s)
	return chc
}

// SetNillableTrackID sets the "track_id" field if the given value is not nil.
func (chc *CommissionHistoryCreate) SetNillableTrackID(s *string) *CommissionHistoryCreate {
	if s != nil {
		chc.SetTrackID(*s)
	}
	return chc
}

// SetCommissionRate sets the "commission_rate" field.
func (chc *CommissionHistoryCreate) SetCommissionRate(f float64) *CommissionHistoryCreate {
	chc.mutation.SetCommissionRate(f)
	return chc
}

// SetIsFirstOrder sets the "is_first_order" field.
func (chc *CommissionHistoryCreate) SetIsFirstOrder(b bool) *CommissionHistoryCreate {
	chc.mutation.SetIsFirstOrder(b)
	return chc
}

// SetDate sets the "date" field.
func (chc *CommissionHistoryCreate) SetDate(s string) *CommissionHistoryCreate {
	chc.mutation.SetDate(s)
	return chc
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (chc *CommissionHistoryCreate) SetNillableDate(s *string) *CommissionHistoryCreate {
	if s != nil {
		chc.SetDate(*s)
	}
	return chc
}

// SetID sets the "id" field.
func (chc *CommissionHistoryCreate) SetID(i int64) *CommissionHistoryCreate {
	chc.mutation.SetID(i)
	return chc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (chc *CommissionHistoryCreate) SetUserID(id int64) *CommissionHistoryCreate {
	chc.mutation.SetUserID(id)
	return chc
}

// SetUser sets the "user" edge to the User entity.
func (chc *CommissionHistoryCreate) SetUser(u *User) *CommissionHistoryCreate {
	return chc.SetUserID(u.ID)
}

// Mutation returns the CommissionHistoryMutation object of the builder.
func (chc *CommissionHistoryCreate) Mutation() *CommissionHistoryMutation {
	return chc.mutation
}

// Save creates the CommissionHistory in the database.
func (chc *CommissionHistoryCreate) Save(ctx context.Context) (*CommissionHistory, error) {
	chc.defaults()
	return withHooks(ctx, chc.sqlSave, chc.mutation, chc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (chc *CommissionHistoryCreate) SaveX(ctx context.Context) *CommissionHistory {
	v, err := chc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (chc *CommissionHistoryCreate) Exec(ctx context.Context) error {
	_, err := chc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (chc *CommissionHistoryCreate) ExecX(ctx context.Context) {
	if err := chc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (chc *CommissionHistoryCreate) defaults() {
	if _, ok := chc.mutation.Date(); !ok {
		v := commissionhistory.DefaultDate
		chc.mutation.SetDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (chc *CommissionHistoryCreate) check() error {
	if _, ok := chc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "CommissionHistory.amount"`)}
	}
	if _, ok := chc.mutation.AffiliateUserID(); !ok {
		return &ValidationError{Name: "affiliate_user_id", err: errors.New(`ent: missing required field "CommissionHistory.affiliate_user_id"`)}
	}
	if _, ok := chc.mutation.CommissionRate(); !ok {
		return &ValidationError{Name: "commission_rate", err: errors.New(`ent: missing required field "CommissionHistory.commission_rate"`)}
	}
	if _, ok := chc.mutation.IsFirstOrder(); !ok {
		return &ValidationError{Name: "is_first_order", err: errors.New(`ent: missing required field "CommissionHistory.is_first_order"`)}
	}
	if _, ok := chc.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "CommissionHistory.date"`)}
	}
	if v, ok := chc.mutation.ID(); ok {
		if err := commissionhistory.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "CommissionHistory.id": %w`, err)}
		}
	}
	if len(chc.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "CommissionHistory.user"`)}
	}
	return nil
}

func (chc *CommissionHistoryCreate) sqlSave(ctx context.Context) (*CommissionHistory, error) {
	if err := chc.check(); err != nil {
		return nil, err
	}
	_node, _spec := chc.createSpec()
	if err := sqlgraph.CreateNode(ctx, chc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	chc.mutation.id = &_node.ID
	chc.mutation.done = true
	return _node, nil
}

func (chc *CommissionHistoryCreate) createSpec() (*CommissionHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &CommissionHistory{config: chc.config}
		_spec = sqlgraph.NewCreateSpec(commissionhistory.Table, sqlgraph.NewFieldSpec(commissionhistory.FieldID, field.TypeInt64))
	)
	if id, ok := chc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := chc.mutation.Amount(); ok {
		_spec.SetField(commissionhistory.FieldAmount, field.TypeFloat64, value)
		_node.Amount = value
	}
	if value, ok := chc.mutation.AffiliateUserID(); ok {
		_spec.SetField(commissionhistory.FieldAffiliateUserID, field.TypeString, value)
		_node.AffiliateUserID = value
	}
	if value, ok := chc.mutation.TrxID(); ok {
		_spec.SetField(commissionhistory.FieldTrxID, field.TypeString, value)
		_node.TrxID = value
	}
	if value, ok := chc.mutation.TrackID(); ok {
		_spec.SetField(commissionhistory.FieldTrackID, field.TypeString, value)
		_node.TrackID = value
	}
	if value, ok := chc.mutation.CommissionRate(); ok {
		_spec.SetField(commissionhistory.FieldCommissionRate, field.TypeFloat64, value)
		_node.CommissionRate = value
	}
	if value, ok := chc.mutation.IsFirstOrder(); ok {
		_spec.SetField(commissionhistory.FieldIsFirstOrder, field.TypeBool, value)
		_node.IsFirstOrder = value
	}
	if value, ok := chc.mutation.Date(); ok {
		_spec.SetField(commissionhistory.FieldDate, field.TypeString, value)
		_node.Date = value
	}
	if nodes := chc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   commissionhistory.UserTable,
			Columns: []string{commissionhistory.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_commission_histories = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CommissionHistoryCreateBulk is the builder for creating many CommissionHistory entities in bulk.
type CommissionHistoryCreateBulk struct {
	config
	err      error
	builders []*CommissionHistoryCreate
}

// Save creates the CommissionHistory entities in the database.
func (chcb *CommissionHistoryCreateBulk) Save(ctx context.Context) ([]*CommissionHistory, error) {
	if chcb.err != nil {
		return nil, chcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(chcb.builders))
	nodes := make([]*CommissionHistory, len(chcb.builders))
	mutators := make([]Mutator, len(chcb.builders))
	for i := range chcb.builders {
		func(i int, root context.Context) {
			builder := chcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommissionHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, chcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, chcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, chcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (chcb *CommissionHistoryCreateBulk) SaveX(ctx context.Context) []*CommissionHistory {
	v, err := chcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (chcb *CommissionHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := chcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (chcb *CommissionHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := chcb.Exec(ctx); err != nil {
		panic(err)
	}
}
