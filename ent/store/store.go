// Code generated by ent, DO NOT EDIT.

package store

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the store type in the database.
	Label = "store"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEtc holds the string denoting the etc field in the database.
	FieldEtc = "etc"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// Table holds the table name of the store in the database.
	Table = "stores"
)

// Columns holds all SQL columns for store fields.
var Columns = []string{
	FieldID,
	FieldEtc,
	FieldStatus,
	FieldCreatedAt,
	FieldUpdateAt,
	FieldTitle,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() time.Time
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() time.Time
)

// OrderOption defines the ordering options for the Store queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEtc orders the results by the etc field.
func ByEtc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEtc, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdateAt orders the results by the update_at field.
func ByUpdateAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateAt, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}
