package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// StoreApp holds the schema definition for the StoreApp entity.
type StoreApp struct {
	ent.Schema
}

// Fields of the StoreApp.
func (StoreApp) Fields() []ent.Field {
	return []ent.Field{

		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),

		field.String("name").NotEmpty(),

		field.Enum("category").Values(
			"ARTIFICIAL_INTELLIGENCE",
			"ART_AND_DESIGN",
			"AUTO_AND_VEHICLES",
			"BEAUTY",
			"BOOKS_AND_REFERENCE",
			"BUSINESS",
			"COMICS",
			"COMMUNICATION",
			"DATING",
			"EDUCATION",
			"ENTERTAINMENT",
			"EVENTS",
			"FAMILY",
			"FINANCE",
			"FOOD_AND_DRINK",
			"GAME",
			"HEALTH_AND_FITNESS",
			"HOUSE_AND_HOME",
			"LIBRARIES_AND_DEMO",
			"LIFESTYLE",
			"MAPS_AND_NAVIGATION",
			"MEDICAL",
			"NEWS_AND_MAGAZINES",
			"PARENTING",
			"PERSONALIZATION",
			"PHOTOGRAPHY",
			"PRODUCTIVITY",
			"SHOPPING",
			"SOCIAL",
			"SPORTS",
			"TOOLS",
			"TRAVEL_AND_LOCAL",
			"VIDEO_PLAYERS",
			"VIRTUAL_REALITY",
			"WEATHER",
			"WEB3_AND_CRYPTO",
		),

		field.Float32("rating").Nillable(),

		field.Int32("reviews").NonNegative(),

		field.String("size").NotEmpty(),

		field.String("installs").NotEmpty(),

		field.Enum("type").Values("Free", "Paid"),

		field.Float32("price").Min(0).Default(0),

		field.String("content_rating").NotEmpty(),

		field.String("last_updated").NotEmpty(),

		field.String("current_ver").NotEmpty(),

		field.Bool("in_app_purchases"),

		field.Bool("ad_supported"),
	}
}

// Edges of the StoreApp.
func (StoreApp) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("genres", StoreGenre.Type),
		edge.To("images", Image.Type),

		edge.From("user_install_list", User.Type).
			Ref("installed_apps"),
	}
}
