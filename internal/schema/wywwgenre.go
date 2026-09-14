package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WywwGenre holds the schema definition for the WywwGenre entity.
type WywwGenre struct {
	ent.Schema
}

// Fields of the WywwGenre.
func (WywwGenre) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),

		field.String("name").NotEmpty(),
	}
}

// Edges of the WywwGenre.
func (WywwGenre) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("app", WywwMovie.Type).
			Ref("genres"),
	}
}
