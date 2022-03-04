// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mark-ignacio/http2sqlite/ent/httprequest"
)

// HTTPRequestCreate is the builder for creating a HTTPRequest entity.
type HTTPRequestCreate struct {
	config
	mutation *HTTPRequestMutation
	hooks    []Hook
}

// SetReceived sets the "received" field.
func (hrc *HTTPRequestCreate) SetReceived(t time.Time) *HTTPRequestCreate {
	hrc.mutation.SetReceived(t)
	return hrc
}

// SetNillableReceived sets the "received" field if the given value is not nil.
func (hrc *HTTPRequestCreate) SetNillableReceived(t *time.Time) *HTTPRequestCreate {
	if t != nil {
		hrc.SetReceived(*t)
	}
	return hrc
}

// SetHost sets the "host" field.
func (hrc *HTTPRequestCreate) SetHost(s string) *HTTPRequestCreate {
	hrc.mutation.SetHost(s)
	return hrc
}

// SetPath sets the "path" field.
func (hrc *HTTPRequestCreate) SetPath(s string) *HTTPRequestCreate {
	hrc.mutation.SetPath(s)
	return hrc
}

// SetMethod sets the "method" field.
func (hrc *HTTPRequestCreate) SetMethod(s string) *HTTPRequestCreate {
	hrc.mutation.SetMethod(s)
	return hrc
}

// SetHeader sets the "header" field.
func (hrc *HTTPRequestCreate) SetHeader(h http.Header) *HTTPRequestCreate {
	hrc.mutation.SetHeader(h)
	return hrc
}

// SetBody sets the "body" field.
func (hrc *HTTPRequestCreate) SetBody(b []byte) *HTTPRequestCreate {
	hrc.mutation.SetBody(b)
	return hrc
}

// Mutation returns the HTTPRequestMutation object of the builder.
func (hrc *HTTPRequestCreate) Mutation() *HTTPRequestMutation {
	return hrc.mutation
}

// Save creates the HTTPRequest in the database.
func (hrc *HTTPRequestCreate) Save(ctx context.Context) (*HTTPRequest, error) {
	var (
		err  error
		node *HTTPRequest
	)
	hrc.defaults()
	if len(hrc.hooks) == 0 {
		if err = hrc.check(); err != nil {
			return nil, err
		}
		node, err = hrc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HTTPRequestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hrc.check(); err != nil {
				return nil, err
			}
			hrc.mutation = mutation
			if node, err = hrc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(hrc.hooks) - 1; i >= 0; i-- {
			if hrc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hrc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hrc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (hrc *HTTPRequestCreate) SaveX(ctx context.Context) *HTTPRequest {
	v, err := hrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hrc *HTTPRequestCreate) Exec(ctx context.Context) error {
	_, err := hrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hrc *HTTPRequestCreate) ExecX(ctx context.Context) {
	if err := hrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hrc *HTTPRequestCreate) defaults() {
	if _, ok := hrc.mutation.Received(); !ok {
		v := httprequest.DefaultReceived()
		hrc.mutation.SetReceived(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hrc *HTTPRequestCreate) check() error {
	if _, ok := hrc.mutation.Received(); !ok {
		return &ValidationError{Name: "received", err: errors.New(`ent: missing required field "HTTPRequest.received"`)}
	}
	if _, ok := hrc.mutation.Host(); !ok {
		return &ValidationError{Name: "host", err: errors.New(`ent: missing required field "HTTPRequest.host"`)}
	}
	if _, ok := hrc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "HTTPRequest.path"`)}
	}
	if _, ok := hrc.mutation.Method(); !ok {
		return &ValidationError{Name: "method", err: errors.New(`ent: missing required field "HTTPRequest.method"`)}
	}
	if _, ok := hrc.mutation.Header(); !ok {
		return &ValidationError{Name: "header", err: errors.New(`ent: missing required field "HTTPRequest.header"`)}
	}
	if _, ok := hrc.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New(`ent: missing required field "HTTPRequest.body"`)}
	}
	return nil
}

func (hrc *HTTPRequestCreate) sqlSave(ctx context.Context) (*HTTPRequest, error) {
	_node, _spec := hrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (hrc *HTTPRequestCreate) createSpec() (*HTTPRequest, *sqlgraph.CreateSpec) {
	var (
		_node = &HTTPRequest{config: hrc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: httprequest.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: httprequest.FieldID,
			},
		}
	)
	if value, ok := hrc.mutation.Received(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: httprequest.FieldReceived,
		})
		_node.Received = value
	}
	if value, ok := hrc.mutation.Host(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: httprequest.FieldHost,
		})
		_node.Host = value
	}
	if value, ok := hrc.mutation.Path(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: httprequest.FieldPath,
		})
		_node.Path = value
	}
	if value, ok := hrc.mutation.Method(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: httprequest.FieldMethod,
		})
		_node.Method = value
	}
	if value, ok := hrc.mutation.Header(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: httprequest.FieldHeader,
		})
		_node.Header = value
	}
	if value, ok := hrc.mutation.Body(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: httprequest.FieldBody,
		})
		_node.Body = value
	}
	return _node, _spec
}

// HTTPRequestCreateBulk is the builder for creating many HTTPRequest entities in bulk.
type HTTPRequestCreateBulk struct {
	config
	builders []*HTTPRequestCreate
}

// Save creates the HTTPRequest entities in the database.
func (hrcb *HTTPRequestCreateBulk) Save(ctx context.Context) ([]*HTTPRequest, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hrcb.builders))
	nodes := make([]*HTTPRequest, len(hrcb.builders))
	mutators := make([]Mutator, len(hrcb.builders))
	for i := range hrcb.builders {
		func(i int, root context.Context) {
			builder := hrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HTTPRequestMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, hrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hrcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, hrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hrcb *HTTPRequestCreateBulk) SaveX(ctx context.Context) []*HTTPRequest {
	v, err := hrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hrcb *HTTPRequestCreateBulk) Exec(ctx context.Context) error {
	_, err := hrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hrcb *HTTPRequestCreateBulk) ExecX(ctx context.Context) {
	if err := hrcb.Exec(ctx); err != nil {
		panic(err)
	}
}
