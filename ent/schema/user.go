package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// User schema represents affiliate users
type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Positive(),
		field.String("username").Unique(),
		field.String("email").Unique(),
		field.String("password_hash"),
		field.String("first_name").Optional(),
		field.String("last_name").Optional(),
		field.Enum("role").Values("admin", "affiliate", "manager"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Bool("is_active").Default(true),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("campaigns", Campaign.Type),
		edge.To("referrals", Referral.Type),
		edge.To("tracks", Track.Type),
		edge.To("payouts", Payout.Type), // Add this line
		edge.To("posts", Post.Type),     // Add this line
	}
}
