// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/banner"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Banner is the model entity for the Banner schema.
type Banner struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Type holds the value of the "type" field.
	Type banner.Type `json:"type,omitempty"`
	// ClickURL holds the value of the "click_url" field.
	ClickURL string `json:"click_url,omitempty"`
	// Size holds the value of the "size" field.
	Size string `json:"size,omitempty"`
	// Status holds the value of the "status" field.
	Status banner.Status `json:"status,omitempty"`
	// AllowedCountries holds the value of the "allowed_countries" field.
	AllowedCountries []string `json:"allowed_countries,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BannerQuery when eager-loading is set.
	Edges        BannerEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BannerEdges holds the relations/edges for other nodes in the graph.
type BannerEdges struct {
	// Campaigns holds the value of the campaigns edge.
	Campaigns []*Campaign `json:"campaigns,omitempty"`
	// Creatives holds the value of the creatives edge.
	Creatives []*BannerCreative `json:"creatives,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CampaignsOrErr returns the Campaigns value or an error if the edge
// was not loaded in eager-loading.
func (e BannerEdges) CampaignsOrErr() ([]*Campaign, error) {
	if e.loadedTypes[0] {
		return e.Campaigns, nil
	}
	return nil, &NotLoadedError{edge: "campaigns"}
}

// CreativesOrErr returns the Creatives value or an error if the edge
// was not loaded in eager-loading.
func (e BannerEdges) CreativesOrErr() ([]*BannerCreative, error) {
	if e.loadedTypes[1] {
		return e.Creatives, nil
	}
	return nil, &NotLoadedError{edge: "creatives"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Banner) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case banner.FieldAllowedCountries:
			values[i] = new([]byte)
		case banner.FieldID:
			values[i] = new(sql.NullInt64)
		case banner.FieldName, banner.FieldDescription, banner.FieldType, banner.FieldClickURL, banner.FieldSize, banner.FieldStatus:
			values[i] = new(sql.NullString)
		case banner.FieldCreatedAt, banner.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Banner fields.
func (b *Banner) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case banner.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int64(value.Int64)
		case banner.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				b.Name = value.String
			}
		case banner.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				b.Description = value.String
			}
		case banner.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				b.Type = banner.Type(value.String)
			}
		case banner.FieldClickURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field click_url", values[i])
			} else if value.Valid {
				b.ClickURL = value.String
			}
		case banner.FieldSize:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				b.Size = value.String
			}
		case banner.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				b.Status = banner.Status(value.String)
			}
		case banner.FieldAllowedCountries:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field allowed_countries", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &b.AllowedCountries); err != nil {
					return fmt.Errorf("unmarshal field allowed_countries: %w", err)
				}
			}
		case banner.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case banner.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Banner.
// This includes values selected through modifiers, order, etc.
func (b *Banner) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryCampaigns queries the "campaigns" edge of the Banner entity.
func (b *Banner) QueryCampaigns() *CampaignQuery {
	return NewBannerClient(b.config).QueryCampaigns(b)
}

// QueryCreatives queries the "creatives" edge of the Banner entity.
func (b *Banner) QueryCreatives() *BannerCreativeQuery {
	return NewBannerClient(b.config).QueryCreatives(b)
}

// Update returns a builder for updating this Banner.
// Note that you need to call Banner.Unwrap() before calling this method if this Banner
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Banner) Update() *BannerUpdateOne {
	return NewBannerClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Banner entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Banner) Unwrap() *Banner {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Banner is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Banner) String() string {
	var builder strings.Builder
	builder.WriteString("Banner(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("name=")
	builder.WriteString(b.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(b.Description)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", b.Type))
	builder.WriteString(", ")
	builder.WriteString("click_url=")
	builder.WriteString(b.ClickURL)
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(b.Size)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", b.Status))
	builder.WriteString(", ")
	builder.WriteString("allowed_countries=")
	builder.WriteString(fmt.Sprintf("%v", b.AllowedCountries))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Banners is a parsable slice of Banner.
type Banners []*Banner
