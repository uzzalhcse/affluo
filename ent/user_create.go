// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/bannerstats"
	"affluo/ent/campaign"
	"affluo/ent/commissionplan"
	"affluo/ent/gigtracking"
	"affluo/ent/payout"
	"affluo/ent/post"
	"affluo/ent/referral"
	"affluo/ent/track"
	"affluo/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetPasswordHash sets the "password_hash" field.
func (uc *UserCreate) SetPasswordHash(s string) *UserCreate {
	uc.mutation.SetPasswordHash(s)
	return uc
}

// SetFirstName sets the "first_name" field.
func (uc *UserCreate) SetFirstName(s string) *UserCreate {
	uc.mutation.SetFirstName(s)
	return uc
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (uc *UserCreate) SetNillableFirstName(s *string) *UserCreate {
	if s != nil {
		uc.SetFirstName(*s)
	}
	return uc
}

// SetLastName sets the "last_name" field.
func (uc *UserCreate) SetLastName(s string) *UserCreate {
	uc.mutation.SetLastName(s)
	return uc
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (uc *UserCreate) SetNillableLastName(s *string) *UserCreate {
	if s != nil {
		uc.SetLastName(*s)
	}
	return uc
}

// SetRole sets the "role" field.
func (uc *UserCreate) SetRole(u user.Role) *UserCreate {
	uc.mutation.SetRole(u)
	return uc
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uc *UserCreate) SetNillableRole(u *user.Role) *UserCreate {
	if u != nil {
		uc.SetRole(*u)
	}
	return uc
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetIsActive sets the "is_active" field.
func (uc *UserCreate) SetIsActive(b bool) *UserCreate {
	uc.mutation.SetIsActive(b)
	return uc
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (uc *UserCreate) SetNillableIsActive(b *bool) *UserCreate {
	if b != nil {
		uc.SetIsActive(*b)
	}
	return uc
}

// SetLastLogin sets the "last_login" field.
func (uc *UserCreate) SetLastLogin(t time.Time) *UserCreate {
	uc.mutation.SetLastLogin(t)
	return uc
}

// SetNillableLastLogin sets the "last_login" field if the given value is not nil.
func (uc *UserCreate) SetNillableLastLogin(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetLastLogin(*t)
	}
	return uc
}

// SetResetToken sets the "reset_token" field.
func (uc *UserCreate) SetResetToken(s string) *UserCreate {
	uc.mutation.SetResetToken(s)
	return uc
}

// SetNillableResetToken sets the "reset_token" field if the given value is not nil.
func (uc *UserCreate) SetNillableResetToken(s *string) *UserCreate {
	if s != nil {
		uc.SetResetToken(*s)
	}
	return uc
}

// SetResetTokenExpiresAt sets the "reset_token_expires_at" field.
func (uc *UserCreate) SetResetTokenExpiresAt(t time.Time) *UserCreate {
	uc.mutation.SetResetTokenExpiresAt(t)
	return uc
}

// SetNillableResetTokenExpiresAt sets the "reset_token_expires_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableResetTokenExpiresAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetResetTokenExpiresAt(*t)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(i int64) *UserCreate {
	uc.mutation.SetID(i)
	return uc
}

// AddCampaignIDs adds the "campaigns" edge to the Campaign entity by IDs.
func (uc *UserCreate) AddCampaignIDs(ids ...int64) *UserCreate {
	uc.mutation.AddCampaignIDs(ids...)
	return uc
}

// AddCampaigns adds the "campaigns" edges to the Campaign entity.
func (uc *UserCreate) AddCampaigns(c ...*Campaign) *UserCreate {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddCampaignIDs(ids...)
}

// AddReferralIDs adds the "referrals" edge to the Referral entity by IDs.
func (uc *UserCreate) AddReferralIDs(ids ...int64) *UserCreate {
	uc.mutation.AddReferralIDs(ids...)
	return uc
}

// AddReferrals adds the "referrals" edges to the Referral entity.
func (uc *UserCreate) AddReferrals(r ...*Referral) *UserCreate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uc.AddReferralIDs(ids...)
}

// AddTrackIDs adds the "tracks" edge to the Track entity by IDs.
func (uc *UserCreate) AddTrackIDs(ids ...int64) *UserCreate {
	uc.mutation.AddTrackIDs(ids...)
	return uc
}

// AddTracks adds the "tracks" edges to the Track entity.
func (uc *UserCreate) AddTracks(t ...*Track) *UserCreate {
	ids := make([]int64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uc.AddTrackIDs(ids...)
}

// AddPayoutIDs adds the "payouts" edge to the Payout entity by IDs.
func (uc *UserCreate) AddPayoutIDs(ids ...string) *UserCreate {
	uc.mutation.AddPayoutIDs(ids...)
	return uc
}

// AddPayouts adds the "payouts" edges to the Payout entity.
func (uc *UserCreate) AddPayouts(p ...*Payout) *UserCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uc.AddPayoutIDs(ids...)
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (uc *UserCreate) AddPostIDs(ids ...string) *UserCreate {
	uc.mutation.AddPostIDs(ids...)
	return uc
}

// AddPosts adds the "posts" edges to the Post entity.
func (uc *UserCreate) AddPosts(p ...*Post) *UserCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uc.AddPostIDs(ids...)
}

// AddStatIDs adds the "stats" edge to the BannerStats entity by IDs.
func (uc *UserCreate) AddStatIDs(ids ...int64) *UserCreate {
	uc.mutation.AddStatIDs(ids...)
	return uc
}

// AddStats adds the "stats" edges to the BannerStats entity.
func (uc *UserCreate) AddStats(b ...*BannerStats) *UserCreate {
	ids := make([]int64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uc.AddStatIDs(ids...)
}

// AddGigTrackingIDs adds the "gig_trackings" edge to the GigTracking entity by IDs.
func (uc *UserCreate) AddGigTrackingIDs(ids ...int64) *UserCreate {
	uc.mutation.AddGigTrackingIDs(ids...)
	return uc
}

// AddGigTrackings adds the "gig_trackings" edges to the GigTracking entity.
func (uc *UserCreate) AddGigTrackings(g ...*GigTracking) *UserCreate {
	ids := make([]int64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uc.AddGigTrackingIDs(ids...)
}

// SetCommissionPlanID sets the "commission_plan" edge to the CommissionPlan entity by ID.
func (uc *UserCreate) SetCommissionPlanID(id int) *UserCreate {
	uc.mutation.SetCommissionPlanID(id)
	return uc
}

// SetNillableCommissionPlanID sets the "commission_plan" edge to the CommissionPlan entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableCommissionPlanID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetCommissionPlanID(*id)
	}
	return uc
}

// SetCommissionPlan sets the "commission_plan" edge to the CommissionPlan entity.
func (uc *UserCreate) SetCommissionPlan(c *CommissionPlan) *UserCreate {
	return uc.SetCommissionPlanID(c.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.Role(); !ok {
		v := user.DefaultRole
		uc.mutation.SetRole(v)
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := user.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
	if _, ok := uc.mutation.IsActive(); !ok {
		v := user.DefaultIsActive
		uc.mutation.SetIsActive(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "User.username"`)}
	}
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "User.email"`)}
	}
	if v, ok := uc.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if _, ok := uc.mutation.PasswordHash(); !ok {
		return &ValidationError{Name: "password_hash", err: errors.New(`ent: missing required field "User.password_hash"`)}
	}
	if _, ok := uc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required field "User.role"`)}
	}
	if v, ok := uc.mutation.Role(); ok {
		if err := user.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "User.role": %w`, err)}
		}
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "User.created_at"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "User.updated_at"`)}
	}
	if _, ok := uc.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "User.is_active"`)}
	}
	if v, ok := uc.mutation.ID(); ok {
		if err := user.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "User.id": %w`, err)}
		}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := uc.mutation.PasswordHash(); ok {
		_spec.SetField(user.FieldPasswordHash, field.TypeString, value)
		_node.PasswordHash = value
	}
	if value, ok := uc.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
		_node.FirstName = value
	}
	if value, ok := uc.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}
	if value, ok := uc.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeEnum, value)
		_node.Role = value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := uc.mutation.IsActive(); ok {
		_spec.SetField(user.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if value, ok := uc.mutation.LastLogin(); ok {
		_spec.SetField(user.FieldLastLogin, field.TypeTime, value)
		_node.LastLogin = value
	}
	if value, ok := uc.mutation.ResetToken(); ok {
		_spec.SetField(user.FieldResetToken, field.TypeString, value)
		_node.ResetToken = value
	}
	if value, ok := uc.mutation.ResetTokenExpiresAt(); ok {
		_spec.SetField(user.FieldResetTokenExpiresAt, field.TypeTime, value)
		_node.ResetTokenExpiresAt = value
	}
	if nodes := uc.mutation.CampaignsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CampaignsTable,
			Columns: []string{user.CampaignsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(campaign.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ReferralsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReferralsTable,
			Columns: []string{user.ReferralsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(referral.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.TracksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TracksTable,
			Columns: []string{user.TracksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.PayoutsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PayoutsTable,
			Columns: []string{user.PayoutsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(payout.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.StatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.StatsTable,
			Columns: []string{user.StatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bannerstats.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.GigTrackingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.GigTrackingsTable,
			Columns: []string{user.GigTrackingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gigtracking.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.CommissionPlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.CommissionPlanTable,
			Columns: []string{user.CommissionPlanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(commissionplan.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.commission_plan_publishers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
