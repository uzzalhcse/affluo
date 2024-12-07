// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/campaign"
	"affluo/ent/campaignlink"
	"affluo/ent/track"
	"affluo/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TrackCreate is the builder for creating a Track entity.
type TrackCreate struct {
	config
	mutation *TrackMutation
	hooks    []Hook
}

// SetIPAddress sets the "ip_address" field.
func (tc *TrackCreate) SetIPAddress(s string) *TrackCreate {
	tc.mutation.SetIPAddress(s)
	return tc
}

// SetUserAgent sets the "user_agent" field.
func (tc *TrackCreate) SetUserAgent(s string) *TrackCreate {
	tc.mutation.SetUserAgent(s)
	return tc
}

// SetDeviceFingerprint sets the "device_fingerprint" field.
func (tc *TrackCreate) SetDeviceFingerprint(s string) *TrackCreate {
	tc.mutation.SetDeviceFingerprint(s)
	return tc
}

// SetReferrer sets the "referrer" field.
func (tc *TrackCreate) SetReferrer(s string) *TrackCreate {
	tc.mutation.SetReferrer(s)
	return tc
}

// SetNillableReferrer sets the "referrer" field if the given value is not nil.
func (tc *TrackCreate) SetNillableReferrer(s *string) *TrackCreate {
	if s != nil {
		tc.SetReferrer(*s)
	}
	return tc
}

// SetType sets the "type" field.
func (tc *TrackCreate) SetType(t track.Type) *TrackCreate {
	tc.mutation.SetType(t)
	return tc
}

// SetStatus sets the "status" field.
func (tc *TrackCreate) SetStatus(t track.Status) *TrackCreate {
	tc.mutation.SetStatus(t)
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TrackCreate) SetCreatedAt(t time.Time) *TrackCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TrackCreate) SetNillableCreatedAt(t *time.Time) *TrackCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetIsUniqueClick sets the "is_unique_click" field.
func (tc *TrackCreate) SetIsUniqueClick(b bool) *TrackCreate {
	tc.mutation.SetIsUniqueClick(b)
	return tc
}

// SetNillableIsUniqueClick sets the "is_unique_click" field if the given value is not nil.
func (tc *TrackCreate) SetNillableIsUniqueClick(b *bool) *TrackCreate {
	if b != nil {
		tc.SetIsUniqueClick(*b)
	}
	return tc
}

// SetAdditionalMetadata sets the "additional_metadata" field.
func (tc *TrackCreate) SetAdditionalMetadata(m map[string]interface{}) *TrackCreate {
	tc.mutation.SetAdditionalMetadata(m)
	return tc
}

