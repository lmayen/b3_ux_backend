package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Color holds the schema definition for the Color entity.
type Color struct {
	ent.Schema
}

// Fields of the Color.
func (Color) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),

		field.Int("r").
			NonNegative().
			Max(255),

		field.Int("g").
			NonNegative().
			Max(255),

		field.Int("b").
			NonNegative().
			Max(255),

		field.Float32("ratio").
			Min(0).
			Max(1),
	}
}

// Edges of the Color.
func (Color) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("image", Image.Type).Ref("colors").Unique(),
	}
}
