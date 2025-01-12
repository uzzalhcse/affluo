// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/banner"
	"affluo/ent/campaign"
	"affluo/ent/campaignlink"
	"affluo/ent/schema"
	"affluo/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CampaignCreate is the builder for creating a Campaign entity.
type CampaignCreate struct {
	config
	mutation *CampaignMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *CampaignCreate) SetName(s string) *CampaignCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetDescription sets the "description" field.
func (cc *CampaignCreate) SetDescription(s string) *CampaignCreate {
	cc.mutation.SetDescription(s)
	return cc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cc *CampaignCreate) SetNillableDescription(s *string) *CampaignCreate {
	if s != nil {
		cc.SetDescription(*s)
	}
	return cc
}

// SetUniqueCode sets the "unique_code" field.
func (cc *CampaignCreate) SetUniqueCode(s string) *CampaignCreate {
	cc.mutation.SetUniqueCode(s)
	return cc
}

// SetType sets the "type" field.
func (cc *CampaignCreate) SetType(c campaign.Type) *CampaignCreate {
	cc.mutation.SetType(c)
	return cc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cc *CampaignCreate) SetNillableType(c *campaign.Type) *CampaignCreate {
	if c != nil {
		cc.SetType(*c)
	}
	return cc
}

// SetCommissionType sets the "commission_type" field.
func (cc *CampaignCreate) SetCommissionType(ct campaign.CommissionType) *CampaignCreate {
	cc.mutation.SetCommissionType(ct)
	return cc
}

// SetNillableCommissionType sets the "commission_type" field if the given value is not nil.
func (cc *CampaignCreate) SetNillableCommissionType(ct *campaign.CommissionType) *CampaignCreate {
	if ct != nil {
		cc.SetCommissionType(*ct)
	}
	return cc
}

// SetBaseCommissionRate sets the "base_commission_rate" field.
func (cc *CampaignCreate) SetBaseCommissionRate(f float64) *CampaignCreate {
	cc.mutation.SetBaseCommissionRate(f)
	return cc
}

// SetNillableBaseCommissionRate sets the "base_commission_rate" field if the given value is not nil.
func (cc *CampaignCreate) SetNillableBaseCommissionRate(f *float64) *CampaignCreate {
	if f != nil {
		cc.SetBaseCommissionRate(*f)
	}
	return cc
}

// SetCommissionTiers sets the "commission_tiers" field.
func (cc *CampaignCreate) SetCommissionTiers(st []schema.CommissionTier) *CampaignCreate {
	cc.mutation.SetCommissionTiers(st)
	return cc
}

// SetTargetGeography sets the "target_geography" field.
func (cc *CampaignCreate) SetTargetGeography(s string) *CampaignCreate {
	cc.mutation.SetTargetGeography(s)
	return cc
}

// SetNillableTargetGeography sets the "target_geography" field if the given value is not nil.
func (cc *CampaignCreate) SetNillableTargetGeography(s *string) *CampaignCreate {
	if s != nil {
		cc.SetTargetGeography(*s)
	}
	return cc
}

// SetTargetDemographics sets the "target_demographics" field.
func (cc *CampaignCreate) SetTargetDemographics(m map[string]interface{}) *CampaignCreate {
	cc.mutation.SetTargetDemographics(m)
	return cc
}

// SetStartDate sets the "start_date" field.
func (cc *CampaignCreate) SetStartDate(t time.Time) *CampaignCreate {
	cc.mutation.SetStartDate(t)
	return cc
}

