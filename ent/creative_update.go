// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/banner"
	"affluo/ent/bannercreative"
	"affluo/ent/creative"
	"affluo/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CreativeUpdate is the builder for updating Creative entities.
type CreativeUpdate struct {
	config
	hooks    []Hook
	mutation *CreativeMutation
}

// Where appends a list predicates to the CreativeUpdate builder.
func (cu *CreativeUpdate) Where(ps ...predicate.Creative) *CreativeUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CreativeUpdate) SetName(s string) *CreativeUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *CreativeUpdate) SetNillableName(s *string) *CreativeUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// ClearName clears the value of the "name" field.
func (cu *CreativeUpdate) ClearName() *CreativeUpdate {
	cu.mutation.ClearName()
	return cu
}

// SetImageURL sets the "image_url" field.
func (cu *CreativeUpdate) SetImageURL(s string) *CreativeUpdate {
	cu.mutation.SetImageURL(s)
	return cu
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (cu *CreativeUpdate) SetNillableImageURL(s *string) *CreativeUpdate {
	if s != nil {
		cu.SetImageURL(*s)
	}
	return cu
}

// ClearImageURL clears the value of the "image_url" field.
func (cu *CreativeUpdate) ClearImageURL() *CreativeUpdate {
	cu.mutation.ClearImageURL()
	return cu
}

// SetSize sets the "size" field.
func (cu *CreativeUpdate) SetSize(s string) *CreativeUpdate {
	cu.mutation.SetSize(s)
	return cu
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (cu *CreativeUpdate) SetNillableSize(s *string) *CreativeUpdate {
	if s != nil {
		cu.SetSize(*s)
	}
	return cu
}

// ClearSize clears the value of the "size" field.
func (cu *CreativeUpdate) ClearSize() *CreativeUpdate {
	cu.mutation.ClearSize()
	return cu
}

// SetEnabled sets the "enabled" field.
func (cu *CreativeUpdate) SetEnabled(b bool) *CreativeUpdate {
	cu.mutation.SetEnabled(b)
	return cu
}

// SetNillableEnabled sets the "enabled" field if the given value is not nil.
func (cu *CreativeUpdate) SetNillableEnabled(b *bool) *CreativeUpdate {
	if b != nil {
		cu.SetEnabled(*b)
	}
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CreativeUpdate) SetUpdatedAt(t time.Time) *CreativeUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// AddBannerIDs adds the "banners" edge to the Banner entity by IDs.
func (cu *CreativeUpdate) AddBannerIDs(ids ...int64) *CreativeUpdate {
	cu.mutation.AddBannerIDs(ids...)
	return cu
}

// AddBanners adds the "banners" edges to the Banner entity.
func (cu *CreativeUpdate) AddBanners(b ...*Banner) *CreativeUpdate {
	ids := make([]int64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cu.AddBannerIDs(ids...)
}

// AddBannerCreativeIDs adds the "banner_creatives" edge to the BannerCreative entity by IDs.
func (cu *CreativeUpdate) AddBannerCreativeIDs(ids ...int) *CreativeUpdate {
	cu.mutation.AddBannerCreativeIDs(ids...)
	return cu
}

// AddBannerCreatives adds the "banner_creatives" edges to the BannerCreative entity.
func (cu *CreativeUpdate) AddBannerCreatives(b ...*BannerCreative) *CreativeUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cu.AddBannerCreativeIDs(ids...)
}

// Mutation returns the CreativeMutation object of the builder.
func (cu *CreativeUpdate) Mutation() *CreativeMutation {
	return cu.mutation
}

// ClearBanners clears all "banners" edges to the Banner entity.
func (cu *CreativeUpdate) ClearBanners() *CreativeUpdate {
	cu.mutation.ClearBanners()
	return cu
}

// RemoveBannerIDs removes the "banners" edge to Banner entities by IDs.
func (cu *CreativeUpdate) RemoveBannerIDs(ids ...int64) *CreativeUpdate {
	cu.mutation.RemoveBannerIDs(ids...)
	return cu
}

// RemoveBanners removes "banners" edges to Banner entities.
func (cu *CreativeUpdate) RemoveBanners(b ...*Banner) *CreativeUpdate {
	ids := make([]int64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cu.RemoveBannerIDs(ids...)
}

// ClearBannerCreatives clears all "banner_creatives" edges to the BannerCreative entity.
func (cu *CreativeUpdate) ClearBannerCreatives() *CreativeUpdate {
	cu.mutation.ClearBannerCreatives()
	return cu
}

// RemoveBannerCreativeIDs removes the "banner_creatives" edge to BannerCreative entities by IDs.
func (cu *CreativeUpdate) RemoveBannerCreativeIDs(ids ...int) *CreativeUpdate {
	cu.mutation.RemoveBannerCreativeIDs(ids...)
	return cu
}

// RemoveBannerCreatives removes "banner_creatives" edges to BannerCreative entities.
func (cu *CreativeUpdate) RemoveBannerCreatives(b ...*BannerCreative) *CreativeUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cu.RemoveBannerCreativeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CreativeUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CreativeUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CreativeUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CreativeUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CreativeUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := creative.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

func (cu *CreativeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(creative.Table, creative.Columns, sqlgraph.NewFieldSpec(creative.FieldID, field.TypeInt64))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(creative.FieldName, field.TypeString, value)
	}
	if cu.mutation.NameCleared() {
		_spec.ClearField(creative.FieldName, field.TypeString)
	}
	if value, ok := cu.mutation.ImageURL(); ok {
		_spec.SetField(creative.FieldImageURL, field.TypeString, value)
	}
	if cu.mutation.ImageURLCleared() {
		_spec.ClearField(creative.FieldImageURL, field.TypeString)
	}
	if value, ok := cu.mutation.Size(); ok {
		_spec.SetField(creative.FieldSize, field.TypeString, value)
	}
	if cu.mutation.SizeCleared() {
		_spec.ClearField(creative.FieldSize, field.TypeString)
	}
	if value, ok := cu.mutation.Enabled(); ok {
		_spec.SetField(creative.FieldEnabled, field.TypeBool, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(creative.FieldUpdatedAt, field.TypeTime, value)
	}
	if cu.mutation.BannersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   creative.BannersTable,
			Columns: creative.BannersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(banner.FieldID, field.TypeInt64),
			},
		}
		createE := &BannerCreativeCreate{config: cu.config, mutation: newBannerCreativeMutation(cu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedBannersIDs(); len(nodes) > 0 && !cu.mutation.BannersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   creative.BannersTable,
			Columns: creative.BannersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(banner.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &BannerCreativeCreate{config: cu.config, mutation: newBannerCreativeMutation(cu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.BannersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   creative.BannersTable,
			Columns: creative.BannersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(banner.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &BannerCreativeCreate{config: cu.config, mutation: newBannerCreativeMutation(cu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.BannerCreativesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   creative.BannerCreativesTable,
			Columns: []string{creative.BannerCreativesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bannercreative.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedBannerCreativesIDs(); len(nodes) > 0 && !cu.mutation.BannerCreativesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   creative.BannerCreativesTable,
			Columns: []string{creative.BannerCreativesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bannercreative.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.BannerCreativesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   creative.BannerCreativesTable,
			Columns: []string{creative.BannerCreativesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bannercreative.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{creative.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CreativeUpdateOne is the builder for updating a single Creative entity.
type CreativeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CreativeMutation
}

// SetName sets the "name" field.
func (cuo *CreativeUpdateOne) SetName(s string) *CreativeUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *CreativeUpdateOne) SetNillableName(s *string) *CreativeUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// ClearName clears the value of the "name" field.
func (cuo *CreativeUpdateOne) ClearName() *CreativeUpdateOne {
	cuo.mutation.ClearName()
	return cuo
}

// SetImageURL sets the "image_url" field.
func (cuo *CreativeUpdateOne) SetImageURL(s string) *CreativeUpdateOne {
	cuo.mutation.SetImageURL(s)
	return cuo
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (cuo *CreativeUpdateOne) SetNillableImageURL(s *string) *CreativeUpdateOne {
	if s != nil {
		cuo.SetImageURL(*s)
	}
	return cuo
}

// ClearImageURL clears the value of the "image_url" field.
func (cuo *CreativeUpdateOne) ClearImageURL() *CreativeUpdateOne {
	cuo.mutation.ClearImageURL()
	return cuo
}

// SetSize sets the "size" field.
func (cuo *CreativeUpdateOne) SetSize(s string) *CreativeUpdateOne {
	cuo.mutation.SetSize(s)
	return cuo
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (cuo *CreativeUpdateOne) SetNillableSize(s *string) *CreativeUpdateOne {
	if s != nil {
		cuo.SetSize(*s)
	}
	return cuo
}

// ClearSize clears the value of the "size" field.
func (cuo *CreativeUpdateOne) ClearSize() *CreativeUpdateOne {
	cuo.mutation.ClearSize()
	return cuo
}

// SetEnabled sets the "enabled" field.
func (cuo *CreativeUpdateOne) SetEnabled(b bool) *CreativeUpdateOne {
	cuo.mutation.SetEnabled(b)
	return cuo
}

// SetNillableEnabled sets the "enabled" field if the given value is not nil.
func (cuo *CreativeUpdateOne) SetNillableEnabled(b *bool) *CreativeUpdateOne {
	if b != nil {
		cuo.SetEnabled(*b)
	}
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CreativeUpdateOne) SetUpdatedAt(t time.Time) *CreativeUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// AddBannerIDs adds the "banners" edge to the Banner entity by IDs.
func (cuo *CreativeUpdateOne) AddBannerIDs(ids ...int64) *CreativeUpdateOne {
	cuo.mutation.AddBannerIDs(ids...)
	return cuo
}

// AddBanners adds the "banners" edges to the Banner entity.
func (cuo *CreativeUpdateOne) AddBanners(b ...*Banner) *CreativeUpdateOne {
	ids := make([]int64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cuo.AddBannerIDs(ids...)
}

// AddBannerCreativeIDs adds the "banner_creatives" edge to the BannerCreative entity by IDs.
func (cuo *CreativeUpdateOne) AddBannerCreativeIDs(ids ...int) *CreativeUpdateOne {
	cuo.mutation.AddBannerCreativeIDs(ids...)
	return cuo
}

// AddBannerCreatives adds the "banner_creatives" edges to the BannerCreative entity.
func (cuo *CreativeUpdateOne) AddBannerCreatives(b ...*BannerCreative) *CreativeUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cuo.AddBannerCreativeIDs(ids...)
}

// Mutation returns the CreativeMutation object of the builder.
func (cuo *CreativeUpdateOne) Mutation() *CreativeMutation {
	return cuo.mutation
}

// ClearBanners clears all "banners" edges to the Banner entity.
func (cuo *CreativeUpdateOne) ClearBanners() *CreativeUpdateOne {
	cuo.mutation.ClearBanners()
	return cuo
}

// RemoveBannerIDs removes the "banners" edge to Banner entities by IDs.
func (cuo *CreativeUpdateOne) RemoveBannerIDs(ids ...int64) *CreativeUpdateOne {
	cuo.mutation.RemoveBannerIDs(ids...)
	return cuo
}

// RemoveBanners removes "banners" edges to Banner entities.
func (cuo *CreativeUpdateOne) RemoveBanners(b ...*Banner) *CreativeUpdateOne {
	ids := make([]int64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cuo.RemoveBannerIDs(ids...)
}

// ClearBannerCreatives clears all "banner_creatives" edges to the BannerCreative entity.
func (cuo *CreativeUpdateOne) ClearBannerCreatives() *CreativeUpdateOne {
	cuo.mutation.ClearBannerCreatives()
	return cuo
}

// RemoveBannerCreativeIDs removes the "banner_creatives" edge to BannerCreative entities by IDs.
func (cuo *CreativeUpdateOne) RemoveBannerCreativeIDs(ids ...int) *CreativeUpdateOne {
	cuo.mutation.RemoveBannerCreativeIDs(ids...)
	return cuo
}

// RemoveBannerCreatives removes "banner_creatives" edges to BannerCreative entities.
func (cuo *CreativeUpdateOne) RemoveBannerCreatives(b ...*BannerCreative) *CreativeUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cuo.RemoveBannerCreativeIDs(ids...)
}

// Where appends a list predicates to the CreativeUpdate builder.
func (cuo *CreativeUpdateOne) Where(ps ...predicate.Creative) *CreativeUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CreativeUpdateOne) Select(field string, fields ...string) *CreativeUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Creative entity.
func (cuo *CreativeUpdateOne) Save(ctx context.Context) (*Creative, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CreativeUpdateOne) SaveX(ctx context.Context) *Creative {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CreativeUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CreativeUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CreativeUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := creative.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

func (cuo *CreativeUpdateOne) sqlSave(ctx context.Context) (_node *Creative, err error) {
	_spec := sqlgraph.NewUpdateSpec(creative.Table, creative.Columns, sqlgraph.NewFieldSpec(creative.FieldID, field.TypeInt64))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Creative.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, creative.FieldID)
		for _, f := range fields {
			if !creative.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != creative.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(creative.FieldName, field.TypeString, value)
	}
	if cuo.mutation.NameCleared() {
		_spec.ClearField(creative.FieldName, field.TypeString)
	}
	if value, ok := cuo.mutation.ImageURL(); ok {
		_spec.SetField(creative.FieldImageURL, field.TypeString, value)
	}
	if cuo.mutation.ImageURLCleared() {
		_spec.ClearField(creative.FieldImageURL, field.TypeString)
	}
	if value, ok := cuo.mutation.Size(); ok {
		_spec.SetField(creative.FieldSize, field.TypeString, value)
	}
	if cuo.mutation.SizeCleared() {
		_spec.ClearField(creative.FieldSize, field.TypeString)
	}
	if value, ok := cuo.mutation.Enabled(); ok {
		_spec.SetField(creative.FieldEnabled, field.TypeBool, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(creative.FieldUpdatedAt, field.TypeTime, value)
	}
	if cuo.mutation.BannersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   creative.BannersTable,
			Columns: creative.BannersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(banner.FieldID, field.TypeInt64),
			},
		}
		createE := &BannerCreativeCreate{config: cuo.config, mutation: newBannerCreativeMutation(cuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedBannersIDs(); len(nodes) > 0 && !cuo.mutation.BannersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   creative.BannersTable,
			Columns: creative.BannersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(banner.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &BannerCreativeCreate{config: cuo.config, mutation: newBannerCreativeMutation(cuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.BannersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   creative.BannersTable,
			Columns: creative.BannersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(banner.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &BannerCreativeCreate{config: cuo.config, mutation: newBannerCreativeMutation(cuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.BannerCreativesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   creative.BannerCreativesTable,
			Columns: []string{creative.BannerCreativesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bannercreative.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedBannerCreativesIDs(); len(nodes) > 0 && !cuo.mutation.BannerCreativesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   creative.BannerCreativesTable,
			Columns: []string{creative.BannerCreativesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bannercreative.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.BannerCreativesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   creative.BannerCreativesTable,
			Columns: []string{creative.BannerCreativesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bannercreative.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Creative{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{creative.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
