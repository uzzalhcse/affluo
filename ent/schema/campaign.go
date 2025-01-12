package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

type Campaign struct {
	ent.Schema
}

func (Campaign) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Positive(),
		field.String("name").NotEmpty().Unique(),
		field.String("description").Optional(),
		field.String("unique_code").Unique(),

		// Campaign Type
		field.Enum("type").Values(
			"sale",
			"lead",
			"click",
			"subscription",
		).Default("sale"),

		// Commission Structure
		field.Enum("commission_type").Values(
			"flat_rate",
			"percentage",
			"tiered",
		).Default("percentage"),

		// Commission Details
		field.Float("base_commission_rate").Default(0.0),
		field.JSON("commission_tiers", []CommissionTier{}).Optional(),

		// Targeting
		field.String("target_geography").Optional(),
		field.JSON("target_demographics", map[string]interface{}{}).Optional(),

		// Timing
		field.Time("start_date"),
		field.Time("end_date").Optional(),

		// Status
		field.Enum("status").Values(
			"draft",
			"active",
			"paused",
			"completed",
		).Default("draft"),

		// Tracking
		field.String("tracking_url").Optional(),

		// Performance Metrics
		field.Int("total_clicks").Default(0),
		field.Int("total_conversions").Default(0),
		field.Float("total_revenue").Default(0.0),
		field.Float("conversion_rate").Default(0.0),

		// Timestamps
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Campaign) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("campaigns").Unique(),
		edge.To("links", CampaignLink.Type),
		edge.To("banners", Banner.Type),
	}
}

func (Campaign) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
		index.Fields("unique_code"),
		index.Fields("type"),
		index.Fields("status"),
	}
}

// CommissionTier defines a tiered commission structure
type CommissionTier struct {
	MinThreshold float64 `json:"min_threshold"`
	MaxThreshold float64 `json:"max_threshold"`
	Rate         float64 `json:"rate"`
}
