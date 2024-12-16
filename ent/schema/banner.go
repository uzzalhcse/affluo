package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type Banner struct {
	ent.Schema
}

func (Banner) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Positive(),
		field.String("name").NotEmpty(),
		field.String("description").Optional(),
		field.Enum("type").Values(
			"static",
			"dynamic",
		).Default("static"),

		// Content fields
		field.String("click_url").Optional(),

		// Dimensions and styling
		field.String("size"),

		// Status and tracking
		field.Enum("status").Values(
			"draft",
			"active",
			"inactive",
		).Default("draft"),

		// Geo-targeting
		field.JSON("allowed_countries", []string{}).Optional(),

		// Timestamps
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Banner) Edges() []ent.Edge {
	return []ent.Edge{
		// Many-to-many relationship with Campaigns
		edge.From("campaigns", Campaign.Type).
			Ref("banners"),

		// One-to-many relationship with BannerCreatives
		edge.To("creatives", BannerCreative.Type),
	}
}
