package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),

		field.String("token").
			Unique().
			Immutable(),

		field.Time("created_at").
			Default(time.Now),

		field.Time("expires_at").
			Default(inOneMonth),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("artist", User.Type).
			Ref("sessions").
			Unique().
			Required().
			Immutable(),
	}
}

func inOneMonth() time.Time {
	now := time.Now()
	return now.AddDate(0, 0, 30)
}
