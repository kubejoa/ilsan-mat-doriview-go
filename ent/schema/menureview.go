package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MenuReview holds the schema definition for the MenuReview entity.
type MenuReview struct {
	ent.Schema
}

// Fields of the MenuReview.
func (MenuReview) Fields() []ent.Field {
	return []ent.Field{
		field.String("etc"),
		field.String("status"),
        field.Time("created_at").
            Default(time.Now),
		field.Time("update_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Int("score"),
		field.String("comment"),
		field.Int("cooltime"),
		field.Int("price"),
	}
}

// Edges of the MenuReview.
func (MenuReview) Edges() []ent.Edge {
	return nil
}
