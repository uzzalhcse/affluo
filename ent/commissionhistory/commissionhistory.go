// Code generated by ent, DO NOT EDIT.

package commissionhistory

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the commissionhistory type in the database.
	Label = "commission_history"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldAffiliateUserID holds the string denoting the affiliate_user_id field in the database.
	FieldAffiliateUserID = "affiliate_user_id"
	// FieldTrxID holds the string denoting the trx_id field in the database.
	FieldTrxID = "trx_id"
	// FieldTrackID holds the string denoting the track_id field in the database.
	FieldTrackID = "track_id"
	// FieldCommissionRate holds the string denoting the commission_rate field in the database.
	FieldCommissionRate = "commission_rate"
	// FieldIsFirstOrder holds the string denoting the is_first_order field in the database.
	FieldIsFirstOrder = "is_first_order"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the commissionhistory in the database.
	Table = "commission_histories"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "commission_histories"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_commission_histories"
)

// Columns holds all SQL columns for commissionhistory fields.
var Columns = []string{
	FieldID,
	FieldAmount,
	FieldAffiliateUserID,
	FieldTrxID,
	FieldTrackID,
	FieldCommissionRate,
	FieldIsFirstOrder,
	FieldDate,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "commission_histories"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_commission_histories",
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
	DefaultDate string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int64) error
)

// OrderOption defines the ordering options for the CommissionHistory queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByAffiliateUserID orders the results by the affiliate_user_id field.
func ByAffiliateUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAffiliateUserID, opts...).ToFunc()
}

// ByTrxID orders the results by the trx_id field.
func ByTrxID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTrxID, opts...).ToFunc()
}

// ByTrackID orders the results by the track_id field.
func ByTrackID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTrackID, opts...).ToFunc()
}

// ByCommissionRate orders the results by the commission_rate field.
func ByCommissionRate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCommissionRate, opts...).ToFunc()
}

// ByIsFirstOrder orders the results by the is_first_order field.
func ByIsFirstOrder(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsFirstOrder, opts...).ToFunc()
}

// ByDate orders the results by the date field.
func ByDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDate, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
