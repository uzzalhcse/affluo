// ent/schema/commissionplan.go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CommissionPlan holds the schema definition for the CommissionPlan entity
type CommissionPlan struct {
	ent.Schema
}

// Fields of the CommissionPlan
func (CommissionPlan) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("description").
			Optional(),
		field.Enum("type").
			Values("fixed", "percentage"),
		field.Float("click_commission").
			Default(0),
		field.Float("impression_commission").
			Default(0),
		field.Float("lead_commission").
			Default(0),
		field.Float("minimum_payout").
			Default(0),
		field.Time("valid_from").
			Optional(),
		field.Time("valid_until").
			Optional(),
		field.Bool("is_active").
			Default(true),
		field.Bool("is_default").
			Default(false),
	}
}

// Edges of the CommissionPlan
func (CommissionPlan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("publishers", User.Type),
	}
}
