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
		field.Float("first_lead_commission").
			Default(25), // 25% for first order
		field.Float("repeat_lead_commission").
			Default(10), // 10% for subsequent orders
		field.Int("valid_months").
			Default(12), // How many months the repeat commission is valid
		field.Float("minimum_payout").
			Default(0),
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
