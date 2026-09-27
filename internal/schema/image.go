package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),

		field.String("filename").
			NotEmpty().
			Unique(),

		field.String("path").
			NotEmpty().
			Unique(),

		field.Int("width").
			NonNegative(),

		field.Int("height").
			NonNegative(),

		field.Enum("kind").
			Values("poster", "banner", "icon"),
	}
}

// Edges of the Image.
func (Image) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("colors", Color.Type),
		edge.From("movie", WywwMovie.Type).Ref("images").Unique(),
		edge.From("store", StoreApp.Type).Ref("images").Unique(),
	}
}
