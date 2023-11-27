package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Menu holds the schema definition for the Menu entity.
type Menu struct {
	ent.Schema
}

// Fields of the Menu.
func (Menu) Fields() []ent.Field {
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

// Edges of the Menu.
func (Menu) Edges() []ent.Edge {
	return nil
}
