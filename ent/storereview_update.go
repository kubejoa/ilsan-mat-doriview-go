// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"ilsan/ent/predicate"
	"ilsan/ent/storereview"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StoreReviewUpdate is the builder for updating StoreReview entities.
type StoreReviewUpdate struct {
	config
	hooks    []Hook
	mutation *StoreReviewMutation
}

// Where appends a list predicates to the StoreReviewUpdate builder.
func (sru *StoreReviewUpdate) Where(ps ...predicate.StoreReview) *StoreReviewUpdate {
	sru.mutation.Where(ps...)
	return sru
}

// SetEtc sets the "etc" field.
func (sru *StoreReviewUpdate) SetEtc(s string) *StoreReviewUpdate {
	sru.mutation.SetEtc(s)
	return sru
}

// SetNillableEtc sets the "etc" field if the given value is not nil.
func (sru *StoreReviewUpdate) SetNillableEtc(s *string) *StoreReviewUpdate {
	if s != nil {
		sru.SetEtc(*s)
	}
	return sru
}

// SetStatus sets the "status" field.
func (sru *StoreReviewUpdate) SetStatus(s string) *StoreReviewUpdate {
	sru.mutation.SetStatus(s)
	return sru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sru *StoreReviewUpdate) SetNillableStatus(s *string) *StoreReviewUpdate {
	if s != nil {
		sru.SetStatus(*s)
	}
	return sru
}

// SetCreatedAt sets the "created_at" field.
func (sru *StoreReviewUpdate) SetCreatedAt(t time.Time) *StoreReviewUpdate {
	sru.mutation.SetCreatedAt(t)
	return sru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sru *StoreReviewUpdate) SetNillableCreatedAt(t *time.Time) *StoreReviewUpdate {
	if t != nil {
		sru.SetCreatedAt(*t)
	}
	return sru
}

// SetUpdateAt sets the "update_at" field.
func (sru *StoreReviewUpdate) SetUpdateAt(t time.Time) *StoreReviewUpdate {
	sru.mutation.SetUpdateAt(t)
	return sru
}

// SetVisitDay sets the "visit_day" field.
func (sru *StoreReviewUpdate) SetVisitDay(t time.Time) *StoreReviewUpdate {
	sru.mutation.SetVisitDay(t)
	return sru
}

// SetNillableVisitDay sets the "visit_day" field if the given value is not nil.
func (sru *StoreReviewUpdate) SetNillableVisitDay(t *time.Time) *StoreReviewUpdate {
	if t != nil {
		sru.SetVisitDay(*t)
	}
	return sru
}

