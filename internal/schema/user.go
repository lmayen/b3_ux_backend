package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),

		field.String("username").
			Unique(),

		field.Bool("is_admin"),

		field.String("email").
			Unique(),

		field.String("password"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("recommended_movies", WywwMovie.Type),
		edge.To("watched_movies", WywwMovie.Type),
		edge.To("installed_apps", StoreApp.Type),
		edge.To("sessions", Session.Type),
	}
}
