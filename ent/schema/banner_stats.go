// schema/stats.go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// BannerStats holds the daily stats for each banner
type BannerStats struct {
	ent.Schema
}

func (BannerStats) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Time("date").Default(time.Now),
		field.Int64("impressions").Default(0),
		field.Int64("clicks").Default(0),
		field.Int64("leads").Default(0),
		field.Float("earnings").Default(0),
		field.Float("ctr").Optional(),
		field.Float("conversion_rate").Optional(),
		field.String("device_type").Optional(),
		field.String("browser").Optional(),
		field.String("os").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (BannerStats) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("banner", Banner.Type).Ref("stats").Unique(),
		edge.From("publisher", User.Type).Ref("stats").Unique(),
	}
}

func (BannerStats) Indexes() []ent.Index {
	return []ent.Index{
		// Unique index for date + banner + publisher
		index.Fields("date").Edges("banner", "publisher").Unique(),

		// Additional indexes for optimization
		index.Edges("banner"),
		index.Edges("publisher"),
	}
}
