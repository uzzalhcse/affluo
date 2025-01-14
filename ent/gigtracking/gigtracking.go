// Code generated by ent, DO NOT EDIT.

package gigtracking

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the gigtracking type in the database.
	Label = "gig_tracking"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldAffiliateUserID holds the string denoting the affiliate_user_id field in the database.
	FieldAffiliateUserID = "affiliate_user_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldUtmQuery holds the string denoting the utm_query field in the database.
	FieldUtmQuery = "utm_query"
	// FieldLp holds the string denoting the lp field in the database.
	FieldLp = "lp"
	// FieldTrackID holds the string denoting the track_id field in the database.
	FieldTrackID = "track_id"
	// FieldRevenue holds the string denoting the revenue field in the database.
	FieldRevenue = "revenue"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgePublisher holds the string denoting the publisher edge name in mutations.
	EdgePublisher = "publisher"
	// Table holds the table name of the gigtracking in the database.
	Table = "gig_trackings"
	// PublisherTable is the table that holds the publisher relation/edge.
	PublisherTable = "gig_trackings"
	// PublisherInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	PublisherInverseTable = "users"
	// PublisherColumn is the table column denoting the publisher relation/edge.
	PublisherColumn = "gig_tracking_publisher"
)

// Columns holds all SQL columns for gigtracking fields.
var Columns = []string{
	FieldID,
	FieldDate,
	FieldAffiliateUserID,
	FieldType,
	FieldUtmQuery,
	FieldLp,
	FieldTrackID,
	FieldRevenue,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "gig_trackings"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"gig_tracking_publisher",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultDate holds the default value on creation for the "date" field.
	DefaultDate func() time.Time
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType string
	// DefaultRevenue holds the default value on creation for the "revenue" field.
	DefaultRevenue float64
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the GigTracking queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDate orders the results by the date field.
func ByDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDate, opts...).ToFunc()
}

// ByAffiliateUserID orders the results by the affiliate_user_id field.
func ByAffiliateUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAffiliateUserID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByUtmQuery orders the results by the utm_query field.
func ByUtmQuery(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUtmQuery, opts...).ToFunc()
}

// ByLp orders the results by the lp field.
func ByLp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLp, opts...).ToFunc()
}

// ByTrackID orders the results by the track_id field.
func ByTrackID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTrackID, opts...).ToFunc()
}

// ByRevenue orders the results by the revenue field.
func ByRevenue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRevenue, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByPublisherField orders the results by publisher field.
func ByPublisherField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPublisherStep(), sql.OrderByField(field, opts...))
	}
}
func newPublisherStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PublisherInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, PublisherTable, PublisherColumn),
	)
}
