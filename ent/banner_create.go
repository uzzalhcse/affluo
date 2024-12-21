// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/banner"
	"affluo/ent/bannercreative"
	"affluo/ent/bannerstats"
	"affluo/ent/campaign"
	"affluo/ent/lead"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BannerCreate is the builder for creating a Banner entity.
type BannerCreate struct {
	config
	mutation *BannerMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (bc *BannerCreate) SetName(s string) *BannerCreate {
	bc.mutation.SetName(s)
	return bc
}

// SetDescription sets the "description" field.
func (bc *BannerCreate) SetDescription(s string) *BannerCreate {
	bc.mutation.SetDescription(s)
	return bc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (bc *BannerCreate) SetNillableDescription(s *string) *BannerCreate {
	if s != nil {
		bc.SetDescription(*s)
	}
	return bc
}

// SetType sets the "type" field.
func (bc *BannerCreate) SetType(b banner.Type) *BannerCreate {
	bc.mutation.SetType(b)
	return bc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (bc *BannerCreate) SetNillableType(b *banner.Type) *BannerCreate {
	if b != nil {
		bc.SetType(*b)
	}
	return bc
}

// SetClickURL sets the "click_url" field.
func (bc *BannerCreate) SetClickURL(s string) *BannerCreate {
	bc.mutation.SetClickURL(s)
	return bc
}

// SetNillableClickURL sets the "click_url" field if the given value is not nil.
func (bc *BannerCreate) SetNillableClickURL(s *string) *BannerCreate {
	if s != nil {
		bc.SetClickURL(*s)
	}
	return bc
}

// SetSize sets the "size" field.
func (bc *BannerCreate) SetSize(s string) *BannerCreate {
	bc.mutation.SetSize(s)
	return bc
}

// SetStatus sets the "status" field.
func (bc *BannerCreate) SetStatus(b banner.Status) *BannerCreate {
	bc.mutation.SetStatus(b)
	return bc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (bc *BannerCreate) SetNillableStatus(b *banner.Status) *BannerCreate {
	if b != nil {
		bc.SetStatus(*b)
	}
	return bc
}

// SetAllowedCountries sets the "allowed_countries" field.
func (bc *BannerCreate) SetAllowedCountries(s []string) *BannerCreate {
	bc.mutation.SetAllowedCountries(s)
	return bc
}

// SetWeight sets the "weight" field.
func (bc *BannerCreate) SetWeight(i int) *BannerCreate {
	bc.mutation.SetWeight(i)
	return bc
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (bc *BannerCreate) SetNillableWeight(i *int) *BannerCreate {
	if i != nil {
		bc.SetWeight(*i)
	}
	return bc
}

// SetSmartWeight sets the "smart_weight" field.
func (bc *BannerCreate) SetSmartWeight(f float64) *BannerCreate {
	bc.mutation.SetSmartWeight(f)
	return bc
}

// SetNillableSmartWeight sets the "smart_weight" field if the given value is not nil.
func (bc *BannerCreate) SetNillableSmartWeight(f *float64) *BannerCreate {
	if f != nil {
		bc.SetSmartWeight(*f)
	}
	return bc
}

// SetLastImpression sets the "last_impression" field.
func (bc *BannerCreate) SetLastImpression(t time.Time) *BannerCreate {
	bc.mutation.SetLastImpression(t)
	return bc
}

// SetNillableLastImpression sets the "last_impression" field if the given value is not nil.
func (bc *BannerCreate) SetNillableLastImpression(t *time.Time) *BannerCreate {
	if t != nil {
		bc.SetLastImpression(*t)
	}
	return bc
}

// SetStartDate sets the "start_date" field.
func (bc *BannerCreate) SetStartDate(t time.Time) *BannerCreate {
	bc.mutation.SetStartDate(t)
	return bc
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (bc *BannerCreate) SetNillableStartDate(t *time.Time) *BannerCreate {
	if t != nil {
		bc.SetStartDate(*t)
	}
	return bc
}

// SetEndDate sets the "end_date" field.
func (bc *BannerCreate) SetEndDate(t time.Time) *BannerCreate {
	bc.mutation.SetEndDate(t)
	return bc
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (bc *BannerCreate) SetNillableEndDate(t *time.Time) *BannerCreate {
	if t != nil {
		bc.SetEndDate(*t)
	}
	return bc
}

// SetAllowedDevices sets the "allowed_devices" field.
func (bc *BannerCreate) SetAllowedDevices(s []string) *BannerCreate {
	bc.mutation.SetAllowedDevices(s)
	return bc
}

// SetAllowedBrowsers sets the "allowed_browsers" field.
func (bc *BannerCreate) SetAllowedBrowsers(s []string) *BannerCreate {
	bc.mutation.SetAllowedBrowsers(s)
	return bc
}

// SetAllowedOs sets the "allowed_os" field.
func (bc *BannerCreate) SetAllowedOs(s []string) *BannerCreate {
	bc.mutation.SetAllowedOs(s)
	return bc
}

// SetCreatedAt sets the "created_at" field.
func (bc *BannerCreate) SetCreatedAt(t time.Time) *BannerCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bc *BannerCreate) SetNillableCreatedAt(t *time.Time) *BannerCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// SetUpdatedAt sets the "updated_at" field.
func (bc *BannerCreate) SetUpdatedAt(t time.Time) *BannerCreate {
	bc.mutation.SetUpdatedAt(t)
	return bc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bc *BannerCreate) SetNillableUpdatedAt(t *time.Time) *BannerCreate {
	if t != nil {
		bc.SetUpdatedAt(*t)
	}
	return bc
}

// SetID sets the "id" field.
func (bc *BannerCreate) SetID(i int64) *BannerCreate {
	bc.mutation.SetID(i)
	return bc
}

// AddCampaignIDs adds the "campaigns" edge to the Campaign entity by IDs.
func (bc *BannerCreate) AddCampaignIDs(ids ...int64) *BannerCreate {
	bc.mutation.AddCampaignIDs(ids...)
	return bc
}

// AddCampaigns adds the "campaigns" edges to the Campaign entity.
func (bc *BannerCreate) AddCampaigns(c ...*Campaign) *BannerCreate {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bc.AddCampaignIDs(ids...)
}

// AddCreativeIDs adds the "creatives" edge to the BannerCreative entity by IDs.
func (bc *BannerCreate) AddCreativeIDs(ids ...int64) *BannerCreate {
	bc.mutation.AddCreativeIDs(ids...)
	return bc
}

// AddCreatives adds the "creatives" edges to the BannerCreative entity.
func (bc *BannerCreate) AddCreatives(b ...*BannerCreative) *BannerCreate {
	ids := make([]int64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bc.AddCreativeIDs(ids...)
}

// AddStatIDs adds the "stats" edge to the BannerStats entity by IDs.
func (bc *BannerCreate) AddStatIDs(ids ...int64) *BannerCreate {
	bc.mutation.AddStatIDs(ids...)
	return bc
}

// AddStats adds the "stats" edges to the BannerStats entity.
func (bc *BannerCreate) AddStats(b ...*BannerStats) *BannerCreate {
	ids := make([]int64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bc.AddStatIDs(ids...)
}

// AddLeadIDs adds the "leads" edge to the Lead entity by IDs.
func (bc *BannerCreate) AddLeadIDs(ids ...int64) *BannerCreate {
	bc.mutation.AddLeadIDs(ids...)
	return bc
}

// AddLeads adds the "leads" edges to the Lead entity.
func (bc *BannerCreate) AddLeads(l ...*Lead) *BannerCreate {
	ids := make([]int64, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return bc.AddLeadIDs(ids...)
}

// Mutation returns the BannerMutation object of the builder.
func (bc *BannerCreate) Mutation() *BannerMutation {
	return bc.mutation
}

// Save creates the Banner in the database.
func (bc *BannerCreate) Save(ctx context.Context) (*Banner, error) {
	bc.defaults()
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BannerCreate) SaveX(ctx context.Context) *Banner {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BannerCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BannerCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BannerCreate) defaults() {
	if _, ok := bc.mutation.GetType(); !ok {
		v := banner.DefaultType
		bc.mutation.SetType(v)
	}
	if _, ok := bc.mutation.Status(); !ok {
		v := banner.DefaultStatus
		bc.mutation.SetStatus(v)
	}
	if _, ok := bc.mutation.Weight(); !ok {
		v := banner.DefaultWeight
		bc.mutation.SetWeight(v)
	}
	if _, ok := bc.mutation.CreatedAt(); !ok {
		v := banner.DefaultCreatedAt()
		bc.mutation.SetCreatedAt(v)
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		v := banner.DefaultUpdatedAt()
		bc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BannerCreate) check() error {
	if _, ok := bc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Banner.name"`)}
	}
	if v, ok := bc.mutation.Name(); ok {
		if err := banner.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Banner.name": %w`, err)}
		}
	}
	if _, ok := bc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Banner.type"`)}
	}
	if v, ok := bc.mutation.GetType(); ok {
		if err := banner.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Banner.type": %w`, err)}
		}
	}
	if _, ok := bc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "Banner.size"`)}
	}
	if _, ok := bc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Banner.status"`)}
	}
	if v, ok := bc.mutation.Status(); ok {
		if err := banner.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Banner.status": %w`, err)}
		}
	}
	if _, ok := bc.mutation.Weight(); !ok {
		return &ValidationError{Name: "weight", err: errors.New(`ent: missing required field "Banner.weight"`)}
	}
	if _, ok := bc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Banner.created_at"`)}
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Banner.updated_at"`)}
	}
	if v, ok := bc.mutation.ID(); ok {
		if err := banner.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Banner.id": %w`, err)}
		}
	}
	return nil
}

