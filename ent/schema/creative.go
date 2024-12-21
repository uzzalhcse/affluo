// creative.go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type Creative struct {
	ent.Schema
}

func (Creative) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Positive(),
		// Creative metadata
		field.String("name").Optional(),
		// Asset information
		field.String("image_url").Optional(),
		field.String("size").Optional(),
		// Status and control
		field.Bool("enabled").Default(true),
		// Timestamps
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Creative) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("banners", Banner.Type).
			Ref("creatives").
			Through("banner_creatives", BannerCreative.Type),
	}
}
