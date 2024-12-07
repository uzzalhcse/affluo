package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Payout struct {
	ent.Schema
}

func (Payout) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.Float("amount"),
		field.Time("paid_at"),
		field.Enum("status").Values("pending", "completed", "failed"),
		field.String("transaction_id").Optional(),
	}
}

func (Payout) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("payouts").Unique(),
	}
}
