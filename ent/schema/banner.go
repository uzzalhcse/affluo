// banner.go
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
		field.Int("weight").Default(1),
		field.Float("smart_weight").Optional(),
		field.Time("last_impression").Optional(),
		field.Time("start_date").Optional(),
		field.Time("end_date").Optional(),
		// Device targeting
		field.JSON("allowed_devices", []string{}).Optional(),
		field.JSON("allowed_browsers", []string{}).Optional(),
		field.JSON("allowed_os", []string{}).Optional(),
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
		// Many-to-many relationship with Creatives
		edge.To("creatives", Creative.Type).
			Through("banner_creatives", BannerCreative.Type),
		edge.To("stats", BannerStats.Type),
		edge.To("leads", Lead.Type),
	}
}
