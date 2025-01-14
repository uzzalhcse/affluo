// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/creative"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Creative is the model entity for the Creative schema.
type Creative struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// ImageURL holds the value of the "image_url" field.
	ImageURL string `json:"image_url,omitempty"`
	// Size holds the value of the "size" field.
	Size string `json:"size,omitempty"`
	// Enabled holds the value of the "enabled" field.
	Enabled bool `json:"enabled,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CreativeQuery when eager-loading is set.
	Edges        CreativeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CreativeEdges holds the relations/edges for other nodes in the graph.
type CreativeEdges struct {
	// Banners holds the value of the banners edge.
	Banners []*Banner `json:"banners,omitempty"`
	// BannerCreatives holds the value of the banner_creatives edge.
	BannerCreatives []*BannerCreative `json:"banner_creatives,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// BannersOrErr returns the Banners value or an error if the edge
// was not loaded in eager-loading.
func (e CreativeEdges) BannersOrErr() ([]*Banner, error) {
	if e.loadedTypes[0] {
		return e.Banners, nil
	}
	return nil, &NotLoadedError{edge: "banners"}
}

// BannerCreativesOrErr returns the BannerCreatives value or an error if the edge
// was not loaded in eager-loading.
func (e CreativeEdges) BannerCreativesOrErr() ([]*BannerCreative, error) {
	if e.loadedTypes[1] {
		return e.BannerCreatives, nil
	}
	return nil, &NotLoadedError{edge: "banner_creatives"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Creative) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case creative.FieldEnabled:
			values[i] = new(sql.NullBool)
		case creative.FieldID:
			values[i] = new(sql.NullInt64)
		case creative.FieldName, creative.FieldImageURL, creative.FieldSize:
			values[i] = new(sql.NullString)
		case creative.FieldCreatedAt, creative.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Creative fields.
func (c *Creative) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case creative.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int64(value.Int64)
		case creative.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case creative.FieldImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_url", values[i])
			} else if value.Valid {
				c.ImageURL = value.String
			}
		case creative.FieldSize:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				c.Size = value.String
			}
		case creative.FieldEnabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enabled", values[i])
			} else if value.Valid {
				c.Enabled = value.Bool
			}
		case creative.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case creative.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Creative.
// This includes values selected through modifiers, order, etc.
func (c *Creative) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryBanners queries the "banners" edge of the Creative entity.
func (c *Creative) QueryBanners() *BannerQuery {
	return NewCreativeClient(c.config).QueryBanners(c)
}

// QueryBannerCreatives queries the "banner_creatives" edge of the Creative entity.
func (c *Creative) QueryBannerCreatives() *BannerCreativeQuery {
	return NewCreativeClient(c.config).QueryBannerCreatives(c)
}

// Update returns a builder for updating this Creative.
// Note that you need to call Creative.Unwrap() before calling this method if this Creative
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Creative) Update() *CreativeUpdateOne {
	return NewCreativeClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Creative entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Creative) Unwrap() *Creative {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Creative is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Creative) String() string {
	var builder strings.Builder
	builder.WriteString("Creative(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("image_url=")
	builder.WriteString(c.ImageURL)
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(c.Size)
	builder.WriteString(", ")
	builder.WriteString("enabled=")
	builder.WriteString(fmt.Sprintf("%v", c.Enabled))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Creatives is a parsable slice of Creative.
type Creatives []*Creative
