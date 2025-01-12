package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"fmt"
	"regexp"
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
		field.String("email").Unique().
			Validate(func(s string) error {
				// Basic email validation
				emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
				if !emailRegex.MatchString(s) {
					return fmt.Errorf("invalid email format")
				}
				return nil
			}),
		field.String("password_hash").Sensitive(),
		field.String("first_name").Optional(),
		field.String("last_name").Optional(),
		field.Enum("role").Values("admin", "affiliate", "manager").Default("affiliate"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Bool("is_active").Default(true),
		field.Time("last_login").Optional(),
		field.String("reset_token").Optional(),
		field.Time("reset_token_expires_at").Optional(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("campaigns", Campaign.Type),
		edge.To("referrals", Referral.Type),
		edge.To("tracks", Track.Type),
		edge.To("payouts", Payout.Type), // Add this line
		edge.To("posts", Post.Type),     // Add this line
		edge.To("stats", BannerStats.Type),
		edge.From("gig_trackings", GigTracking.Type).
			Ref("publisher"),
		edge.From("commission_plan", CommissionPlan.Type).
			Ref("publishers").
			Unique(),
	}
}
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email"),
	}
}
