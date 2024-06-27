// Code generated by ent, DO NOT EDIT.

package reservation

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the reservation type in the database.
	Label = "reservation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStartsAt holds the string denoting the starts_at field in the database.
	FieldStartsAt = "starts_at"
	// FieldEndsAt holds the string denoting the ends_at field in the database.
	FieldEndsAt = "ends_at"
	// FieldRoomID holds the string denoting the room_id field in the database.
	FieldRoomID = "room_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeRoom holds the string denoting the room edge name in mutations.
	EdgeRoom = "room"
	// Table holds the table name of the reservation in the database.
	Table = "reservations"
	// RoomTable is the table that holds the room relation/edge.
	RoomTable = "reservations"
	// RoomInverseTable is the table name for the Room entity.
	// It exists in this package in order to avoid circular dependency with the "room" package.
	RoomInverseTable = "rooms"
	// RoomColumn is the table column denoting the room relation/edge.
	RoomColumn = "room_id"
)

// Columns holds all SQL columns for reservation fields.
var Columns = []string{
	FieldID,
	FieldStartsAt,
	FieldEndsAt,
	FieldRoomID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Reservation queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByStartsAt orders the results by the starts_at field.
func ByStartsAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartsAt, opts...).ToFunc()
}

// ByEndsAt orders the results by the ends_at field.
func ByEndsAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndsAt, opts...).ToFunc()
}

// ByRoomID orders the results by the room_id field.
func ByRoomID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoomID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByRoomField orders the results by room field.
func ByRoomField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoomStep(), sql.OrderByField(field, opts...))
	}
}
func newRoomStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoomInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, RoomTable, RoomColumn),
	)
}
