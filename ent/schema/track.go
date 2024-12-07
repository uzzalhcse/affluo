package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Track schema for tracking individual clicks/impressions
type Track struct {
	ent.Schema
}

func (Track) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Positive(),
		field.String("ip_address"),
		field.String("user_agent"),
		field.String("device_fingerprint"),
		field.String("referrer").Optional(),
		field.Enum("type").Values("click", "impression", "conversion"),
		field.Enum("status").Values("valid", "suspected_fraud", "blacklisted"),
		field.Time("created_at").Default(time.Now),
		field.Bool("is_unique_click").Default(false),
		field.JSON("additional_metadata", map[string]interface{}{}).Optional(),
	}
}

func (Track) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("tracks").Unique(),
		edge.From("campaign", Campaign.Type).Ref("tracks").Unique(),
		edge.From("link", CampaignLink.Type).Ref("tracks").Unique(),
	}
}
