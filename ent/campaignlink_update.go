// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/campaign"
	"affluo/ent/campaignlink"
	"affluo/ent/predicate"
	"affluo/ent/track"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CampaignLinkUpdate is the builder for updating CampaignLink entities.
type CampaignLinkUpdate struct {
	config
	hooks    []Hook
	mutation *CampaignLinkMutation
}

// Where appends a list predicates to the CampaignLinkUpdate builder.
func (clu *CampaignLinkUpdate) Where(ps ...predicate.CampaignLink) *CampaignLinkUpdate {
	clu.mutation.Where(ps...)
	return clu
}

// SetUniqueCode sets the "unique_code" field.
func (clu *CampaignLinkUpdate) SetUniqueCode(s string) *CampaignLinkUpdate {
	clu.mutation.SetUniqueCode(s)
	return clu
}

// SetNillableUniqueCode sets the "unique_code" field if the given value is not nil.
func (clu *CampaignLinkUpdate) SetNillableUniqueCode(s *string) *CampaignLinkUpdate {
	if s != nil {
		clu.SetUniqueCode(*s)
	}
	return clu
}

// SetOriginalURL sets the "original_url" field.
func (clu *CampaignLinkUpdate) SetOriginalURL(s string) *CampaignLinkUpdate {
	clu.mutation.SetOriginalURL(s)
	return clu
}

// SetNillableOriginalURL sets the "original_url" field if the given value is not nil.
func (clu *CampaignLinkUpdate) SetNillableOriginalURL(s *string) *CampaignLinkUpdate {
	if s != nil {
		clu.SetOriginalURL(*s)
	}
	return clu
}

// SetTrackingURL sets the "tracking_url" field.
func (clu *CampaignLinkUpdate) SetTrackingURL(s string) *CampaignLinkUpdate {
	clu.mutation.SetTrackingURL(s)
	return clu
}

// SetNillableTrackingURL sets the "tracking_url" field if the given value is not nil.
func (clu *CampaignLinkUpdate) SetNillableTrackingURL(s *string) *CampaignLinkUpdate {
	if s != nil {
		clu.SetTrackingURL(*s)
	}
	return clu
}

