// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/campaign"
	"affluo/ent/campaignlink"
	"affluo/ent/track"
	"affluo/ent/user"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Track is the model entity for the Track schema.
type Track struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// IPAddress holds the value of the "ip_address" field.
	IPAddress string `json:"ip_address,omitempty"`
	// UserAgent holds the value of the "user_agent" field.
	UserAgent string `json:"user_agent,omitempty"`
	// DeviceFingerprint holds the value of the "device_fingerprint" field.
	DeviceFingerprint string `json:"device_fingerprint,omitempty"`
	// Referrer holds the value of the "referrer" field.
	Referrer string `json:"referrer,omitempty"`
	// Type holds the value of the "type" field.
	Type track.Type `json:"type,omitempty"`
	// Status holds the value of the "status" field.
	Status track.Status `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// IsUniqueClick holds the value of the "is_unique_click" field.
	IsUniqueClick bool `json:"is_unique_click,omitempty"`
	// AdditionalMetadata holds the value of the "additional_metadata" field.
	AdditionalMetadata map[string]interface{} `json:"additional_metadata,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TrackQuery when eager-loading is set.
	Edges                TrackEdges `json:"edges"`
	campaign_tracks      *int64
	campaign_link_tracks *int64
	user_tracks          *int64
	selectValues         sql.SelectValues
}

// TrackEdges holds the relations/edges for other nodes in the graph.
type TrackEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Campaign holds the value of the campaign edge.
	Campaign *Campaign `json:"campaign,omitempty"`
	// Link holds the value of the link edge.
	Link *CampaignLink `json:"link,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TrackEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// CampaignOrErr returns the Campaign value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TrackEdges) CampaignOrErr() (*Campaign, error) {
	if e.Campaign != nil {
		return e.Campaign, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: campaign.Label}
	}
	return nil, &NotLoadedError{edge: "campaign"}
}

// LinkOrErr returns the Link value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TrackEdges) LinkOrErr() (*CampaignLink, error) {
	if e.Link != nil {
		return e.Link, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: campaignlink.Label}
	}
	return nil, &NotLoadedError{edge: "link"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Track) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case track.FieldAdditionalMetadata:
			values[i] = new([]byte)
		case track.FieldIsUniqueClick:
			values[i] = new(sql.NullBool)
		case track.FieldID:
			values[i] = new(sql.NullInt64)
		case track.FieldIPAddress, track.FieldUserAgent, track.FieldDeviceFingerprint, track.FieldReferrer, track.FieldType, track.FieldStatus:
			values[i] = new(sql.NullString)
		case track.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case track.ForeignKeys[0]: // campaign_tracks
			values[i] = new(sql.NullInt64)
		case track.ForeignKeys[1]: // campaign_link_tracks
			values[i] = new(sql.NullInt64)
		case track.ForeignKeys[2]: // user_tracks
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Track fields.
func (t *Track) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case track.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int64(value.Int64)
		case track.FieldIPAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip_address", values[i])
			} else if value.Valid {
				t.IPAddress = value.String
			}
		case track.FieldUserAgent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_agent", values[i])
			} else if value.Valid {
				t.UserAgent = value.String
			}
		case track.FieldDeviceFingerprint:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device_fingerprint", values[i])
			} else if value.Valid {
				t.DeviceFingerprint = value.String
			}
		case track.FieldReferrer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field referrer", values[i])
			} else if value.Valid {
				t.Referrer = value.String
			}
		case track.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				t.Type = track.Type(value.String)
			}
		case track.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = track.Status(value.String)
			}
		case track.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case track.FieldIsUniqueClick:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_unique_click", values[i])
			} else if value.Valid {
				t.IsUniqueClick = value.Bool
			}
		case track.FieldAdditionalMetadata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field additional_metadata", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.AdditionalMetadata); err != nil {
					return fmt.Errorf("unmarshal field additional_metadata: %w", err)
				}
			}
		case track.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field campaign_tracks", value)
			} else if value.Valid {
				t.campaign_tracks = new(int64)
				*t.campaign_tracks = int64(value.Int64)
			}
		case track.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field campaign_link_tracks", value)
			} else if value.Valid {
				t.campaign_link_tracks = new(int64)
				*t.campaign_link_tracks = int64(value.Int64)
			}
		case track.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_tracks", value)
			} else if value.Valid {
				t.user_tracks = new(int64)
				*t.user_tracks = int64(value.Int64)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Track.
// This includes values selected through modifiers, order, etc.
func (t *Track) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Track entity.
func (t *Track) QueryUser() *UserQuery {
	return NewTrackClient(t.config).QueryUser(t)
}

// QueryCampaign queries the "campaign" edge of the Track entity.
func (t *Track) QueryCampaign() *CampaignQuery {
	return NewTrackClient(t.config).QueryCampaign(t)
}

// QueryLink queries the "link" edge of the Track entity.
func (t *Track) QueryLink() *CampaignLinkQuery {
	return NewTrackClient(t.config).QueryLink(t)
}

// Update returns a builder for updating this Track.
// Note that you need to call Track.Unwrap() before calling this method if this Track
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Track) Update() *TrackUpdateOne {
	return NewTrackClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Track entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Track) Unwrap() *Track {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Track is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Track) String() string {
	var builder strings.Builder
	builder.WriteString("Track(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("ip_address=")
	builder.WriteString(t.IPAddress)
	builder.WriteString(", ")
	builder.WriteString("user_agent=")
	builder.WriteString(t.UserAgent)
	builder.WriteString(", ")
	builder.WriteString("device_fingerprint=")
	builder.WriteString(t.DeviceFingerprint)
	builder.WriteString(", ")
	builder.WriteString("referrer=")
	builder.WriteString(t.Referrer)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", t.Type))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("is_unique_click=")
	builder.WriteString(fmt.Sprintf("%v", t.IsUniqueClick))
	builder.WriteString(", ")
	builder.WriteString("additional_metadata=")
	builder.WriteString(fmt.Sprintf("%v", t.AdditionalMetadata))
	builder.WriteByte(')')
	return builder.String()
}

// Tracks is a parsable slice of Track.
type Tracks []*Track