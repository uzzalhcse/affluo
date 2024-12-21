// Code generated by ent, DO NOT EDIT.

package bannercreative

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the bannercreative type in the database.
	Label = "banner_creative"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBannerID holds the string denoting the banner_id field in the database.
	FieldBannerID = "banner_id"
	// FieldCreativeID holds the string denoting the creative_id field in the database.
	FieldCreativeID = "creative_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldIsPrimary holds the string denoting the is_primary field in the database.
	FieldIsPrimary = "is_primary"
	// FieldDisplayOrder holds the string denoting the display_order field in the database.
	FieldDisplayOrder = "display_order"
	// EdgeBanner holds the string denoting the banner edge name in mutations.
	EdgeBanner = "banner"
	// EdgeCreative holds the string denoting the creative edge name in mutations.
	EdgeCreative = "creative"
	// Table holds the table name of the bannercreative in the database.
	Table = "banner_creatives"
	// BannerTable is the table that holds the banner relation/edge.
	BannerTable = "banner_creatives"
	// BannerInverseTable is the table name for the Banner entity.
	// It exists in this package in order to avoid circular dependency with the "banner" package.
	BannerInverseTable = "banners"
	// BannerColumn is the table column denoting the banner relation/edge.
	BannerColumn = "banner_id"
	// CreativeTable is the table that holds the creative relation/edge.
	CreativeTable = "banner_creatives"
	// CreativeInverseTable is the table name for the Creative entity.
	// It exists in this package in order to avoid circular dependency with the "creative" package.
	CreativeInverseTable = "creatives"
	// CreativeColumn is the table column denoting the creative relation/edge.
	CreativeColumn = "creative_id"
)

// Columns holds all SQL columns for bannercreative fields.
var Columns = []string{
	FieldID,
	FieldBannerID,
	FieldCreativeID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldIsPrimary,
	FieldDisplayOrder,
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
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultIsPrimary holds the default value on creation for the "is_primary" field.
	DefaultIsPrimary bool
)

// OrderOption defines the ordering options for the BannerCreative queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByBannerID orders the results by the banner_id field.
func ByBannerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBannerID, opts...).ToFunc()
}

// ByCreativeID orders the results by the creative_id field.
func ByCreativeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreativeID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByIsPrimary orders the results by the is_primary field.
func ByIsPrimary(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsPrimary, opts...).ToFunc()
}

// ByDisplayOrder orders the results by the display_order field.
func ByDisplayOrder(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisplayOrder, opts...).ToFunc()
}

// ByBannerField orders the results by banner field.
func ByBannerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBannerStep(), sql.OrderByField(field, opts...))
	}
}

// ByCreativeField orders the results by creative field.
func ByCreativeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreativeStep(), sql.OrderByField(field, opts...))
	}
}
func newBannerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BannerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, BannerTable, BannerColumn),
	)
}
func newCreativeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreativeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CreativeTable, CreativeColumn),
	)
}
