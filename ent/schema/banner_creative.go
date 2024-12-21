// banner_creative.go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type BannerCreative struct {
	ent.Schema
}

func (BannerCreative) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("banner_id"),
		field.Int64("creative_id"),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		// You can add additional fields specific to the relationship here
		field.Bool("is_primary").Default(false),
		field.Int("display_order").Optional(),
	}
}

func (BannerCreative) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("banner", Banner.Type).
			Field("banner_id").
			Required().
			Unique(),
		edge.To("creative", Creative.Type).
			Field("creative_id").
			Required().
			Unique(),
	}
}
