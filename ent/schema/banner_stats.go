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
		field.Float("ctr").Optional(),
		field.Float("conversion_rate").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (BannerStats) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("banner", Banner.Type).
			Ref("stats").
			Unique(),
	}
}

func (BannerStats) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("date"),
		index.Fields("date").Edges("banner").Unique(),
	}
}
