package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Reservation holds the schema definition for the Reservation entity.
type Reservation struct {
	ent.Schema
}

// Fields of the Reservation.
func (Reservation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("starts_at"),
		field.Time("ends_at"),
		field.UUID("room_id", uuid.UUID{}),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
		field.Time("deleted_at").Nillable().Optional(),
	}
}

// Edges of the Reservation.
func (Reservation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", Room.Type).Ref("reservations").Unique().Required().Field("room_id"),
	}
}

// Indexes of the Reservation.
func (Reservation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("starts_at", "ends_at"),
	}
}
