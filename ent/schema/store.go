package schema

import (
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Store holds the schema definition for the Store entity.
type Store struct {
	ent.Schema
}

// Fields of the Store.
func (Store) Fields() []ent.Field {
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

// Edges of the Store.
func (Store) Edges() []ent.Edge {
	return nil
}