// SetEndDate sets the "end_date" field.
func (cc *CampaignCreate) SetEndDate(t time.Time) *CampaignCreate {
	cc.mutation.SetEndDate(t)
	return cc
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (cc *CampaignCreate) SetNillableEndDate(t *time.Time) *CampaignCreate {
	if t != nil {
		cc.SetEndDate(*t)
	}
	return cc
}

// SetStatus sets the "status" field.
func (cc *CampaignCreate) SetStatus(c campaign.Status) *CampaignCreate {
	cc.mutation.SetStatus(c)
	return cc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cc *CampaignCreate) SetNillableStatus(c *campaign.Status) *CampaignCreate {
	if c != nil {
		cc.SetStatus(*c)
	}
	return cc
}

// SetTrackingURL sets the "tracking_url" field.
func (cc *CampaignCreate) SetTrackingURL(s string) *CampaignCreate {
	cc.mutation.SetTrackingURL(s)
	return cc
}

// SetNillableTrackingURL sets the "tracking_url" field if the given value is not nil.
func (cc *CampaignCreate) SetNillableTrackingURL(s *string) *CampaignCreate {
	if s != nil {
		cc.SetTrackingURL(*s)
	}
	return cc
}

// SetTotalClicks sets the "total_clicks" field.
func (cc *CampaignCreate) SetTotalClicks(i int) *CampaignCreate {
	cc.mutation.SetTotalClicks(i)
	return cc
}

// SetNillableTotalClicks sets the "total_clicks" field if the given value is not nil.
func (cc *CampaignCreate) SetNillableTotalClicks(i *int) *CampaignCreate {
	if i != nil {
		cc.SetTotalClicks(*i)
	}
	return cc
}

// SetTotalConversions sets the "total_conversions" field.
func (cc *CampaignCreate) SetTotalConversions(i int) *CampaignCreate {
	cc.mutation.SetTotalConversions(i)
	return cc
}

// SetNillableTotalConversions sets the "total_conversions" field if the given value is not nil.
func (cc *CampaignCreate) SetNillableTotalConversions(i *int) *CampaignCreate {
	if i != nil {
		cc.SetTotalConversions(*i)
	}
	return cc
}

// SetTotalRevenue sets the "total_revenue" field.
func (cc *CampaignCreate) SetTotalRevenue(f float64) *CampaignCreate {
	cc.mutation.SetTotalRevenue(f)
	return cc
}

// SetNillableTotalRevenue sets the "total_revenue" field if the given value is not nil.
func (cc *CampaignCreate) SetNillableTotalRevenue(f *float64) *CampaignCreate {
	if f != nil {
		cc.SetTotalRevenue(*f)
	}
	return cc
}

// SetConversionRate sets the "conversion_rate" field.
func (cc *CampaignCreate) SetConversionRate(f float64) *CampaignCreate {
	cc.mutation.SetConversionRate(f)
	return cc
}

// SetNillableConversionRate sets the "conversion_rate" field if the given value is not nil.
func (cc *CampaignCreate) SetNillableConversionRate(f *float64) *CampaignCreate {
	if f != nil {
		cc.SetConversionRate(*f)
	}
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *CampaignCreate) SetCreatedAt(t time.Time) *CampaignCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CampaignCreate) SetNillableCreatedAt(t *time.Time) *CampaignCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CampaignCreate) SetUpdatedAt(t time.Time) *CampaignCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CampaignCreate) SetNillableUpdatedAt(t *time.Time) *CampaignCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CampaignCreate) SetID(i int64) *CampaignCreate {
	cc.mutation.SetID(i)
	return cc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cc *CampaignCreate) SetOwnerID(id int64) *CampaignCreate {
	cc.mutation.SetOwnerID(id)
	return cc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (cc *CampaignCreate) SetNillableOwnerID(id *int64) *CampaignCreate {
	if id != nil {
		cc = cc.SetOwnerID(*id)
	}
	return cc
}

// SetOwner sets the "owner" edge to the User entity.
func (cc *CampaignCreate) SetOwner(u *User) *CampaignCreate {
	return cc.SetOwnerID(u.ID)
}

// AddLinkIDs adds the "links" edge to the CampaignLink entity by IDs.
func (cc *CampaignCreate) AddLinkIDs(ids ...int64) *CampaignCreate {
	cc.mutation.AddLinkIDs(ids...)
	return cc
}

// AddLinks adds the "links" edges to the CampaignLink entity.
func (cc *CampaignCreate) AddLinks(c ...*CampaignLink) *CampaignCreate {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddLinkIDs(ids...)
}

// AddBannerIDs adds the "banners" edge to the Banner entity by IDs.
func (cc *CampaignCreate) AddBannerIDs(ids ...int64) *CampaignCreate {
	cc.mutation.AddBannerIDs(ids...)
	return cc
}

// AddBanners adds the "banners" edges to the Banner entity.
func (cc *CampaignCreate) AddBanners(b ...*Banner) *CampaignCreate {
	ids := make([]int64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cc.AddBannerIDs(ids...)
}

// Mutation returns the CampaignMutation object of the builder.
func (cc *CampaignCreate) Mutation() *CampaignMutation {
	return cc.mutation
}

// Save creates the Campaign in the database.
func (cc *CampaignCreate) Save(ctx context.Context) (*Campaign, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CampaignCreate) SaveX(ctx context.Context) *Campaign {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CampaignCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CampaignCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CampaignCreate) defaults() {
	if _, ok := cc.mutation.GetType(); !ok {
		v := campaign.DefaultType
		cc.mutation.SetType(v)
	}
	if _, ok := cc.mutation.CommissionType(); !ok {
		v := campaign.DefaultCommissionType
		cc.mutation.SetCommissionType(v)
	}
	if _, ok := cc.mutation.BaseCommissionRate(); !ok {
		v := campaign.DefaultBaseCommissionRate
		cc.mutation.SetBaseCommissionRate(v)
	}
	if _, ok := cc.mutation.Status(); !ok {
		v := campaign.DefaultStatus
		cc.mutation.SetStatus(v)
	}
	if _, ok := cc.mutation.TotalClicks(); !ok {
		v := campaign.DefaultTotalClicks
		cc.mutation.SetTotalClicks(v)
	}
	if _, ok := cc.mutation.TotalConversions(); !ok {
		v := campaign.DefaultTotalConversions
		cc.mutation.SetTotalConversions(v)
	}
	if _, ok := cc.mutation.TotalRevenue(); !ok {
		v := campaign.DefaultTotalRevenue
		cc.mutation.SetTotalRevenue(v)
	}
	if _, ok := cc.mutation.ConversionRate(); !ok {
		v := campaign.DefaultConversionRate
		cc.mutation.SetConversionRate(v)
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := campaign.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := campaign.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CampaignCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Campaign.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := campaign.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Campaign.name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.UniqueCode(); !ok {
		return &ValidationError{Name: "unique_code", err: errors.New(`ent: missing required field "Campaign.unique_code"`)}
	}
	if _, ok := cc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Campaign.type"`)}
	}
	if v, ok := cc.mutation.GetType(); ok {
		if err := campaign.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Campaign.type": %w`, err)}
		}
	}
	if _, ok := cc.mutation.CommissionType(); !ok {
		return &ValidationError{Name: "commission_type", err: errors.New(`ent: missing required field "Campaign.commission_type"`)}
	}
	if v, ok := cc.mutation.CommissionType(); ok {
		if err := campaign.CommissionTypeValidator(v); err != nil {
			return &ValidationError{Name: "commission_type", err: fmt.Errorf(`ent: validator failed for field "Campaign.commission_type": %w`, err)}
		}
	}
	if _, ok := cc.mutation.BaseCommissionRate(); !ok {
		return &ValidationError{Name: "base_commission_rate", err: errors.New(`ent: missing required field "Campaign.base_commission_rate"`)}
	}
	if _, ok := cc.mutation.StartDate(); !ok {
		return &ValidationError{Name: "start_date", err: errors.New(`ent: missing required field "Campaign.start_date"`)}
	}
	if _, ok := cc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Campaign.status"`)}
	}
	if v, ok := cc.mutation.Status(); ok {
		if err := campaign.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Campaign.status": %w`, err)}
		}
	}
	if _, ok := cc.mutation.TotalClicks(); !ok {
		return &ValidationError{Name: "total_clicks", err: errors.New(`ent: missing required field "Campaign.total_clicks"`)}
	}
	if _, ok := cc.mutation.TotalConversions(); !ok {
		return &ValidationError{Name: "total_conversions", err: errors.New(`ent: missing required field "Campaign.total_conversions"`)}
	}
	if _, ok := cc.mutation.TotalRevenue(); !ok {
		return &ValidationError{Name: "total_revenue", err: errors.New(`ent: missing required field "Campaign.total_revenue"`)}
	}
	if _, ok := cc.mutation.ConversionRate(); !ok {
		return &ValidationError{Name: "conversion_rate", err: errors.New(`ent: missing required field "Campaign.conversion_rate"`)}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Campaign.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Campaign.updated_at"`)}
	}
	if v, ok := cc.mutation.ID(); ok {
		if err := campaign.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Campaign.id": %w`, err)}
		}
	}
	return nil
}

