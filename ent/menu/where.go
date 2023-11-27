// Code generated by ent, DO NOT EDIT.

package menu

import (
	"ilsan/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldID, id))
}

// Etc applies equality check predicate on the "etc" field. It's identical to EtcEQ.
func Etc(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldEtc, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldStatus, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldUpdateAt, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldTitle, v))
}

// EtcEQ applies the EQ predicate on the "etc" field.
func EtcEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldEtc, v))
}

// EtcNEQ applies the NEQ predicate on the "etc" field.
func EtcNEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldEtc, v))
}

// EtcIn applies the In predicate on the "etc" field.
func EtcIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldEtc, vs...))
}

// EtcNotIn applies the NotIn predicate on the "etc" field.
func EtcNotIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldEtc, vs...))
}

// EtcGT applies the GT predicate on the "etc" field.
func EtcGT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldEtc, v))
}

// EtcGTE applies the GTE predicate on the "etc" field.
func EtcGTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldEtc, v))
}

// EtcLT applies the LT predicate on the "etc" field.
func EtcLT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldEtc, v))
}

// EtcLTE applies the LTE predicate on the "etc" field.
func EtcLTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldEtc, v))
}

// EtcContains applies the Contains predicate on the "etc" field.
func EtcContains(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContains(FieldEtc, v))
}

// EtcHasPrefix applies the HasPrefix predicate on the "etc" field.
func EtcHasPrefix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasPrefix(FieldEtc, v))
}

// EtcHasSuffix applies the HasSuffix predicate on the "etc" field.
func EtcHasSuffix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasSuffix(FieldEtc, v))
}

// EtcEqualFold applies the EqualFold predicate on the "etc" field.
func EtcEqualFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEqualFold(FieldEtc, v))
}

// EtcContainsFold applies the ContainsFold predicate on the "etc" field.
func EtcContainsFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContainsFold(FieldEtc, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContainsFold(FieldStatus, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldUpdateAt, v))
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldUpdateAt, v))
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldUpdateAt, vs...))
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldUpdateAt, vs...))
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldUpdateAt, v))
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldUpdateAt, v))
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldUpdateAt, v))
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldUpdateAt, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContainsFold(FieldTitle, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Menu) predicate.Menu {
	return predicate.Menu(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Menu) predicate.Menu {
	return predicate.Menu(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Menu) predicate.Menu {
	return predicate.Menu(sql.NotPredicates(p))
}
