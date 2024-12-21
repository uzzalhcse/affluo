// Code generated by ent, DO NOT EDIT.

package bannerstats

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the bannerstats type in the database.
	Label = "banner_stats"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldImpressions holds the string denoting the impressions field in the database.
	FieldImpressions = "impressions"
	// FieldClicks holds the string denoting the clicks field in the database.
	FieldClicks = "clicks"
	// FieldLeads holds the string denoting the leads field in the database.
	FieldLeads = "leads"
	// FieldCtr holds the string denoting the ctr field in the database.
	FieldCtr = "ctr"
	// FieldConversionRate holds the string denoting the conversion_rate field in the database.
	FieldConversionRate = "conversion_rate"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeBanner holds the string denoting the banner edge name in mutations.
	EdgeBanner = "banner"
	// Table holds the table name of the bannerstats in the database.
	Table = "banner_stats"
	// BannerTable is the table that holds the banner relation/edge.
	BannerTable = "banner_stats"
	// BannerInverseTable is the table name for the Banner entity.
	// It exists in this package in order to avoid circular dependency with the "banner" package.
	BannerInverseTable = "banners"
	// BannerColumn is the table column denoting the banner relation/edge.
	BannerColumn = "banner_stats"
)

// Columns holds all SQL columns for bannerstats fields.
var Columns = []string{
	FieldID,
	FieldDate,
	FieldImpressions,
	FieldClicks,
	FieldLeads,
	FieldCtr,
	FieldConversionRate,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "banner_stats"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"banner_stats",
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
	// DefaultImpressions holds the default value on creation for the "impressions" field.
	DefaultImpressions int64
	// DefaultClicks holds the default value on creation for the "clicks" field.
	DefaultClicks int64
	// DefaultLeads holds the default value on creation for the "leads" field.
	DefaultLeads int64
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the BannerStats queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDate orders the results by the date field.
func ByDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDate, opts...).ToFunc()
}

// ByImpressions orders the results by the impressions field.
func ByImpressions(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImpressions, opts...).ToFunc()
}

// ByClicks orders the results by the clicks field.
func ByClicks(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClicks, opts...).ToFunc()
}

// ByLeads orders the results by the leads field.
func ByLeads(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLeads, opts...).ToFunc()
}

// ByCtr orders the results by the ctr field.
func ByCtr(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCtr, opts...).ToFunc()
}

// ByConversionRate orders the results by the conversion_rate field.
func ByConversionRate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConversionRate, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByBannerField orders the results by banner field.
func ByBannerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBannerStep(), sql.OrderByField(field, opts...))
	}
}
func newBannerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BannerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, BannerTable, BannerColumn),
	)
}