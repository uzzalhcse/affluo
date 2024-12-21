// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/banner"
	"affluo/ent/bannerstats"
	"affluo/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// BannerStats is the model entity for the BannerStats schema.
type BannerStats struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Date holds the value of the "date" field.
	Date time.Time `json:"date,omitempty"`
	// Impressions holds the value of the "impressions" field.
	Impressions int64 `json:"impressions,omitempty"`
	// Clicks holds the value of the "clicks" field.
	Clicks int64 `json:"clicks,omitempty"`
	// Leads holds the value of the "leads" field.
	Leads int64 `json:"leads,omitempty"`
	// Earnings holds the value of the "earnings" field.
	Earnings float64 `json:"earnings,omitempty"`
	// Ctr holds the value of the "ctr" field.
	Ctr float64 `json:"ctr,omitempty"`
	// ConversionRate holds the value of the "conversion_rate" field.
	ConversionRate float64 `json:"conversion_rate,omitempty"`
	// DeviceType holds the value of the "device_type" field.
	DeviceType string `json:"device_type,omitempty"`
	// Browser holds the value of the "browser" field.
	Browser string `json:"browser,omitempty"`
	// Os holds the value of the "os" field.
	Os string `json:"os,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BannerStatsQuery when eager-loading is set.
	Edges        BannerStatsEdges `json:"edges"`
	banner_stats *int64
	user_stats   *int64
	selectValues sql.SelectValues
}

// BannerStatsEdges holds the relations/edges for other nodes in the graph.
type BannerStatsEdges struct {
	// Banner holds the value of the banner edge.
	Banner *Banner `json:"banner,omitempty"`
	// Publisher holds the value of the publisher edge.
	Publisher *User `json:"publisher,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// BannerOrErr returns the Banner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BannerStatsEdges) BannerOrErr() (*Banner, error) {
	if e.Banner != nil {
		return e.Banner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: banner.Label}
	}
	return nil, &NotLoadedError{edge: "banner"}
}

// PublisherOrErr returns the Publisher value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BannerStatsEdges) PublisherOrErr() (*User, error) {
	if e.Publisher != nil {
		return e.Publisher, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "publisher"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BannerStats) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bannerstats.FieldEarnings, bannerstats.FieldCtr, bannerstats.FieldConversionRate:
			values[i] = new(sql.NullFloat64)
		case bannerstats.FieldID, bannerstats.FieldImpressions, bannerstats.FieldClicks, bannerstats.FieldLeads:
			values[i] = new(sql.NullInt64)
		case bannerstats.FieldDeviceType, bannerstats.FieldBrowser, bannerstats.FieldOs:
			values[i] = new(sql.NullString)
		case bannerstats.FieldDate, bannerstats.FieldCreatedAt, bannerstats.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case bannerstats.ForeignKeys[0]: // banner_stats
			values[i] = new(sql.NullInt64)
		case bannerstats.ForeignKeys[1]: // user_stats
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BannerStats fields.
func (bs *BannerStats) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bannerstats.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bs.ID = int64(value.Int64)
		case bannerstats.FieldDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				bs.Date = value.Time
			}
		case bannerstats.FieldImpressions:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field impressions", values[i])
			} else if value.Valid {
				bs.Impressions = value.Int64
			}
		case bannerstats.FieldClicks:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field clicks", values[i])
			} else if value.Valid {
				bs.Clicks = value.Int64
			}
		case bannerstats.FieldLeads:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field leads", values[i])
			} else if value.Valid {
				bs.Leads = value.Int64
			}
		case bannerstats.FieldEarnings:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field earnings", values[i])
			} else if value.Valid {
				bs.Earnings = value.Float64
			}
		case bannerstats.FieldCtr:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field ctr", values[i])
			} else if value.Valid {
				bs.Ctr = value.Float64
			}
		case bannerstats.FieldConversionRate:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field conversion_rate", values[i])
			} else if value.Valid {
				bs.ConversionRate = value.Float64
			}
		case bannerstats.FieldDeviceType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device_type", values[i])
			} else if value.Valid {
				bs.DeviceType = value.String
			}
		case bannerstats.FieldBrowser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field browser", values[i])
			} else if value.Valid {
				bs.Browser = value.String
			}
		case bannerstats.FieldOs:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field os", values[i])
			} else if value.Valid {
				bs.Os = value.String
			}
		case bannerstats.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				bs.CreatedAt = value.Time
			}
		case bannerstats.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				bs.UpdatedAt = value.Time
			}
		case bannerstats.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field banner_stats", value)
			} else if value.Valid {
				bs.banner_stats = new(int64)
				*bs.banner_stats = int64(value.Int64)
			}
		case bannerstats.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_stats", value)
			} else if value.Valid {
				bs.user_stats = new(int64)
				*bs.user_stats = int64(value.Int64)
			}
		default:
			bs.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BannerStats.
// This includes values selected through modifiers, order, etc.
func (bs *BannerStats) Value(name string) (ent.Value, error) {
	return bs.selectValues.Get(name)
}

// QueryBanner queries the "banner" edge of the BannerStats entity.
func (bs *BannerStats) QueryBanner() *BannerQuery {
	return NewBannerStatsClient(bs.config).QueryBanner(bs)
}

// QueryPublisher queries the "publisher" edge of the BannerStats entity.
func (bs *BannerStats) QueryPublisher() *UserQuery {
	return NewBannerStatsClient(bs.config).QueryPublisher(bs)
}

// Update returns a builder for updating this BannerStats.
// Note that you need to call BannerStats.Unwrap() before calling this method if this BannerStats
// was returned from a transaction, and the transaction was committed or rolled back.
func (bs *BannerStats) Update() *BannerStatsUpdateOne {
	return NewBannerStatsClient(bs.config).UpdateOne(bs)
}

// Unwrap unwraps the BannerStats entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bs *BannerStats) Unwrap() *BannerStats {
	_tx, ok := bs.config.driver.(*txDriver)
	if !ok {
		panic("ent: BannerStats is not a transactional entity")
	}
	bs.config.driver = _tx.drv
	return bs
}

// String implements the fmt.Stringer.
func (bs *BannerStats) String() string {
	var builder strings.Builder
	builder.WriteString("BannerStats(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bs.ID))
	builder.WriteString("date=")
	builder.WriteString(bs.Date.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("impressions=")
	builder.WriteString(fmt.Sprintf("%v", bs.Impressions))
	builder.WriteString(", ")
	builder.WriteString("clicks=")
	builder.WriteString(fmt.Sprintf("%v", bs.Clicks))
	builder.WriteString(", ")
	builder.WriteString("leads=")
	builder.WriteString(fmt.Sprintf("%v", bs.Leads))
	builder.WriteString(", ")
	builder.WriteString("earnings=")
	builder.WriteString(fmt.Sprintf("%v", bs.Earnings))
	builder.WriteString(", ")
	builder.WriteString("ctr=")
	builder.WriteString(fmt.Sprintf("%v", bs.Ctr))
	builder.WriteString(", ")
	builder.WriteString("conversion_rate=")
	builder.WriteString(fmt.Sprintf("%v", bs.ConversionRate))
	builder.WriteString(", ")
	builder.WriteString("device_type=")
	builder.WriteString(bs.DeviceType)
	builder.WriteString(", ")
	builder.WriteString("browser=")
	builder.WriteString(bs.Browser)
	builder.WriteString(", ")
	builder.WriteString("os=")
	builder.WriteString(bs.Os)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(bs.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(bs.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// BannerStatsSlice is a parsable slice of BannerStats.
type BannerStatsSlice []*BannerStats
