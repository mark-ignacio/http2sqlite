// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mark-ignacio/http2sqlite/ent/httprequest"
	"github.com/mark-ignacio/http2sqlite/ent/predicate"
)

// HTTPRequestUpdate is the builder for updating HTTPRequest entities.
type HTTPRequestUpdate struct {
	config
	hooks    []Hook
	mutation *HTTPRequestMutation
}

// Where appends a list predicates to the HTTPRequestUpdate builder.
func (hru *HTTPRequestUpdate) Where(ps ...predicate.HTTPRequest) *HTTPRequestUpdate {
	hru.mutation.Where(ps...)
	return hru
}

// SetReceived sets the "received" field.
func (hru *HTTPRequestUpdate) SetReceived(t time.Time) *HTTPRequestUpdate {
	hru.mutation.SetReceived(t)
	return hru
}

// SetNillableReceived sets the "received" field if the given value is not nil.
func (hru *HTTPRequestUpdate) SetNillableReceived(t *time.Time) *HTTPRequestUpdate {
	if t != nil {
		hru.SetReceived(*t)
	}
	return hru
}

// SetHost sets the "host" field.
func (hru *HTTPRequestUpdate) SetHost(s string) *HTTPRequestUpdate {
	hru.mutation.SetHost(s)
	return hru
}

// SetPath sets the "path" field.
func (hru *HTTPRequestUpdate) SetPath(s string) *HTTPRequestUpdate {
	hru.mutation.SetPath(s)
	return hru
}

// SetMethod sets the "method" field.
func (hru *HTTPRequestUpdate) SetMethod(s string) *HTTPRequestUpdate {
	hru.mutation.SetMethod(s)
	return hru
}

// SetHeader sets the "header" field.
func (hru *HTTPRequestUpdate) SetHeader(h http.Header) *HTTPRequestUpdate {
	hru.mutation.SetHeader(h)
	return hru
}

// SetBody sets the "body" field.
func (hru *HTTPRequestUpdate) SetBody(b []byte) *HTTPRequestUpdate {
	hru.mutation.SetBody(b)
	return hru
}

// Mutation returns the HTTPRequestMutation object of the builder.
func (hru *HTTPRequestUpdate) Mutation() *HTTPRequestMutation {
	return hru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hru *HTTPRequestUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(hru.hooks) == 0 {
		affected, err = hru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HTTPRequestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hru.mutation = mutation
			affected, err = hru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hru.hooks) - 1; i >= 0; i-- {
			if hru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (hru *HTTPRequestUpdate) SaveX(ctx context.Context) int {
	affected, err := hru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hru *HTTPRequestUpdate) Exec(ctx context.Context) error {
	_, err := hru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hru *HTTPRequestUpdate) ExecX(ctx context.Context) {
	if err := hru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (hru *HTTPRequestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   httprequest.Table,
			Columns: httprequest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: httprequest.FieldID,
			},
		},
	}
	if ps := hru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hru.mutation.Received(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: httprequest.FieldReceived,
		})
	}
	if value, ok := hru.mutation.Host(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: httprequest.FieldHost,
		})
	}
	if value, ok := hru.mutation.Path(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: httprequest.FieldPath,
		})
	}
	if value, ok := hru.mutation.Method(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: httprequest.FieldMethod,
		})
	}
	if value, ok := hru.mutation.Header(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: httprequest.FieldHeader,
		})
	}
	if value, ok := hru.mutation.Body(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: httprequest.FieldBody,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{httprequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// HTTPRequestUpdateOne is the builder for updating a single HTTPRequest entity.
type HTTPRequestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HTTPRequestMutation
}

// SetReceived sets the "received" field.
func (hruo *HTTPRequestUpdateOne) SetReceived(t time.Time) *HTTPRequestUpdateOne {
	hruo.mutation.SetReceived(t)
	return hruo
}

// SetNillableReceived sets the "received" field if the given value is not nil.
func (hruo *HTTPRequestUpdateOne) SetNillableReceived(t *time.Time) *HTTPRequestUpdateOne {
	if t != nil {
		hruo.SetReceived(*t)
	}
	return hruo
}

// SetHost sets the "host" field.
func (hruo *HTTPRequestUpdateOne) SetHost(s string) *HTTPRequestUpdateOne {
	hruo.mutation.SetHost(s)
	return hruo
}

// SetPath sets the "path" field.
func (hruo *HTTPRequestUpdateOne) SetPath(s string) *HTTPRequestUpdateOne {
	hruo.mutation.SetPath(s)
	return hruo
}

// SetMethod sets the "method" field.
func (hruo *HTTPRequestUpdateOne) SetMethod(s string) *HTTPRequestUpdateOne {
	hruo.mutation.SetMethod(s)
	return hruo
}

// SetHeader sets the "header" field.
func (hruo *HTTPRequestUpdateOne) SetHeader(h http.Header) *HTTPRequestUpdateOne {
	hruo.mutation.SetHeader(h)
	return hruo
}

// SetBody sets the "body" field.
func (hruo *HTTPRequestUpdateOne) SetBody(b []byte) *HTTPRequestUpdateOne {
	hruo.mutation.SetBody(b)
	return hruo
}

// Mutation returns the HTTPRequestMutation object of the builder.
func (hruo *HTTPRequestUpdateOne) Mutation() *HTTPRequestMutation {
	return hruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (hruo *HTTPRequestUpdateOne) Select(field string, fields ...string) *HTTPRequestUpdateOne {
	hruo.fields = append([]string{field}, fields...)
	return hruo
}

// Save executes the query and returns the updated HTTPRequest entity.
func (hruo *HTTPRequestUpdateOne) Save(ctx context.Context) (*HTTPRequest, error) {
	var (
		err  error
		node *HTTPRequest
	)
	if len(hruo.hooks) == 0 {
		node, err = hruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HTTPRequestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hruo.mutation = mutation
			node, err = hruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(hruo.hooks) - 1; i >= 0; i-- {
			if hruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (hruo *HTTPRequestUpdateOne) SaveX(ctx context.Context) *HTTPRequest {
	node, err := hruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (hruo *HTTPRequestUpdateOne) Exec(ctx context.Context) error {
	_, err := hruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hruo *HTTPRequestUpdateOne) ExecX(ctx context.Context) {
	if err := hruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (hruo *HTTPRequestUpdateOne) sqlSave(ctx context.Context) (_node *HTTPRequest, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   httprequest.Table,
			Columns: httprequest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: httprequest.FieldID,
			},
		},
	}
	id, ok := hruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "HTTPRequest.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := hruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, httprequest.FieldID)
		for _, f := range fields {
			if !httprequest.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != httprequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := hruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hruo.mutation.Received(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: httprequest.FieldReceived,
		})
	}
	if value, ok := hruo.mutation.Host(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: httprequest.FieldHost,
		})
	}
	if value, ok := hruo.mutation.Path(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: httprequest.FieldPath,
		})
	}
	if value, ok := hruo.mutation.Method(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: httprequest.FieldMethod,
		})
	}
	if value, ok := hruo.mutation.Header(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: httprequest.FieldHeader,
		})
	}
	if value, ok := hruo.mutation.Body(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: httprequest.FieldBody,
		})
	}
	_node = &HTTPRequest{config: hruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, hruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{httprequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
