// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/lead"
	"affluo/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LeadDelete is the builder for deleting a Lead entity.
type LeadDelete struct {
	config
	hooks    []Hook
	mutation *LeadMutation
}

// Where appends a list predicates to the LeadDelete builder.
func (ld *LeadDelete) Where(ps ...predicate.Lead) *LeadDelete {
	ld.mutation.Where(ps...)
	return ld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ld *LeadDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ld.sqlExec, ld.mutation, ld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ld *LeadDelete) ExecX(ctx context.Context) int {
	n, err := ld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ld *LeadDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(lead.Table, sqlgraph.NewFieldSpec(lead.FieldID, field.TypeInt64))
	if ps := ld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ld.mutation.done = true
	return affected, err
}

// LeadDeleteOne is the builder for deleting a single Lead entity.
type LeadDeleteOne struct {
	ld *LeadDelete
}

// Where appends a list predicates to the LeadDelete builder.
func (ldo *LeadDeleteOne) Where(ps ...predicate.Lead) *LeadDeleteOne {
	ldo.ld.mutation.Where(ps...)
	return ldo
}

// Exec executes the deletion query.
func (ldo *LeadDeleteOne) Exec(ctx context.Context) error {
	n, err := ldo.ld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{lead.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ldo *LeadDeleteOne) ExecX(ctx context.Context) {
	if err := ldo.Exec(ctx); err != nil {
		panic(err)
	}
}