func (cc *CampaignCreate) sqlSave(ctx context.Context) (*Campaign, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CampaignCreate) createSpec() (*Campaign, *sqlgraph.CreateSpec) {
	var (
		_node = &Campaign{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(campaign.Table, sqlgraph.NewFieldSpec(campaign.FieldID, field.TypeInt64))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(campaign.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Description(); ok {
		_spec.SetField(campaign.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := cc.mutation.UniqueCode(); ok {
		_spec.SetField(campaign.FieldUniqueCode, field.TypeString, value)
		_node.UniqueCode = value
	}
	if value, ok := cc.mutation.GetType(); ok {
		_spec.SetField(campaign.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := cc.mutation.CommissionType(); ok {
		_spec.SetField(campaign.FieldCommissionType, field.TypeEnum, value)
		_node.CommissionType = value
	}
	if value, ok := cc.mutation.BaseCommissionRate(); ok {
		_spec.SetField(campaign.FieldBaseCommissionRate, field.TypeFloat64, value)
		_node.BaseCommissionRate = value
	}
	if value, ok := cc.mutation.CommissionTiers(); ok {
		_spec.SetField(campaign.FieldCommissionTiers, field.TypeJSON, value)
		_node.CommissionTiers = value
	}
	if value, ok := cc.mutation.TargetGeography(); ok {
		_spec.SetField(campaign.FieldTargetGeography, field.TypeString, value)
		_node.TargetGeography = value
	}
	if value, ok := cc.mutation.TargetDemographics(); ok {
		_spec.SetField(campaign.FieldTargetDemographics, field.TypeJSON, value)
		_node.TargetDemographics = value
	}
	if value, ok := cc.mutation.StartDate(); ok {
		_spec.SetField(campaign.FieldStartDate, field.TypeTime, value)
		_node.StartDate = value
	}
	if value, ok := cc.mutation.EndDate(); ok {
		_spec.SetField(campaign.FieldEndDate, field.TypeTime, value)
		_node.EndDate = value
	}
	if value, ok := cc.mutation.Status(); ok {
		_spec.SetField(campaign.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := cc.mutation.TrackingURL(); ok {
		_spec.SetField(campaign.FieldTrackingURL, field.TypeString, value)
		_node.TrackingURL = value
	}
	if value, ok := cc.mutation.TotalClicks(); ok {
		_spec.SetField(campaign.FieldTotalClicks, field.TypeInt, value)
		_node.TotalClicks = value
	}
	if value, ok := cc.mutation.TotalConversions(); ok {
		_spec.SetField(campaign.FieldTotalConversions, field.TypeInt, value)
		_node.TotalConversions = value
	}
	if value, ok := cc.mutation.TotalRevenue(); ok {
		_spec.SetField(campaign.FieldTotalRevenue, field.TypeFloat64, value)
		_node.TotalRevenue = value
	}
	if value, ok := cc.mutation.ConversionRate(); ok {
		_spec.SetField(campaign.FieldConversionRate, field.TypeFloat64, value)
		_node.ConversionRate = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(campaign.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(campaign.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := cc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   campaign.OwnerTable,
			Columns: []string{campaign.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_campaigns = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.LinksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   campaign.LinksTable,
			Columns: []string{campaign.LinksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(campaignlink.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.BannersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   campaign.BannersTable,
			Columns: campaign.BannersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(banner.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CampaignCreateBulk is the builder for creating many Campaign entities in bulk.
type CampaignCreateBulk struct {
	config
	err      error
	builders []*CampaignCreate
}

// Save creates the Campaign entities in the database.
func (ccb *CampaignCreateBulk) Save(ctx context.Context) ([]*Campaign, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Campaign, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CampaignMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CampaignCreateBulk) SaveX(ctx context.Context) []*Campaign {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CampaignCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CampaignCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
