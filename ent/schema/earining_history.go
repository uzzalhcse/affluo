package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type EarningHistory struct {
	ent.Schema
}

func (EarningHistory) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Positive(),
		field.Float("amount"),
		field.String("event_type").Comment("lead,account_create,click,impression"),
		field.String("source").Comment("banner,gigs"),
		field.String("track_id").Optional(),
		field.String("date").Default(time.Now().Format("2006-01-02")),
	}
}

func (EarningHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("earning_histories").
			Unique().
			Required(),
	}
}
