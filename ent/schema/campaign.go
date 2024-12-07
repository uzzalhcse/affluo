package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Campaign schema for tracking marketing campaigns
type Campaign struct {
	ent.Schema
}

func (Campaign) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Positive(),
		field.String("name").NotEmpty(),
		field.String("description").Optional(),
		field.Enum("type").Values(
			"cpc",       // Cost Per Click
			"cpa",       // Cost Per Action
			"cpm",       // Cost Per Mille
			"rev_share", // Revenue Share
		),
		field.Float("payout_rate"),
		field.Time("start_date"),
		field.Time("end_date").Optional(),
		field.Enum("status").Values("active", "paused", "completed"),
		field.String("tracking_url"),
		field.String("unique_code").Unique(),
	}
}

func (Campaign) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("campaigns").Unique(),
		edge.To("links", CampaignLink.Type),
		edge.To("tracks", Track.Type),
		edge.To("referrals", Referral.Type), // Add this to complete the back-reference
	}
}
