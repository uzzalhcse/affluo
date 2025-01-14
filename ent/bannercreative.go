// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/banner"
	"affluo/ent/bannercreative"
	"affluo/ent/creative"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// BannerCreative is the model entity for the BannerCreative schema.
type BannerCreative struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BannerID holds the value of the "banner_id" field.
	BannerID int64 `json:"banner_id,omitempty"`
	// CreativeID holds the value of the "creative_id" field.
	CreativeID int64 `json:"creative_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// IsPrimary holds the value of the "is_primary" field.
	IsPrimary bool `json:"is_primary,omitempty"`
	// DisplayOrder holds the value of the "display_order" field.
	DisplayOrder int `json:"display_order,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BannerCreativeQuery when eager-loading is set.
	Edges        BannerCreativeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BannerCreativeEdges holds the relations/edges for other nodes in the graph.
type BannerCreativeEdges struct {
	// Banner holds the value of the banner edge.
	Banner *Banner `json:"banner,omitempty"`
	// Creative holds the value of the creative edge.
	Creative *Creative `json:"creative,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// BannerOrErr returns the Banner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BannerCreativeEdges) BannerOrErr() (*Banner, error) {
	if e.Banner != nil {
		return e.Banner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: banner.Label}
	}
	return nil, &NotLoadedError{edge: "banner"}
}

// CreativeOrErr returns the Creative value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BannerCreativeEdges) CreativeOrErr() (*Creative, error) {
	if e.Creative != nil {
		return e.Creative, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: creative.Label}
	}
	return nil, &NotLoadedError{edge: "creative"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BannerCreative) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bannercreative.FieldIsPrimary:
			values[i] = new(sql.NullBool)
		case bannercreative.FieldID, bannercreative.FieldBannerID, bannercreative.FieldCreativeID, bannercreative.FieldDisplayOrder:
			values[i] = new(sql.NullInt64)
		case bannercreative.FieldCreatedAt, bannercreative.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BannerCreative fields.
func (bc *BannerCreative) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bannercreative.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bc.ID = int(value.Int64)
		case bannercreative.FieldBannerID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field banner_id", values[i])
			} else if value.Valid {
				bc.BannerID = value.Int64
			}
		case bannercreative.FieldCreativeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field creative_id", values[i])
			} else if value.Valid {
				bc.CreativeID = value.Int64
			}
		case bannercreative.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				bc.CreatedAt = value.Time
			}
		case bannercreative.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				bc.UpdatedAt = value.Time
			}
		case bannercreative.FieldIsPrimary:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_primary", values[i])
			} else if value.Valid {
				bc.IsPrimary = value.Bool
			}
		case bannercreative.FieldDisplayOrder:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field display_order", values[i])
			} else if value.Valid {
				bc.DisplayOrder = int(value.Int64)
			}
		default:
			bc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BannerCreative.
// This includes values selected through modifiers, order, etc.
func (bc *BannerCreative) Value(name string) (ent.Value, error) {
	return bc.selectValues.Get(name)
}

// QueryBanner queries the "banner" edge of the BannerCreative entity.
func (bc *BannerCreative) QueryBanner() *BannerQuery {
	return NewBannerCreativeClient(bc.config).QueryBanner(bc)
}

// QueryCreative queries the "creative" edge of the BannerCreative entity.
func (bc *BannerCreative) QueryCreative() *CreativeQuery {
	return NewBannerCreativeClient(bc.config).QueryCreative(bc)
}

// Update returns a builder for updating this BannerCreative.
// Note that you need to call BannerCreative.Unwrap() before calling this method if this BannerCreative
// was returned from a transaction, and the transaction was committed or rolled back.
func (bc *BannerCreative) Update() *BannerCreativeUpdateOne {
	return NewBannerCreativeClient(bc.config).UpdateOne(bc)
}

// Unwrap unwraps the BannerCreative entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bc *BannerCreative) Unwrap() *BannerCreative {
	_tx, ok := bc.config.driver.(*txDriver)
	if !ok {
		panic("ent: BannerCreative is not a transactional entity")
	}
	bc.config.driver = _tx.drv
	return bc
}

// String implements the fmt.Stringer.
func (bc *BannerCreative) String() string {
	var builder strings.Builder
	builder.WriteString("BannerCreative(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bc.ID))
	builder.WriteString("banner_id=")
	builder.WriteString(fmt.Sprintf("%v", bc.BannerID))
	builder.WriteString(", ")
	builder.WriteString("creative_id=")
	builder.WriteString(fmt.Sprintf("%v", bc.CreativeID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(bc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(bc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("is_primary=")
	builder.WriteString(fmt.Sprintf("%v", bc.IsPrimary))
	builder.WriteString(", ")
	builder.WriteString("display_order=")
	builder.WriteString(fmt.Sprintf("%v", bc.DisplayOrder))
	builder.WriteByte(')')
	return builder.String()
}

// BannerCreatives is a parsable slice of BannerCreative.
type BannerCreatives []*BannerCreative