// SetCreatedAt sets the "created_at" field.
func (clu *CampaignLinkUpdate) SetCreatedAt(t time.Time) *CampaignLinkUpdate {
	clu.mutation.SetCreatedAt(t)
	return clu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (clu *CampaignLinkUpdate) SetNillableCreatedAt(t *time.Time) *CampaignLinkUpdate {
	if t != nil {
		clu.SetCreatedAt(*t)
	}
	return clu
}

// SetIsActive sets the "is_active" field.
func (clu *CampaignLinkUpdate) SetIsActive(b bool) *CampaignLinkUpdate {
	clu.mutation.SetIsActive(b)
	return clu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (clu *CampaignLinkUpdate) SetNillableIsActive(b *bool) *CampaignLinkUpdate {
	if b != nil {
		clu.SetIsActive(*b)
	}
	return clu
}

// SetCampaignID sets the "campaign" edge to the Campaign entity by ID.
func (clu *CampaignLinkUpdate) SetCampaignID(id int64) *CampaignLinkUpdate {
	clu.mutation.SetCampaignID(id)
	return clu
}

// SetNillableCampaignID sets the "campaign" edge to the Campaign entity by ID if the given value is not nil.
func (clu *CampaignLinkUpdate) SetNillableCampaignID(id *int64) *CampaignLinkUpdate {
	if id != nil {
		clu = clu.SetCampaignID(*id)
	}
	return clu
}

// SetCampaign sets the "campaign" edge to the Campaign entity.
func (clu *CampaignLinkUpdate) SetCampaign(c *Campaign) *CampaignLinkUpdate {
	return clu.SetCampaignID(c.ID)
}

// AddTrackIDs adds the "tracks" edge to the Track entity by IDs.
func (clu *CampaignLinkUpdate) AddTrackIDs(ids ...int64) *CampaignLinkUpdate {
	clu.mutation.AddTrackIDs(ids...)
	return clu
}

// AddTracks adds the "tracks" edges to the Track entity.
func (clu *CampaignLinkUpdate) AddTracks(t ...*Track) *CampaignLinkUpdate {
	ids := make([]int64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return clu.AddTrackIDs(ids...)
}

// Mutation returns the CampaignLinkMutation object of the builder.
func (clu *CampaignLinkUpdate) Mutation() *CampaignLinkMutation {
	return clu.mutation
}

// ClearCampaign clears the "campaign" edge to the Campaign entity.
func (clu *CampaignLinkUpdate) ClearCampaign() *CampaignLinkUpdate {
	clu.mutation.ClearCampaign()
	return clu
}

// ClearTracks clears all "tracks" edges to the Track entity.
func (clu *CampaignLinkUpdate) ClearTracks() *CampaignLinkUpdate {
	clu.mutation.ClearTracks()
	return clu
}

// RemoveTrackIDs removes the "tracks" edge to Track entities by IDs.
func (clu *CampaignLinkUpdate) RemoveTrackIDs(ids ...int64) *CampaignLinkUpdate {
	clu.mutation.RemoveTrackIDs(ids...)
	return clu
}

// RemoveTracks removes "tracks" edges to Track entities.
func (clu *CampaignLinkUpdate) RemoveTracks(t ...*Track) *CampaignLinkUpdate {
	ids := make([]int64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return clu.RemoveTrackIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (clu *CampaignLinkUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, clu.sqlSave, clu.mutation, clu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (clu *CampaignLinkUpdate) SaveX(ctx context.Context) int {
	affected, err := clu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (clu *CampaignLinkUpdate) Exec(ctx context.Context) error {
	_, err := clu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (clu *CampaignLinkUpdate) ExecX(ctx context.Context) {
	if err := clu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (clu *CampaignLinkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(campaignlink.Table, campaignlink.Columns, sqlgraph.NewFieldSpec(campaignlink.FieldID, field.TypeInt64))
	if ps := clu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := clu.mutation.UniqueCode(); ok {
		_spec.SetField(campaignlink.FieldUniqueCode, field.TypeString, value)
	}
	if value, ok := clu.mutation.OriginalURL(); ok {
		_spec.SetField(campaignlink.FieldOriginalURL, field.TypeString, value)
	}
	if value, ok := clu.mutation.TrackingURL(); ok {
		_spec.SetField(campaignlink.FieldTrackingURL, field.TypeString, value)
	}
	if value, ok := clu.mutation.CreatedAt(); ok {
		_spec.SetField(campaignlink.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := clu.mutation.IsActive(); ok {
		_spec.SetField(campaignlink.FieldIsActive, field.TypeBool, value)
	}
	if clu.mutation.CampaignCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   campaignlink.CampaignTable,
			Columns: []string{campaignlink.CampaignColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(campaign.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := clu.mutation.CampaignIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   campaignlink.CampaignTable,
			Columns: []string{campaignlink.CampaignColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(campaign.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if clu.mutation.TracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   campaignlink.TracksTable,
			Columns: []string{campaignlink.TracksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := clu.mutation.RemovedTracksIDs(); len(nodes) > 0 && !clu.mutation.TracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   campaignlink.TracksTable,
			Columns: []string{campaignlink.TracksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := clu.mutation.TracksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   campaignlink.TracksTable,
			Columns: []string{campaignlink.TracksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, clu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{campaignlink.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	clu.mutation.done = true
	return n, nil
}

// CampaignLinkUpdateOne is the builder for updating a single CampaignLink entity.
type CampaignLinkUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CampaignLinkMutation
}

// SetUniqueCode sets the "unique_code" field.
func (cluo *CampaignLinkUpdateOne) SetUniqueCode(s string) *CampaignLinkUpdateOne {
	cluo.mutation.SetUniqueCode(s)
	return cluo
}

// SetNillableUniqueCode sets the "unique_code" field if the given value is not nil.
func (cluo *CampaignLinkUpdateOne) SetNillableUniqueCode(s *string) *CampaignLinkUpdateOne {
	if s != nil {
		cluo.SetUniqueCode(*s)
	}
	return cluo
}

// SetOriginalURL sets the "original_url" field.
func (cluo *CampaignLinkUpdateOne) SetOriginalURL(s string) *CampaignLinkUpdateOne {
	cluo.mutation.SetOriginalURL(s)
	return cluo
}

// SetNillableOriginalURL sets the "original_url" field if the given value is not nil.
func (cluo *CampaignLinkUpdateOne) SetNillableOriginalURL(s *string) *CampaignLinkUpdateOne {
	if s != nil {
		cluo.SetOriginalURL(*s)
	}
	return cluo
}

// SetTrackingURL sets the "tracking_url" field.
func (cluo *CampaignLinkUpdateOne) SetTrackingURL(s string) *CampaignLinkUpdateOne {
	cluo.mutation.SetTrackingURL(s)
	return cluo
}

// SetNillableTrackingURL sets the "tracking_url" field if the given value is not nil.
func (cluo *CampaignLinkUpdateOne) SetNillableTrackingURL(s *string) *CampaignLinkUpdateOne {
	if s != nil {
		cluo.SetTrackingURL(*s)
	}
	return cluo
}

// SetCreatedAt sets the "created_at" field.
func (cluo *CampaignLinkUpdateOne) SetCreatedAt(t time.Time) *CampaignLinkUpdateOne {
	cluo.mutation.SetCreatedAt(t)
	return cluo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cluo *CampaignLinkUpdateOne) SetNillableCreatedAt(t *time.Time) *CampaignLinkUpdateOne {
	if t != nil {
		cluo.SetCreatedAt(*t)
	}
	return cluo
}

// SetIsActive sets the "is_active" field.
func (cluo *CampaignLinkUpdateOne) SetIsActive(b bool) *CampaignLinkUpdateOne {
	cluo.mutation.SetIsActive(b)
	return cluo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (cluo *CampaignLinkUpdateOne) SetNillableIsActive(b *bool) *CampaignLinkUpdateOne {
	if b != nil {
		cluo.SetIsActive(*b)
	}
	return cluo
}

// SetCampaignID sets the "campaign" edge to the Campaign entity by ID.
func (cluo *CampaignLinkUpdateOne) SetCampaignID(id int64) *CampaignLinkUpdateOne {
	cluo.mutation.SetCampaignID(id)
	return cluo
}

// SetNillableCampaignID sets the "campaign" edge to the Campaign entity by ID if the given value is not nil.
func (cluo *CampaignLinkUpdateOne) SetNillableCampaignID(id *int64) *CampaignLinkUpdateOne {
	if id != nil {
		cluo = cluo.SetCampaignID(*id)
	}
	return cluo
}

// SetCampaign sets the "campaign" edge to the Campaign entity.
func (cluo *CampaignLinkUpdateOne) SetCampaign(c *Campaign) *CampaignLinkUpdateOne {
	return cluo.SetCampaignID(c.ID)
}

// AddTrackIDs adds the "tracks" edge to the Track entity by IDs.
func (cluo *CampaignLinkUpdateOne) AddTrackIDs(ids ...int64) *CampaignLinkUpdateOne {
	cluo.mutation.AddTrackIDs(ids...)
	return cluo
}

// AddTracks adds the "tracks" edges to the Track entity.
func (cluo *CampaignLinkUpdateOne) AddTracks(t ...*Track) *CampaignLinkUpdateOne {
	ids := make([]int64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cluo.AddTrackIDs(ids...)
}

// Mutation returns the CampaignLinkMutation object of the builder.
func (cluo *CampaignLinkUpdateOne) Mutation() *CampaignLinkMutation {
	return cluo.mutation
}

// ClearCampaign clears the "campaign" edge to the Campaign entity.
func (cluo *CampaignLinkUpdateOne) ClearCampaign() *CampaignLinkUpdateOne {
	cluo.mutation.ClearCampaign()
	return cluo
}

// ClearTracks clears all "tracks" edges to the Track entity.
func (cluo *CampaignLinkUpdateOne) ClearTracks() *CampaignLinkUpdateOne {
	cluo.mutation.ClearTracks()
	return cluo
}

// RemoveTrackIDs removes the "tracks" edge to Track entities by IDs.
func (cluo *CampaignLinkUpdateOne) RemoveTrackIDs(ids ...int64) *CampaignLinkUpdateOne {
	cluo.mutation.RemoveTrackIDs(ids...)
	return cluo
}

// RemoveTracks removes "tracks" edges to Track entities.
func (cluo *CampaignLinkUpdateOne) RemoveTracks(t ...*Track) *CampaignLinkUpdateOne {
	ids := make([]int64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cluo.RemoveTrackIDs(ids...)
}

// Where appends a list predicates to the CampaignLinkUpdate builder.
func (cluo *CampaignLinkUpdateOne) Where(ps ...predicate.CampaignLink) *CampaignLinkUpdateOne {
	cluo.mutation.Where(ps...)
	return cluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cluo *CampaignLinkUpdateOne) Select(field string, fields ...string) *CampaignLinkUpdateOne {
	cluo.fields = append([]string{field}, fields...)
	return cluo
}

// Save executes the query and returns the updated CampaignLink entity.
func (cluo *CampaignLinkUpdateOne) Save(ctx context.Context) (*CampaignLink, error) {
	return withHooks(ctx, cluo.sqlSave, cluo.mutation, cluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cluo *CampaignLinkUpdateOne) SaveX(ctx context.Context) *CampaignLink {
	node, err := cluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cluo *CampaignLinkUpdateOne) Exec(ctx context.Context) error {
	_, err := cluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cluo *CampaignLinkUpdateOne) ExecX(ctx context.Context) {
	if err := cluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cluo *CampaignLinkUpdateOne) sqlSave(ctx context.Context) (_node *CampaignLink, err error) {
	_spec := sqlgraph.NewUpdateSpec(campaignlink.Table, campaignlink.Columns, sqlgraph.NewFieldSpec(campaignlink.FieldID, field.TypeInt64))
	id, ok := cluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CampaignLink.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, campaignlink.FieldID)
		for _, f := range fields {
			if !campaignlink.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != campaignlink.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cluo.mutation.UniqueCode(); ok {
		_spec.SetField(campaignlink.FieldUniqueCode, field.TypeString, value)
	}
	if value, ok := cluo.mutation.OriginalURL(); ok {
		_spec.SetField(campaignlink.FieldOriginalURL, field.TypeString, value)
	}
	if value, ok := cluo.mutation.TrackingURL(); ok {
		_spec.SetField(campaignlink.FieldTrackingURL, field.TypeString, value)
	}
	if value, ok := cluo.mutation.CreatedAt(); ok {
		_spec.SetField(campaignlink.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cluo.mutation.IsActive(); ok {
		_spec.SetField(campaignlink.FieldIsActive, field.TypeBool, value)
	}
	if cluo.mutation.CampaignCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   campaignlink.CampaignTable,
			Columns: []string{campaignlink.CampaignColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(campaign.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cluo.mutation.CampaignIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   campaignlink.CampaignTable,
			Columns: []string{campaignlink.CampaignColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(campaign.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cluo.mutation.TracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   campaignlink.TracksTable,
			Columns: []string{campaignlink.TracksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cluo.mutation.RemovedTracksIDs(); len(nodes) > 0 && !cluo.mutation.TracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   campaignlink.TracksTable,
			Columns: []string{campaignlink.TracksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cluo.mutation.TracksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   campaignlink.TracksTable,
			Columns: []string{campaignlink.TracksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CampaignLink{config: cluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{campaignlink.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cluo.mutation.done = true
	return _node, nil
}
