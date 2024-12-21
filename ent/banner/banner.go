// Code generated by ent, DO NOT EDIT.

package banner

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the banner type in the database.
	Label = "banner"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldClickURL holds the string denoting the click_url field in the database.
	FieldClickURL = "click_url"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldAllowedCountries holds the string denoting the allowed_countries field in the database.
	FieldAllowedCountries = "allowed_countries"
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// FieldSmartWeight holds the string denoting the smart_weight field in the database.
	FieldSmartWeight = "smart_weight"
	// FieldLastImpression holds the string denoting the last_impression field in the database.
	FieldLastImpression = "last_impression"
	// FieldStartDate holds the string denoting the start_date field in the database.
	FieldStartDate = "start_date"
	// FieldEndDate holds the string denoting the end_date field in the database.
	FieldEndDate = "end_date"
	// FieldAllowedDevices holds the string denoting the allowed_devices field in the database.
	FieldAllowedDevices = "allowed_devices"
	// FieldAllowedBrowsers holds the string denoting the allowed_browsers field in the database.
	FieldAllowedBrowsers = "allowed_browsers"
	// FieldAllowedOs holds the string denoting the allowed_os field in the database.
	FieldAllowedOs = "allowed_os"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeCampaigns holds the string denoting the campaigns edge name in mutations.
	EdgeCampaigns = "campaigns"
	// EdgeCreatives holds the string denoting the creatives edge name in mutations.
	EdgeCreatives = "creatives"
	// EdgeStats holds the string denoting the stats edge name in mutations.
	EdgeStats = "stats"
	// EdgeLeads holds the string denoting the leads edge name in mutations.
	EdgeLeads = "leads"
	// Table holds the table name of the banner in the database.
	Table = "banners"
	// CampaignsTable is the table that holds the campaigns relation/edge. The primary key declared below.
	CampaignsTable = "campaign_banners"
	// CampaignsInverseTable is the table name for the Campaign entity.
	// It exists in this package in order to avoid circular dependency with the "campaign" package.
	CampaignsInverseTable = "campaigns"
	// CreativesTable is the table that holds the creatives relation/edge.
	CreativesTable = "banner_creatives"
	// CreativesInverseTable is the table name for the BannerCreative entity.
	// It exists in this package in order to avoid circular dependency with the "bannercreative" package.
	CreativesInverseTable = "banner_creatives"
	// CreativesColumn is the table column denoting the creatives relation/edge.
	CreativesColumn = "banner_creatives"
	// StatsTable is the table that holds the stats relation/edge.
	StatsTable = "banner_stats"
	// StatsInverseTable is the table name for the BannerStats entity.
	// It exists in this package in order to avoid circular dependency with the "bannerstats" package.
	StatsInverseTable = "banner_stats"
	// StatsColumn is the table column denoting the stats relation/edge.
	StatsColumn = "banner_stats"
	// LeadsTable is the table that holds the leads relation/edge.
	LeadsTable = "leads"
	// LeadsInverseTable is the table name for the Lead entity.
	// It exists in this package in order to avoid circular dependency with the "lead" package.
	LeadsInverseTable = "leads"
	// LeadsColumn is the table column denoting the leads relation/edge.
	LeadsColumn = "banner_leads"
)

// Columns holds all SQL columns for banner fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldType,
	FieldClickURL,
	FieldSize,
	FieldStatus,
	FieldAllowedCountries,
	FieldWeight,
	FieldSmartWeight,
	FieldLastImpression,
	FieldStartDate,
	FieldEndDate,
	FieldAllowedDevices,
	FieldAllowedBrowsers,
	FieldAllowedOs,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// CampaignsPrimaryKey and CampaignsColumn2 are the table columns denoting the
	// primary key for the campaigns relation (M2M).
	CampaignsPrimaryKey = []string{"campaign_id", "banner_id"}
)

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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultWeight holds the default value on creation for the "weight" field.
	DefaultWeight int
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int64) error
)

// Type defines the type for the "type" enum field.
type Type string

// TypeStatic is the default value of the Type enum.
const DefaultType = TypeStatic

// Type values.
const (
	TypeStatic  Type = "static"
	TypeDynamic Type = "dynamic"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeStatic, TypeDynamic:
		return nil
	default:
		return fmt.Errorf("banner: invalid enum value for type field: %q", _type)
	}
}

// Status defines the type for the "status" enum field.
type Status string

// StatusDraft is the default value of the Status enum.
const DefaultStatus = StatusDraft

// Status values.
const (
	StatusDraft    Status = "draft"
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusDraft, StatusActive, StatusInactive:
		return nil
	default:
		return fmt.Errorf("banner: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Banner queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByClickURL orders the results by the click_url field.
func ByClickURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClickURL, opts...).ToFunc()
}

// BySize orders the results by the size field.
func BySize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSize, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByWeight orders the results by the weight field.
func ByWeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeight, opts...).ToFunc()
}

// BySmartWeight orders the results by the smart_weight field.
func BySmartWeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSmartWeight, opts...).ToFunc()
}

// ByLastImpression orders the results by the last_impression field.
func ByLastImpression(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastImpression, opts...).ToFunc()
}

// ByStartDate orders the results by the start_date field.
func ByStartDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartDate, opts...).ToFunc()
}

// ByEndDate orders the results by the end_date field.
func ByEndDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndDate, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCampaignsCount orders the results by campaigns count.
func ByCampaignsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCampaignsStep(), opts...)
	}
}

// ByCampaigns orders the results by campaigns terms.
func ByCampaigns(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCampaignsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCreativesCount orders the results by creatives count.
func ByCreativesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCreativesStep(), opts...)
	}
}

// ByCreatives orders the results by creatives terms.
func ByCreatives(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreativesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByStatsCount orders the results by stats count.
func ByStatsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStatsStep(), opts...)
	}
}

// ByStats orders the results by stats terms.
func ByStats(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStatsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLeadsCount orders the results by leads count.
func ByLeadsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLeadsStep(), opts...)
	}
}

// ByLeads orders the results by leads terms.
func ByLeads(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLeadsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCampaignsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CampaignsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, CampaignsTable, CampaignsPrimaryKey...),
	)
}
func newCreativesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreativesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CreativesTable, CreativesColumn),
	)
}
func newStatsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StatsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, StatsTable, StatsColumn),
	)
}
func newLeadsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LeadsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LeadsTable, LeadsColumn),
	)
}
