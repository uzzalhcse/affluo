package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type Lead struct {
	ent.Schema
}

func (Lead) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.String("reference_id").Optional(), // External reference ID
		field.Enum("type").Values("product", "service"),
		field.Float("amount").Optional(),
		field.String("currency").Default("USD"),
		field.String("ip_address").Optional(),
		field.String("user_agent").Optional(),
		field.JSON("metadata", map[string]interface{}{}).Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

func (Lead) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("banner", Banner.Type).
			Ref("leads").
			Unique(),
	}
}
