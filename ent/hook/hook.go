// Code generated by ent, DO NOT EDIT.

package hook

import (
	"affluo/ent"
	"context"
	"fmt"
)

// The BannerFunc type is an adapter to allow the use of ordinary
// function as Banner mutator.
type BannerFunc func(context.Context, *ent.BannerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BannerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.BannerMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BannerMutation", m)
}

// The BannerCreativeFunc type is an adapter to allow the use of ordinary
// function as BannerCreative mutator.
type BannerCreativeFunc func(context.Context, *ent.BannerCreativeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BannerCreativeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.BannerCreativeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BannerCreativeMutation", m)
}

// The BannerStatsFunc type is an adapter to allow the use of ordinary
// function as BannerStats mutator.
type BannerStatsFunc func(context.Context, *ent.BannerStatsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BannerStatsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.BannerStatsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BannerStatsMutation", m)
}

// The CampaignFunc type is an adapter to allow the use of ordinary
// function as Campaign mutator.
type CampaignFunc func(context.Context, *ent.CampaignMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CampaignFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CampaignMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CampaignMutation", m)
}

// The CampaignLinkFunc type is an adapter to allow the use of ordinary
// function as CampaignLink mutator.
type CampaignLinkFunc func(context.Context, *ent.CampaignLinkMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CampaignLinkFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CampaignLinkMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CampaignLinkMutation", m)
}

// The CommissionPlanFunc type is an adapter to allow the use of ordinary
// function as CommissionPlan mutator.
type CommissionPlanFunc func(context.Context, *ent.CommissionPlanMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CommissionPlanFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CommissionPlanMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CommissionPlanMutation", m)
}

// The CreativeFunc type is an adapter to allow the use of ordinary
// function as Creative mutator.
type CreativeFunc func(context.Context, *ent.CreativeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CreativeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CreativeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CreativeMutation", m)
}

// The GigTrackingFunc type is an adapter to allow the use of ordinary
// function as GigTracking mutator.
type GigTrackingFunc func(context.Context, *ent.GigTrackingMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GigTrackingFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.GigTrackingMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GigTrackingMutation", m)
}

// The LeadFunc type is an adapter to allow the use of ordinary
// function as Lead mutator.
type LeadFunc func(context.Context, *ent.LeadMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f LeadFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.LeadMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.LeadMutation", m)
}

// The PayoutFunc type is an adapter to allow the use of ordinary
// function as Payout mutator.
type PayoutFunc func(context.Context, *ent.PayoutMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PayoutFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PayoutMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PayoutMutation", m)
}

// The PostFunc type is an adapter to allow the use of ordinary
// function as Post mutator.
type PostFunc func(context.Context, *ent.PostMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PostFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PostMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PostMutation", m)
}

// The ReferralFunc type is an adapter to allow the use of ordinary
// function as Referral mutator.
type ReferralFunc func(context.Context, *ent.ReferralMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ReferralFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ReferralMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ReferralMutation", m)
}

// The TestFunc type is an adapter to allow the use of ordinary
// function as Test mutator.
type TestFunc func(context.Context, *ent.TestMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TestFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TestMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TestMutation", m)
}

// The TrackFunc type is an adapter to allow the use of ordinary
// function as Track mutator.
type TrackFunc func(context.Context, *ent.TrackMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TrackFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TrackMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TrackMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
