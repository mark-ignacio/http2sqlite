// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mark-ignacio/http2sqlite/ent/httprequest"
	"github.com/mark-ignacio/http2sqlite/ent/predicate"
)

// HTTPRequestDelete is the builder for deleting a HTTPRequest entity.
type HTTPRequestDelete struct {
	config
	hooks    []Hook
	mutation *HTTPRequestMutation
}

// Where appends a list predicates to the HTTPRequestDelete builder.
func (hrd *HTTPRequestDelete) Where(ps ...predicate.HTTPRequest) *HTTPRequestDelete {
	hrd.mutation.Where(ps...)
	return hrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hrd *HTTPRequestDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(hrd.hooks) == 0 {
		affected, err = hrd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HTTPRequestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hrd.mutation = mutation
			affected, err = hrd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hrd.hooks) - 1; i >= 0; i-- {
			if hrd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hrd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hrd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (hrd *HTTPRequestDelete) ExecX(ctx context.Context) int {
	n, err := hrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hrd *HTTPRequestDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: httprequest.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: httprequest.FieldID,
			},
		},
	}
	if ps := hrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, hrd.driver, _spec)
}

// HTTPRequestDeleteOne is the builder for deleting a single HTTPRequest entity.
type HTTPRequestDeleteOne struct {
	hrd *HTTPRequestDelete
}

// Exec executes the deletion query.
func (hrdo *HTTPRequestDeleteOne) Exec(ctx context.Context) error {
	n, err := hrdo.hrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{httprequest.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hrdo *HTTPRequestDeleteOne) ExecX(ctx context.Context) {
	hrdo.hrd.ExecX(ctx)
}
