// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"ilsan/ent/store"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StoreCreate is the builder for creating a Store entity.
type StoreCreate struct {
	config
	mutation *StoreMutation
	hooks    []Hook
}

// SetEtc sets the "etc" field.
func (sc *StoreCreate) SetEtc(s string) *StoreCreate {
	sc.mutation.SetEtc(s)
	return sc
}

// SetStatus sets the "status" field.
func (sc *StoreCreate) SetStatus(s string) *StoreCreate {
	sc.mutation.SetStatus(s)
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *StoreCreate) SetCreatedAt(t time.Time) *StoreCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *StoreCreate) SetNillableCreatedAt(t *time.Time) *StoreCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdateAt sets the "update_at" field.
func (sc *StoreCreate) SetUpdateAt(t time.Time) *StoreCreate {
	sc.mutation.SetUpdateAt(t)
	return sc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (sc *StoreCreate) SetNillableUpdateAt(t *time.Time) *StoreCreate {
	if t != nil {
		sc.SetUpdateAt(*t)
	}
	return sc
}

// SetTitle sets the "title" field.
func (sc *StoreCreate) SetTitle(s string) *StoreCreate {
	sc.mutation.SetTitle(s)
	return sc
}

// Mutation returns the StoreMutation object of the builder.
func (sc *StoreCreate) Mutation() *StoreMutation {
	return sc.mutation
}

// Save creates the Store in the database.
func (sc *StoreCreate) Save(ctx context.Context) (*Store, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StoreCreate) SaveX(ctx context.Context) *Store {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StoreCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StoreCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StoreCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := store.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdateAt(); !ok {
		v := store.DefaultUpdateAt()
		sc.mutation.SetUpdateAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StoreCreate) check() error {
	if _, ok := sc.mutation.Etc(); !ok {
		return &ValidationError{Name: "etc", err: errors.New(`ent: missing required field "Store.etc"`)}
	}
	if _, ok := sc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Store.status"`)}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Store.created_at"`)}
	}
	if _, ok := sc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "Store.update_at"`)}
	}
	if _, ok := sc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Store.title"`)}
	}
	return nil
}

func (sc *StoreCreate) sqlSave(ctx context.Context) (*Store, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *StoreCreate) createSpec() (*Store, *sqlgraph.CreateSpec) {
	var (
		_node = &Store{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(store.Table, sqlgraph.NewFieldSpec(store.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.Etc(); ok {
		_spec.SetField(store.FieldEtc, field.TypeString, value)
		_node.Etc = value
	}
	if value, ok := sc.mutation.Status(); ok {
		_spec.SetField(store.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(store.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdateAt(); ok {
		_spec.SetField(store.FieldUpdateAt, field.TypeTime, value)
		_node.UpdateAt = value
	}
	if value, ok := sc.mutation.Title(); ok {
		_spec.SetField(store.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	return _node, _spec
}

// StoreCreateBulk is the builder for creating many Store entities in bulk.
type StoreCreateBulk struct {
	config
	err      error
	builders []*StoreCreate
}

// Save creates the Store entities in the database.
func (scb *StoreCreateBulk) Save(ctx context.Context) ([]*Store, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Store, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StoreMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StoreCreateBulk) SaveX(ctx context.Context) []*Store {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StoreCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StoreCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
