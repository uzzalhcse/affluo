// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/bannerstats"
	"affluo/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BannerStatsDelete is the builder for deleting a BannerStats entity.
type BannerStatsDelete struct {
	config
	hooks    []Hook
	mutation *BannerStatsMutation
}

// Where appends a list predicates to the BannerStatsDelete builder.
func (bsd *BannerStatsDelete) Where(ps ...predicate.BannerStats) *BannerStatsDelete {
	bsd.mutation.Where(ps...)
	return bsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bsd *BannerStatsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bsd.sqlExec, bsd.mutation, bsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bsd *BannerStatsDelete) ExecX(ctx context.Context) int {
	n, err := bsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bsd *BannerStatsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(bannerstats.Table, sqlgraph.NewFieldSpec(bannerstats.FieldID, field.TypeInt64))
	if ps := bsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bsd.mutation.done = true
	return affected, err
}

// BannerStatsDeleteOne is the builder for deleting a single BannerStats entity.
type BannerStatsDeleteOne struct {
	bsd *BannerStatsDelete
}

// Where appends a list predicates to the BannerStatsDelete builder.
func (bsdo *BannerStatsDeleteOne) Where(ps ...predicate.BannerStats) *BannerStatsDeleteOne {
	bsdo.bsd.mutation.Where(ps...)
	return bsdo
}

// Exec executes the deletion query.
func (bsdo *BannerStatsDeleteOne) Exec(ctx context.Context) error {
	n, err := bsdo.bsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{bannerstats.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bsdo *BannerStatsDeleteOne) ExecX(ctx context.Context) {
	if err := bsdo.Exec(ctx); err != nil {
		panic(err)
	}
}