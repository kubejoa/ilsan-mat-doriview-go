// Code generated by ent, DO NOT EDIT.

package menureview

import (
	"ilsan/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldLTE(FieldID, id))
}

// Etc applies equality check predicate on the "etc" field. It's identical to EtcEQ.
func Etc(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldEQ(FieldEtc, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldEQ(FieldStatus, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v time.Time) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldEQ(FieldUpdateAt, v))
}

// Score applies equality check predicate on the "score" field. It's identical to ScoreEQ.
func Score(v int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldEQ(FieldScore, v))
}

// Comment applies equality check predicate on the "comment" field. It's identical to CommentEQ.
func Comment(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldEQ(FieldComment, v))
}

// Cooltime applies equality check predicate on the "cooltime" field. It's identical to CooltimeEQ.
func Cooltime(v int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldEQ(FieldCooltime, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldEQ(FieldPrice, v))
}

// EtcEQ applies the EQ predicate on the "etc" field.
func EtcEQ(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldEQ(FieldEtc, v))
}

// EtcNEQ applies the NEQ predicate on the "etc" field.
func EtcNEQ(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldNEQ(FieldEtc, v))
}

// EtcIn applies the In predicate on the "etc" field.
func EtcIn(vs ...string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldIn(FieldEtc, vs...))
}

// EtcNotIn applies the NotIn predicate on the "etc" field.
func EtcNotIn(vs ...string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldNotIn(FieldEtc, vs...))
}

// EtcGT applies the GT predicate on the "etc" field.
func EtcGT(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldGT(FieldEtc, v))
}

// EtcGTE applies the GTE predicate on the "etc" field.
func EtcGTE(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldGTE(FieldEtc, v))
}

// EtcLT applies the LT predicate on the "etc" field.
func EtcLT(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldLT(FieldEtc, v))
}

// EtcLTE applies the LTE predicate on the "etc" field.
func EtcLTE(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldLTE(FieldEtc, v))
}

// EtcContains applies the Contains predicate on the "etc" field.
func EtcContains(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldContains(FieldEtc, v))
}

// EtcHasPrefix applies the HasPrefix predicate on the "etc" field.
func EtcHasPrefix(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldHasPrefix(FieldEtc, v))
}

// EtcHasSuffix applies the HasSuffix predicate on the "etc" field.
func EtcHasSuffix(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldHasSuffix(FieldEtc, v))
}

// EtcEqualFold applies the EqualFold predicate on the "etc" field.
func EtcEqualFold(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldEqualFold(FieldEtc, v))
}

// EtcContainsFold applies the ContainsFold predicate on the "etc" field.
func EtcContainsFold(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldContainsFold(FieldEtc, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldContainsFold(FieldStatus, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v time.Time) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldEQ(FieldUpdateAt, v))
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v time.Time) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldNEQ(FieldUpdateAt, v))
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...time.Time) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldIn(FieldUpdateAt, vs...))
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...time.Time) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldNotIn(FieldUpdateAt, vs...))
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v time.Time) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldGT(FieldUpdateAt, v))
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v time.Time) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldGTE(FieldUpdateAt, v))
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v time.Time) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldLT(FieldUpdateAt, v))
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v time.Time) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldLTE(FieldUpdateAt, v))
}

// ScoreEQ applies the EQ predicate on the "score" field.
func ScoreEQ(v int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldEQ(FieldScore, v))
}

// ScoreNEQ applies the NEQ predicate on the "score" field.
func ScoreNEQ(v int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldNEQ(FieldScore, v))
}

// ScoreIn applies the In predicate on the "score" field.
func ScoreIn(vs ...int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldIn(FieldScore, vs...))
}

// ScoreNotIn applies the NotIn predicate on the "score" field.
func ScoreNotIn(vs ...int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldNotIn(FieldScore, vs...))
}

// ScoreGT applies the GT predicate on the "score" field.
func ScoreGT(v int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldGT(FieldScore, v))
}

// ScoreGTE applies the GTE predicate on the "score" field.
func ScoreGTE(v int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldGTE(FieldScore, v))
}

// ScoreLT applies the LT predicate on the "score" field.
func ScoreLT(v int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldLT(FieldScore, v))
}

// ScoreLTE applies the LTE predicate on the "score" field.
func ScoreLTE(v int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldLTE(FieldScore, v))
}

// CommentEQ applies the EQ predicate on the "comment" field.
func CommentEQ(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldEQ(FieldComment, v))
}

// CommentNEQ applies the NEQ predicate on the "comment" field.
func CommentNEQ(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldNEQ(FieldComment, v))
}

// CommentIn applies the In predicate on the "comment" field.
func CommentIn(vs ...string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldIn(FieldComment, vs...))
}

// CommentNotIn applies the NotIn predicate on the "comment" field.
func CommentNotIn(vs ...string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldNotIn(FieldComment, vs...))
}

// CommentGT applies the GT predicate on the "comment" field.
func CommentGT(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldGT(FieldComment, v))
}

// CommentGTE applies the GTE predicate on the "comment" field.
func CommentGTE(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldGTE(FieldComment, v))
}

// CommentLT applies the LT predicate on the "comment" field.
func CommentLT(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldLT(FieldComment, v))
}

// CommentLTE applies the LTE predicate on the "comment" field.
func CommentLTE(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldLTE(FieldComment, v))
}

// CommentContains applies the Contains predicate on the "comment" field.
func CommentContains(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldContains(FieldComment, v))
}

// CommentHasPrefix applies the HasPrefix predicate on the "comment" field.
func CommentHasPrefix(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldHasPrefix(FieldComment, v))
}

// CommentHasSuffix applies the HasSuffix predicate on the "comment" field.
func CommentHasSuffix(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldHasSuffix(FieldComment, v))
}

// CommentEqualFold applies the EqualFold predicate on the "comment" field.
func CommentEqualFold(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldEqualFold(FieldComment, v))
}

// CommentContainsFold applies the ContainsFold predicate on the "comment" field.
func CommentContainsFold(v string) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldContainsFold(FieldComment, v))
}

// CooltimeEQ applies the EQ predicate on the "cooltime" field.
func CooltimeEQ(v int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldEQ(FieldCooltime, v))
}

// CooltimeNEQ applies the NEQ predicate on the "cooltime" field.
func CooltimeNEQ(v int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldNEQ(FieldCooltime, v))
}

// CooltimeIn applies the In predicate on the "cooltime" field.
func CooltimeIn(vs ...int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldIn(FieldCooltime, vs...))
}

// CooltimeNotIn applies the NotIn predicate on the "cooltime" field.
func CooltimeNotIn(vs ...int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldNotIn(FieldCooltime, vs...))
}

// CooltimeGT applies the GT predicate on the "cooltime" field.
func CooltimeGT(v int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldGT(FieldCooltime, v))
}

// CooltimeGTE applies the GTE predicate on the "cooltime" field.
func CooltimeGTE(v int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldGTE(FieldCooltime, v))
}

// CooltimeLT applies the LT predicate on the "cooltime" field.
func CooltimeLT(v int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldLT(FieldCooltime, v))
}

// CooltimeLTE applies the LTE predicate on the "cooltime" field.
func CooltimeLTE(v int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldLTE(FieldCooltime, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v int) predicate.MenuReview {
	return predicate.MenuReview(sql.FieldLTE(FieldPrice, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MenuReview) predicate.MenuReview {
	return predicate.MenuReview(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MenuReview) predicate.MenuReview {
	return predicate.MenuReview(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MenuReview) predicate.MenuReview {
	return predicate.MenuReview(sql.NotPredicates(p))
}