// SetScore sets the "score" field.
func (sru *StoreReviewUpdate) SetScore(f float64) *StoreReviewUpdate {
	sru.mutation.ResetScore()
	sru.mutation.SetScore(f)
	return sru
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (sru *StoreReviewUpdate) SetNillableScore(f *float64) *StoreReviewUpdate {
	if f != nil {
		sru.SetScore(*f)
	}
	return sru
}

// AddScore adds f to the "score" field.
func (sru *StoreReviewUpdate) AddScore(f float64) *StoreReviewUpdate {
	sru.mutation.AddScore(f)
	return sru
}

// SetComment sets the "comment" field.
func (sru *StoreReviewUpdate) SetComment(s string) *StoreReviewUpdate {
	sru.mutation.SetComment(s)
	return sru
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (sru *StoreReviewUpdate) SetNillableComment(s *string) *StoreReviewUpdate {
	if s != nil {
		sru.SetComment(*s)
	}
	return sru
}

// Mutation returns the StoreReviewMutation object of the builder.
func (sru *StoreReviewUpdate) Mutation() *StoreReviewMutation {
	return sru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sru *StoreReviewUpdate) Save(ctx context.Context) (int, error) {
	sru.defaults()
	return withHooks(ctx, sru.sqlSave, sru.mutation, sru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sru *StoreReviewUpdate) SaveX(ctx context.Context) int {
	affected, err := sru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sru *StoreReviewUpdate) Exec(ctx context.Context) error {
	_, err := sru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sru *StoreReviewUpdate) ExecX(ctx context.Context) {
	if err := sru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sru *StoreReviewUpdate) defaults() {
	if _, ok := sru.mutation.UpdateAt(); !ok {
		v := storereview.UpdateDefaultUpdateAt()
		sru.mutation.SetUpdateAt(v)
	}
}

func (sru *StoreReviewUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(storereview.Table, storereview.Columns, sqlgraph.NewFieldSpec(storereview.FieldID, field.TypeInt))
	if ps := sru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sru.mutation.Etc(); ok {
		_spec.SetField(storereview.FieldEtc, field.TypeString, value)
	}
	if value, ok := sru.mutation.Status(); ok {
		_spec.SetField(storereview.FieldStatus, field.TypeString, value)
	}
	if value, ok := sru.mutation.CreatedAt(); ok {
		_spec.SetField(storereview.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := sru.mutation.UpdateAt(); ok {
		_spec.SetField(storereview.FieldUpdateAt, field.TypeTime, value)
	}
	if value, ok := sru.mutation.VisitDay(); ok {
		_spec.SetField(storereview.FieldVisitDay, field.TypeTime, value)
	}
	if value, ok := sru.mutation.Score(); ok {
		_spec.SetField(storereview.FieldScore, field.TypeFloat64, value)
	}
	if value, ok := sru.mutation.AddedScore(); ok {
		_spec.AddField(storereview.FieldScore, field.TypeFloat64, value)
	}
	if value, ok := sru.mutation.Comment(); ok {
		_spec.SetField(storereview.FieldComment, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{storereview.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	sru.mutation.done = true
	return n, nil
}

// StoreReviewUpdateOne is the builder for updating a single StoreReview entity.
type StoreReviewUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StoreReviewMutation
}

// SetEtc sets the "etc" field.
func (sruo *StoreReviewUpdateOne) SetEtc(s string) *StoreReviewUpdateOne {
	sruo.mutation.SetEtc(s)
	return sruo
}

// SetNillableEtc sets the "etc" field if the given value is not nil.
func (sruo *StoreReviewUpdateOne) SetNillableEtc(s *string) *StoreReviewUpdateOne {
	if s != nil {
		sruo.SetEtc(*s)
	}
	return sruo
}

// SetStatus sets the "status" field.
func (sruo *StoreReviewUpdateOne) SetStatus(s string) *StoreReviewUpdateOne {
	sruo.mutation.SetStatus(s)
	return sruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sruo *StoreReviewUpdateOne) SetNillableStatus(s *string) *StoreReviewUpdateOne {
	if s != nil {
		sruo.SetStatus(*s)
	}
	return sruo
}

// SetCreatedAt sets the "created_at" field.
func (sruo *StoreReviewUpdateOne) SetCreatedAt(t time.Time) *StoreReviewUpdateOne {
	sruo.mutation.SetCreatedAt(t)
	return sruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sruo *StoreReviewUpdateOne) SetNillableCreatedAt(t *time.Time) *StoreReviewUpdateOne {
	if t != nil {
		sruo.SetCreatedAt(*t)
	}
	return sruo
}

// SetUpdateAt sets the "update_at" field.
func (sruo *StoreReviewUpdateOne) SetUpdateAt(t time.Time) *StoreReviewUpdateOne {
	sruo.mutation.SetUpdateAt(t)
	return sruo
}

// SetVisitDay sets the "visit_day" field.
func (sruo *StoreReviewUpdateOne) SetVisitDay(t time.Time) *StoreReviewUpdateOne {
	sruo.mutation.SetVisitDay(t)
	return sruo
}

// SetNillableVisitDay sets the "visit_day" field if the given value is not nil.
func (sruo *StoreReviewUpdateOne) SetNillableVisitDay(t *time.Time) *StoreReviewUpdateOne {
	if t != nil {
		sruo.SetVisitDay(*t)
	}
	return sruo
}

// SetScore sets the "score" field.
func (sruo *StoreReviewUpdateOne) SetScore(f float64) *StoreReviewUpdateOne {
	sruo.mutation.ResetScore()
	sruo.mutation.SetScore(f)
	return sruo
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (sruo *StoreReviewUpdateOne) SetNillableScore(f *float64) *StoreReviewUpdateOne {
	if f != nil {
		sruo.SetScore(*f)
	}
	return sruo
}

// AddScore adds f to the "score" field.
func (sruo *StoreReviewUpdateOne) AddScore(f float64) *StoreReviewUpdateOne {
	sruo.mutation.AddScore(f)
	return sruo
}

// SetComment sets the "comment" field.
func (sruo *StoreReviewUpdateOne) SetComment(s string) *StoreReviewUpdateOne {
	sruo.mutation.SetComment(s)
	return sruo
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (sruo *StoreReviewUpdateOne) SetNillableComment(s *string) *StoreReviewUpdateOne {
	if s != nil {
		sruo.SetComment(*s)
	}
	return sruo
}

// Mutation returns the StoreReviewMutation object of the builder.
func (sruo *StoreReviewUpdateOne) Mutation() *StoreReviewMutation {
	return sruo.mutation
}

// Where appends a list predicates to the StoreReviewUpdate builder.
func (sruo *StoreReviewUpdateOne) Where(ps ...predicate.StoreReview) *StoreReviewUpdateOne {
	sruo.mutation.Where(ps...)
	return sruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sruo *StoreReviewUpdateOne) Select(field string, fields ...string) *StoreReviewUpdateOne {
	sruo.fields = append([]string{field}, fields...)
	return sruo
}

// Save executes the query and returns the updated StoreReview entity.
func (sruo *StoreReviewUpdateOne) Save(ctx context.Context) (*StoreReview, error) {
	sruo.defaults()
	return withHooks(ctx, sruo.sqlSave, sruo.mutation, sruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sruo *StoreReviewUpdateOne) SaveX(ctx context.Context) *StoreReview {
	node, err := sruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sruo *StoreReviewUpdateOne) Exec(ctx context.Context) error {
	_, err := sruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sruo *StoreReviewUpdateOne) ExecX(ctx context.Context) {
	if err := sruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sruo *StoreReviewUpdateOne) defaults() {
	if _, ok := sruo.mutation.UpdateAt(); !ok {
		v := storereview.UpdateDefaultUpdateAt()
		sruo.mutation.SetUpdateAt(v)
	}
}

func (sruo *StoreReviewUpdateOne) sqlSave(ctx context.Context) (_node *StoreReview, err error) {
	_spec := sqlgraph.NewUpdateSpec(storereview.Table, storereview.Columns, sqlgraph.NewFieldSpec(storereview.FieldID, field.TypeInt))
	id, ok := sruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "StoreReview.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, storereview.FieldID)
		for _, f := range fields {
			if !storereview.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != storereview.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sruo.mutation.Etc(); ok {
		_spec.SetField(storereview.FieldEtc, field.TypeString, value)
	}
	if value, ok := sruo.mutation.Status(); ok {
		_spec.SetField(storereview.FieldStatus, field.TypeString, value)
	}
	if value, ok := sruo.mutation.CreatedAt(); ok {
		_spec.SetField(storereview.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := sruo.mutation.UpdateAt(); ok {
		_spec.SetField(storereview.FieldUpdateAt, field.TypeTime, value)
	}
	if value, ok := sruo.mutation.VisitDay(); ok {
		_spec.SetField(storereview.FieldVisitDay, field.TypeTime, value)
	}
	if value, ok := sruo.mutation.Score(); ok {
		_spec.SetField(storereview.FieldScore, field.TypeFloat64, value)
	}
	if value, ok := sruo.mutation.AddedScore(); ok {
		_spec.AddField(storereview.FieldScore, field.TypeFloat64, value)
	}
	if value, ok := sruo.mutation.Comment(); ok {
		_spec.SetField(storereview.FieldComment, field.TypeString, value)
	}
	_node = &StoreReview{config: sruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{storereview.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sruo.mutation.done = true
	return _node, nil
}
