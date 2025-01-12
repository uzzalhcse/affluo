// schema/gig_tracking.go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// GigTracking holds the daily stats for each banner
type GigTracking struct {
	ent.Schema
}

func (GigTracking) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Time("date").Default(time.Now),
		field.String("affiliate_user_id").Optional(),
		field.String("type").Default("services"),
		field.String("utm_query").Optional(),
		field.String("lp").Optional(),
		field.String("track_id").Optional(),
		field.Float("revenue").Default(0), // in cents
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (GigTracking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("publisher", User.Type).
			Unique().
			Required(),
	}
}

func (GigTracking) Indexes() []ent.Index {
	return []ent.Index{
		// Unique index for date + publisher
		index.Fields("date", "lp", "type", "track_id").Edges("publisher").Unique(),
		index.Edges("publisher"),
	}
}
