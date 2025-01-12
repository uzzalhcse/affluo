package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type Affiliate struct {
	ent.Schema
}

func (Affiliate) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Positive(),
		field.String("tracking_code").Optional(),
		field.String("affiliate_user_id"),
		field.Enum("source").Values(
			"banner",
			"services",
			"products",
		).Default("banner"),
		field.Time("registration_date").Default(time.Now()),
		field.Time("first_transaction_date").Optional(),
		field.Float("commission"),
		field.Time("date").Default(time.Now),
	}
}

func (Affiliate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("affiliates").
			Unique().
			Required(), // Each affiliate must belong to a user
	}
}
