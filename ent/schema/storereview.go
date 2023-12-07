package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// StoreReview holds the schema definition for the StoreReview entity.
type StoreReview struct {
	ent.Schema
}

// Fields of the StoreReview.
func (StoreReview) Fields() []ent.Field {
	return []ent.Field{
		field.String("etc"),
		field.String("status"),
        field.Time("created_at").
            Default(time.Now),
		field.Time("update_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("visit_day"),
		field.Float("score"),
		field.String("comment"),
	}
}

// Edges of the StoreReview.
func (StoreReview) Edges() []ent.Edge {
	return nil
}
