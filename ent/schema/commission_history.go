package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type CommissionHistory struct {
	ent.Schema
}

func (CommissionHistory) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Positive(),
		field.Float("amount"),
		field.String("affiliate_user_id"),
		field.String("trx_id").Optional(),
		field.String("track_id").Optional(),
		field.Float("commission_rate"),
		field.Bool("is_first_order"),
		field.String("date").Default(time.Now().Format("2006-01-02")),
	}
}

func (CommissionHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("commission_histories").
			Unique().
			Required(),
	}
}
