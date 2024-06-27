// Code generated by ent, DO NOT EDIT.

package ent

import (
	"bookit/internal/repo/ent/reservation"
	"bookit/internal/repo/ent/room"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Reservation is the model entity for the Reservation schema.
type Reservation struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// StartsAt holds the value of the "starts_at" field.
	StartsAt time.Time `json:"starts_at,omitempty"`
	// EndsAt holds the value of the "ends_at" field.
	EndsAt time.Time `json:"ends_at,omitempty"`
	// RoomID holds the value of the "room_id" field.
	RoomID uuid.UUID `json:"room_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ReservationQuery when eager-loading is set.
	Edges        ReservationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ReservationEdges holds the relations/edges for other nodes in the graph.
type ReservationEdges struct {
	// Room holds the value of the room edge.
	Room *Room `json:"room,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RoomOrErr returns the Room value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReservationEdges) RoomOrErr() (*Room, error) {
	if e.Room != nil {
		return e.Room, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: room.Label}
	}
	return nil, &NotLoadedError{edge: "room"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Reservation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case reservation.FieldStartsAt, reservation.FieldEndsAt, reservation.FieldCreatedAt, reservation.FieldUpdatedAt, reservation.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case reservation.FieldID, reservation.FieldRoomID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Reservation fields.
func (r *Reservation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case reservation.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				r.ID = *value
			}
		case reservation.FieldStartsAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field starts_at", values[i])
			} else if value.Valid {
				r.StartsAt = value.Time
			}
		case reservation.FieldEndsAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ends_at", values[i])
			} else if value.Valid {
				r.EndsAt = value.Time
			}
		case reservation.FieldRoomID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field room_id", values[i])
			} else if value != nil {
				r.RoomID = *value
			}
		case reservation.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case reservation.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case reservation.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				r.DeletedAt = new(time.Time)
				*r.DeletedAt = value.Time
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Reservation.
// This includes values selected through modifiers, order, etc.
func (r *Reservation) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryRoom queries the "room" edge of the Reservation entity.
func (r *Reservation) QueryRoom() *RoomQuery {
	return NewReservationClient(r.config).QueryRoom(r)
}

// Update returns a builder for updating this Reservation.
// Note that you need to call Reservation.Unwrap() before calling this method if this Reservation
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Reservation) Update() *ReservationUpdateOne {
	return NewReservationClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Reservation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Reservation) Unwrap() *Reservation {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Reservation is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Reservation) String() string {
	var builder strings.Builder
	builder.WriteString("Reservation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("starts_at=")
	builder.WriteString(r.StartsAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ends_at=")
	builder.WriteString(r.EndsAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("room_id=")
	builder.WriteString(fmt.Sprintf("%v", r.RoomID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := r.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Reservations is a parsable slice of Reservation.
type Reservations []*Reservation
