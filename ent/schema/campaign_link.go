package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// CampaignLink schema for unique tracking links
type CampaignLink struct {
	ent.Schema
}

func (CampaignLink) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Positive(),
		field.String("unique_code").Unique(),
		field.String("original_url"),
		field.String("tracking_url"),
		field.Time("created_at").Default(time.Now),
		field.Bool("is_active").Default(true),
	}
}

func (CampaignLink) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("campaign", Campaign.Type).Ref("links").Unique(),
		edge.To("tracks", Track.Type),
	}
}
