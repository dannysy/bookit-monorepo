// Code generated by ent, DO NOT EDIT.

package ent

import (
	"bookit/internal/repo/ent/reservation"
	"bookit/internal/repo/ent/room"
	"bookit/internal/repo/ent/schema"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	reservationFields := schema.Reservation{}.Fields()
	_ = reservationFields
	// reservationDescCreatedAt is the schema descriptor for created_at field.
	reservationDescCreatedAt := reservationFields[4].Descriptor()
	// reservation.DefaultCreatedAt holds the default value on creation for the created_at field.
	reservation.DefaultCreatedAt = reservationDescCreatedAt.Default.(func() time.Time)
	// reservationDescUpdatedAt is the schema descriptor for updated_at field.
	reservationDescUpdatedAt := reservationFields[5].Descriptor()
	// reservation.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	reservation.DefaultUpdatedAt = reservationDescUpdatedAt.Default.(func() time.Time)
	// reservationDescID is the schema descriptor for id field.
	reservationDescID := reservationFields[0].Descriptor()
	// reservation.DefaultID holds the default value on creation for the id field.
	reservation.DefaultID = reservationDescID.Default.(func() uuid.UUID)
	roomFields := schema.Room{}.Fields()
	_ = roomFields
	// roomDescName is the schema descriptor for name field.
	roomDescName := roomFields[1].Descriptor()
	// room.NameValidator is a validator for the "name" field. It is called by the builders before save.
	room.NameValidator = roomDescName.Validators[0].(func(string) error)
	// roomDescDescription is the schema descriptor for description field.
	roomDescDescription := roomFields[2].Descriptor()
	// room.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	room.DescriptionValidator = roomDescDescription.Validators[0].(func(string) error)
	// roomDescIsAvailable is the schema descriptor for is_available field.
	roomDescIsAvailable := roomFields[3].Descriptor()
	// room.DefaultIsAvailable holds the default value on creation for the is_available field.
	room.DefaultIsAvailable = roomDescIsAvailable.Default.(bool)
	// roomDescCreatedAt is the schema descriptor for created_at field.
	roomDescCreatedAt := roomFields[4].Descriptor()
	// room.DefaultCreatedAt holds the default value on creation for the created_at field.
	room.DefaultCreatedAt = roomDescCreatedAt.Default.(func() time.Time)
	// roomDescUpdatedAt is the schema descriptor for updated_at field.
	roomDescUpdatedAt := roomFields[5].Descriptor()
	// room.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	room.DefaultUpdatedAt = roomDescUpdatedAt.Default.(func() time.Time)
	// roomDescID is the schema descriptor for id field.
	roomDescID := roomFields[0].Descriptor()
	// room.DefaultID holds the default value on creation for the id field.
	room.DefaultID = roomDescID.Default.(func() uuid.UUID)
}
