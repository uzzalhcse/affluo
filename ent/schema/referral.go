package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Referral schema to track affiliate referrals
type Referral struct {
	ent.Schema
}

func (Referral) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Positive(),
		field.Enum("status").Values("pending", "approved", "rejected"),
		field.Float("commission_amount"),
		field.Time("created_at").Default(time.Now),
		field.Time("processed_at").Optional(),
	}
}

func (Referral) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("referrer", User.Type).Ref("referrals").Unique(),
		edge.From("campaign", Campaign.Type).Ref("referrals").Unique(),
	}
}
