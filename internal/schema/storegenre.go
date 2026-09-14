package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// StoreGenre holds the schema definition for the StoreGenre entity.
type StoreGenre struct {
	ent.Schema
}

// Fields of the StoreGenre.
func (StoreGenre) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),

		field.String("name").NotEmpty(),
	}
}

// Edges of the StoreGenre.
func (StoreGenre) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("app", StoreApp.Type).
			Ref("genres"),
	}
}
