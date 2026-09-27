package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WywwMovie holds the schema definition for the WywwMovie entity.
type WywwMovie struct {
	ent.Schema
}

// Fields of the WywwMovie.
func (WywwMovie) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),

		field.String("movie_title").NotEmpty(),

		field.String("original_title").NotEmpty(),

		field.Enum("content_rating").Values(
			"Approved",
			"G",
			"NC-17",
			"Not Rated",
			"PG",
			"PG-13",
			"Passed",
			"R",
		),

		field.String("description").NotEmpty(),

		field.Int32("released_year").NonNegative(),

		field.Int32("runtime").NonNegative(),

		field.Float32("rating"),
	}
}

// Edges of the WywwMovie.
func (WywwMovie) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("genres", WywwGenre.Type),
		edge.To("images", Image.Type),

		edge.From("user_recommendations", User.Type).
			Ref("recommended_movies"),

		edge.From("user_watchlist", User.Type).
			Ref("watched_movies"),
	}
}