func (bc *BannerCreate) sqlSave(ctx context.Context) (*Banner, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BannerCreate) createSpec() (*Banner, *sqlgraph.CreateSpec) {
	var (
		_node = &Banner{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(banner.Table, sqlgraph.NewFieldSpec(banner.FieldID, field.TypeInt64))
	)
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bc.mutation.Name(); ok {
		_spec.SetField(banner.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := bc.mutation.Description(); ok {
		_spec.SetField(banner.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := bc.mutation.GetType(); ok {
		_spec.SetField(banner.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := bc.mutation.ClickURL(); ok {
		_spec.SetField(banner.FieldClickURL, field.TypeString, value)
		_node.ClickURL = value
	}
	if value, ok := bc.mutation.Size(); ok {
		_spec.SetField(banner.FieldSize, field.TypeString, value)
		_node.Size = value
	}
	if value, ok := bc.mutation.Status(); ok {
		_spec.SetField(banner.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := bc.mutation.AllowedCountries(); ok {
		_spec.SetField(banner.FieldAllowedCountries, field.TypeJSON, value)
		_node.AllowedCountries = value
	}
	if value, ok := bc.mutation.Weight(); ok {
		_spec.SetField(banner.FieldWeight, field.TypeInt, value)
		_node.Weight = value
	}
	if value, ok := bc.mutation.SmartWeight(); ok {
		_spec.SetField(banner.FieldSmartWeight, field.TypeFloat64, value)
		_node.SmartWeight = value
	}
	if value, ok := bc.mutation.LastImpression(); ok {
		_spec.SetField(banner.FieldLastImpression, field.TypeTime, value)
		_node.LastImpression = value
	}
	if value, ok := bc.mutation.StartDate(); ok {
		_spec.SetField(banner.FieldStartDate, field.TypeTime, value)
		_node.StartDate = value
	}
	if value, ok := bc.mutation.EndDate(); ok {
		_spec.SetField(banner.FieldEndDate, field.TypeTime, value)
		_node.EndDate = value
	}
	if value, ok := bc.mutation.AllowedDevices(); ok {
		_spec.SetField(banner.FieldAllowedDevices, field.TypeJSON, value)
		_node.AllowedDevices = value
	}
	if value, ok := bc.mutation.AllowedBrowsers(); ok {
		_spec.SetField(banner.FieldAllowedBrowsers, field.TypeJSON, value)
		_node.AllowedBrowsers = value
	}
	if value, ok := bc.mutation.AllowedOs(); ok {
		_spec.SetField(banner.FieldAllowedOs, field.TypeJSON, value)
		_node.AllowedOs = value
	}
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.SetField(banner.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := bc.mutation.UpdatedAt(); ok {
		_spec.SetField(banner.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := bc.mutation.CampaignsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   banner.CampaignsTable,
			Columns: banner.CampaignsPrimaryKey,
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
	if nodes := bc.mutation.CreativesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   banner.CreativesTable,
			Columns: []string{banner.CreativesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bannercreative.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.StatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   banner.StatsTable,
			Columns: []string{banner.StatsColumn},
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
	if nodes := bc.mutation.LeadsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   banner.LeadsTable,
			Columns: []string{banner.LeadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lead.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BannerCreateBulk is the builder for creating many Banner entities in bulk.
type BannerCreateBulk struct {
	config
	err      error
	builders []*BannerCreate
}

// Save creates the Banner entities in the database.
func (bcb *BannerCreateBulk) Save(ctx context.Context) ([]*Banner, error) {
	if bcb.err != nil {
		return nil, bcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Banner, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BannerMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BannerCreateBulk) SaveX(ctx context.Context) []*Banner {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BannerCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BannerCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
