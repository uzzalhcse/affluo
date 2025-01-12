// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/affiliate"
	"affluo/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AffiliateCreate is the builder for creating a Affiliate entity.
type AffiliateCreate struct {
	config
	mutation *AffiliateMutation
	hooks    []Hook
}

// SetTrackingCode sets the "tracking_code" field.
func (ac *AffiliateCreate) SetTrackingCode(s string) *AffiliateCreate {
	ac.mutation.SetTrackingCode(s)
	return ac
}

// SetNillableTrackingCode sets the "tracking_code" field if the given value is not nil.
func (ac *AffiliateCreate) SetNillableTrackingCode(s *string) *AffiliateCreate {
	if s != nil {
		ac.SetTrackingCode(*s)
	}
	return ac
}

// SetAffiliateUserID sets the "affiliate_user_id" field.
func (ac *AffiliateCreate) SetAffiliateUserID(s string) *AffiliateCreate {
	ac.mutation.SetAffiliateUserID(s)
	return ac
}

// SetSource sets the "source" field.
func (ac *AffiliateCreate) SetSource(a affiliate.Source) *AffiliateCreate {
	ac.mutation.SetSource(a)
	return ac
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (ac *AffiliateCreate) SetNillableSource(a *affiliate.Source) *AffiliateCreate {
	if a != nil {
		ac.SetSource(*a)
	}
	return ac
}

// SetRegistrationDate sets the "registration_date" field.
func (ac *AffiliateCreate) SetRegistrationDate(t time.Time) *AffiliateCreate {
	ac.mutation.SetRegistrationDate(t)
	return ac
}

// SetNillableRegistrationDate sets the "registration_date" field if the given value is not nil.
func (ac *AffiliateCreate) SetNillableRegistrationDate(t *time.Time) *AffiliateCreate {
	if t != nil {
		ac.SetRegistrationDate(*t)
	}
	return ac
}

// SetFirstTransactionDate sets the "first_transaction_date" field.
func (ac *AffiliateCreate) SetFirstTransactionDate(t time.Time) *AffiliateCreate {
	ac.mutation.SetFirstTransactionDate(t)
	return ac
}

// SetNillableFirstTransactionDate sets the "first_transaction_date" field if the given value is not nil.
func (ac *AffiliateCreate) SetNillableFirstTransactionDate(t *time.Time) *AffiliateCreate {
	if t != nil {
		ac.SetFirstTransactionDate(*t)
	}
	return ac
}

// SetCommission sets the "commission" field.
func (ac *AffiliateCreate) SetCommission(f float64) *AffiliateCreate {
	ac.mutation.SetCommission(f)
	return ac
}

// SetDate sets the "date" field.
func (ac *AffiliateCreate) SetDate(t time.Time) *AffiliateCreate {
	ac.mutation.SetDate(t)
	return ac
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (ac *AffiliateCreate) SetNillableDate(t *time.Time) *AffiliateCreate {
	if t != nil {
		ac.SetDate(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AffiliateCreate) SetID(i int64) *AffiliateCreate {
	ac.mutation.SetID(i)
	return ac
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ac *AffiliateCreate) SetUserID(id int64) *AffiliateCreate {
	ac.mutation.SetUserID(id)
	return ac
}

// SetUser sets the "user" edge to the User entity.
func (ac *AffiliateCreate) SetUser(u *User) *AffiliateCreate {
	return ac.SetUserID(u.ID)
}

// Mutation returns the AffiliateMutation object of the builder.
func (ac *AffiliateCreate) Mutation() *AffiliateMutation {
	return ac.mutation
}

// Save creates the Affiliate in the database.
func (ac *AffiliateCreate) Save(ctx context.Context) (*Affiliate, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AffiliateCreate) SaveX(ctx context.Context) *Affiliate {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AffiliateCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AffiliateCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AffiliateCreate) defaults() {
	if _, ok := ac.mutation.Source(); !ok {
		v := affiliate.DefaultSource
		ac.mutation.SetSource(v)
	}
	if _, ok := ac.mutation.RegistrationDate(); !ok {
		v := affiliate.DefaultRegistrationDate
		ac.mutation.SetRegistrationDate(v)
	}
	if _, ok := ac.mutation.Date(); !ok {
		v := affiliate.DefaultDate()
		ac.mutation.SetDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AffiliateCreate) check() error {
	if v, ok := ac.mutation.TrackingCode(); ok {
		if err := affiliate.TrackingCodeValidator(v); err != nil {
			return &ValidationError{Name: "tracking_code", err: fmt.Errorf(`ent: validator failed for field "Affiliate.tracking_code": %w`, err)}
		}
	}
	if _, ok := ac.mutation.AffiliateUserID(); !ok {
		return &ValidationError{Name: "affiliate_user_id", err: errors.New(`ent: missing required field "Affiliate.affiliate_user_id"`)}
	}
	if v, ok := ac.mutation.AffiliateUserID(); ok {
		if err := affiliate.AffiliateUserIDValidator(v); err != nil {
			return &ValidationError{Name: "affiliate_user_id", err: fmt.Errorf(`ent: validator failed for field "Affiliate.affiliate_user_id": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Source(); !ok {
		return &ValidationError{Name: "source", err: errors.New(`ent: missing required field "Affiliate.source"`)}
	}
	if v, ok := ac.mutation.Source(); ok {
		if err := affiliate.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`ent: validator failed for field "Affiliate.source": %w`, err)}
		}
	}
	if _, ok := ac.mutation.RegistrationDate(); !ok {
		return &ValidationError{Name: "registration_date", err: errors.New(`ent: missing required field "Affiliate.registration_date"`)}
	}
	if _, ok := ac.mutation.Commission(); !ok {
		return &ValidationError{Name: "commission", err: errors.New(`ent: missing required field "Affiliate.commission"`)}
	}
	if _, ok := ac.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "Affiliate.date"`)}
	}
	if v, ok := ac.mutation.ID(); ok {
		if err := affiliate.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Affiliate.id": %w`, err)}
		}
	}
	if len(ac.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Affiliate.user"`)}
	}
	return nil
}

func (ac *AffiliateCreate) sqlSave(ctx context.Context) (*Affiliate, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AffiliateCreate) createSpec() (*Affiliate, *sqlgraph.CreateSpec) {
	var (
		_node = &Affiliate{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(affiliate.Table, sqlgraph.NewFieldSpec(affiliate.FieldID, field.TypeInt64))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.TrackingCode(); ok {
		_spec.SetField(affiliate.FieldTrackingCode, field.TypeString, value)
		_node.TrackingCode = value
	}
	if value, ok := ac.mutation.AffiliateUserID(); ok {
		_spec.SetField(affiliate.FieldAffiliateUserID, field.TypeString, value)
		_node.AffiliateUserID = value
	}
	if value, ok := ac.mutation.Source(); ok {
		_spec.SetField(affiliate.FieldSource, field.TypeEnum, value)
		_node.Source = value
	}
	if value, ok := ac.mutation.RegistrationDate(); ok {
		_spec.SetField(affiliate.FieldRegistrationDate, field.TypeTime, value)
		_node.RegistrationDate = value
	}
	if value, ok := ac.mutation.FirstTransactionDate(); ok {
		_spec.SetField(affiliate.FieldFirstTransactionDate, field.TypeTime, value)
		_node.FirstTransactionDate = value
	}
	if value, ok := ac.mutation.Commission(); ok {
		_spec.SetField(affiliate.FieldCommission, field.TypeFloat64, value)
		_node.Commission = value
	}
	if value, ok := ac.mutation.Date(); ok {
		_spec.SetField(affiliate.FieldDate, field.TypeTime, value)
		_node.Date = value
	}
	if nodes := ac.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   affiliate.UserTable,
			Columns: []string{affiliate.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_affiliates = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AffiliateCreateBulk is the builder for creating many Affiliate entities in bulk.
type AffiliateCreateBulk struct {
	config
	err      error
	builders []*AffiliateCreate
}

// Save creates the Affiliate entities in the database.
func (acb *AffiliateCreateBulk) Save(ctx context.Context) ([]*Affiliate, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Affiliate, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AffiliateMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AffiliateCreateBulk) SaveX(ctx context.Context) []*Affiliate {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AffiliateCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AffiliateCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