// SetID sets the "id" field.
func (tc *TrackCreate) SetID(i int64) *TrackCreate {
	tc.mutation.SetID(i)
	return tc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (tc *TrackCreate) SetUserID(id int64) *TrackCreate {
	tc.mutation.SetUserID(id)
	return tc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (tc *TrackCreate) SetNillableUserID(id *int64) *TrackCreate {
	if id != nil {
		tc = tc.SetUserID(*id)
	}
	return tc
}

// SetUser sets the "user" edge to the User entity.
func (tc *TrackCreate) SetUser(u *User) *TrackCreate {
	return tc.SetUserID(u.ID)
}

// SetCampaignID sets the "campaign" edge to the Campaign entity by ID.
func (tc *TrackCreate) SetCampaignID(id int64) *TrackCreate {
	tc.mutation.SetCampaignID(id)
	return tc
}

// SetNillableCampaignID sets the "campaign" edge to the Campaign entity by ID if the given value is not nil.
func (tc *TrackCreate) SetNillableCampaignID(id *int64) *TrackCreate {
	if id != nil {
		tc = tc.SetCampaignID(*id)
	}
	return tc
}

// SetCampaign sets the "campaign" edge to the Campaign entity.
func (tc *TrackCreate) SetCampaign(c *Campaign) *TrackCreate {
	return tc.SetCampaignID(c.ID)
}

// SetLinkID sets the "link" edge to the CampaignLink entity by ID.
func (tc *TrackCreate) SetLinkID(id int64) *TrackCreate {
	tc.mutation.SetLinkID(id)
	return tc
}

// SetNillableLinkID sets the "link" edge to the CampaignLink entity by ID if the given value is not nil.
func (tc *TrackCreate) SetNillableLinkID(id *int64) *TrackCreate {
	if id != nil {
		tc = tc.SetLinkID(*id)
	}
	return tc
}

// SetLink sets the "link" edge to the CampaignLink entity.
func (tc *TrackCreate) SetLink(c *CampaignLink) *TrackCreate {
	return tc.SetLinkID(c.ID)
}

// Mutation returns the TrackMutation object of the builder.
func (tc *TrackCreate) Mutation() *TrackMutation {
	return tc.mutation
}

// Save creates the Track in the database.
func (tc *TrackCreate) Save(ctx context.Context) (*Track, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TrackCreate) SaveX(ctx context.Context) *Track {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TrackCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TrackCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TrackCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := track.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.IsUniqueClick(); !ok {
		v := track.DefaultIsUniqueClick
		tc.mutation.SetIsUniqueClick(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TrackCreate) check() error {
	if _, ok := tc.mutation.IPAddress(); !ok {
		return &ValidationError{Name: "ip_address", err: errors.New(`ent: missing required field "Track.ip_address"`)}
	}
	if _, ok := tc.mutation.UserAgent(); !ok {
		return &ValidationError{Name: "user_agent", err: errors.New(`ent: missing required field "Track.user_agent"`)}
	}
	if _, ok := tc.mutation.DeviceFingerprint(); !ok {
		return &ValidationError{Name: "device_fingerprint", err: errors.New(`ent: missing required field "Track.device_fingerprint"`)}
	}
	if _, ok := tc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Track.type"`)}
	}
	if v, ok := tc.mutation.GetType(); ok {
		if err := track.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Track.type": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Track.status"`)}
	}
	if v, ok := tc.mutation.Status(); ok {
		if err := track.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Track.status": %w`, err)}
		}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Track.created_at"`)}
	}
	if _, ok := tc.mutation.IsUniqueClick(); !ok {
		return &ValidationError{Name: "is_unique_click", err: errors.New(`ent: missing required field "Track.is_unique_click"`)}
	}
	if v, ok := tc.mutation.ID(); ok {
		if err := track.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Track.id": %w`, err)}
		}
	}
	return nil
}

func (tc *TrackCreate) sqlSave(ctx context.Context) (*Track, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TrackCreate) createSpec() (*Track, *sqlgraph.CreateSpec) {
	var (
		_node = &Track{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(track.Table, sqlgraph.NewFieldSpec(track.FieldID, field.TypeInt64))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.IPAddress(); ok {
		_spec.SetField(track.FieldIPAddress, field.TypeString, value)
		_node.IPAddress = value
	}
	if value, ok := tc.mutation.UserAgent(); ok {
		_spec.SetField(track.FieldUserAgent, field.TypeString, value)
		_node.UserAgent = value
	}
	if value, ok := tc.mutation.DeviceFingerprint(); ok {
		_spec.SetField(track.FieldDeviceFingerprint, field.TypeString, value)
		_node.DeviceFingerprint = value
	}
	if value, ok := tc.mutation.Referrer(); ok {
		_spec.SetField(track.FieldReferrer, field.TypeString, value)
		_node.Referrer = value
	}
	if value, ok := tc.mutation.GetType(); ok {
		_spec.SetField(track.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := tc.mutation.Status(); ok {
		_spec.SetField(track.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(track.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.IsUniqueClick(); ok {
		_spec.SetField(track.FieldIsUniqueClick, field.TypeBool, value)
		_node.IsUniqueClick = value
	}
	if value, ok := tc.mutation.AdditionalMetadata(); ok {
		_spec.SetField(track.FieldAdditionalMetadata, field.TypeJSON, value)
		_node.AdditionalMetadata = value
	}
	if nodes := tc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   track.UserTable,
			Columns: []string{track.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_tracks = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.CampaignIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   track.CampaignTable,
			Columns: []string{track.CampaignColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(campaign.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.campaign_tracks = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.LinkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   track.LinkTable,
			Columns: []string{track.LinkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(campaignlink.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.campaign_link_tracks = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TrackCreateBulk is the builder for creating many Track entities in bulk.
type TrackCreateBulk struct {
	config
	err      error
	builders []*TrackCreate
}

// Save creates the Track entities in the database.
func (tcb *TrackCreateBulk) Save(ctx context.Context) ([]*Track, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Track, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TrackMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TrackCreateBulk) SaveX(ctx context.Context) []*Track {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TrackCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TrackCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
