package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("etc"),
		field.String("status"),
        field.Time("created_at").
            Default(time.Now),
		field.Time("update_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.String("title"),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return nil
}
