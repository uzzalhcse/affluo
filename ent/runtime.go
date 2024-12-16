// Code generated by ent, DO NOT EDIT.

package ent

import (
	"affluo/ent/banner"
	"affluo/ent/bannercreative"
	"affluo/ent/campaign"
	"affluo/ent/campaignlink"
	"affluo/ent/referral"
	"affluo/ent/schema"
	"affluo/ent/track"
	"affluo/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	bannerFields := schema.Banner{}.Fields()
	_ = bannerFields
	// bannerDescName is the schema descriptor for name field.
	bannerDescName := bannerFields[1].Descriptor()
	// banner.NameValidator is a validator for the "name" field. It is called by the builders before save.
	banner.NameValidator = bannerDescName.Validators[0].(func(string) error)
	// bannerDescCreatedAt is the schema descriptor for created_at field.
	bannerDescCreatedAt := bannerFields[8].Descriptor()
	// banner.DefaultCreatedAt holds the default value on creation for the created_at field.
	banner.DefaultCreatedAt = bannerDescCreatedAt.Default.(func() time.Time)
	// bannerDescUpdatedAt is the schema descriptor for updated_at field.
	bannerDescUpdatedAt := bannerFields[9].Descriptor()
	// banner.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	banner.DefaultUpdatedAt = bannerDescUpdatedAt.Default.(func() time.Time)
	// banner.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	banner.UpdateDefaultUpdatedAt = bannerDescUpdatedAt.UpdateDefault.(func() time.Time)
	// bannerDescID is the schema descriptor for id field.
	bannerDescID := bannerFields[0].Descriptor()
	// banner.IDValidator is a validator for the "id" field. It is called by the builders before save.
	banner.IDValidator = bannerDescID.Validators[0].(func(int64) error)
	bannercreativeFields := schema.BannerCreative{}.Fields()
	_ = bannercreativeFields
	// bannercreativeDescEnabled is the schema descriptor for enabled field.
	bannercreativeDescEnabled := bannercreativeFields[4].Descriptor()
	// bannercreative.DefaultEnabled holds the default value on creation for the enabled field.
	bannercreative.DefaultEnabled = bannercreativeDescEnabled.Default.(bool)
	// bannercreativeDescCreatedAt is the schema descriptor for created_at field.
	bannercreativeDescCreatedAt := bannercreativeFields[5].Descriptor()
	// bannercreative.DefaultCreatedAt holds the default value on creation for the created_at field.
	bannercreative.DefaultCreatedAt = bannercreativeDescCreatedAt.Default.(func() time.Time)
	// bannercreativeDescUpdatedAt is the schema descriptor for updated_at field.
	bannercreativeDescUpdatedAt := bannercreativeFields[6].Descriptor()
	// bannercreative.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	bannercreative.DefaultUpdatedAt = bannercreativeDescUpdatedAt.Default.(func() time.Time)
	// bannercreative.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	bannercreative.UpdateDefaultUpdatedAt = bannercreativeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// bannercreativeDescID is the schema descriptor for id field.
	bannercreativeDescID := bannercreativeFields[0].Descriptor()
	// bannercreative.IDValidator is a validator for the "id" field. It is called by the builders before save.
	bannercreative.IDValidator = bannercreativeDescID.Validators[0].(func(int64) error)
	campaignFields := schema.Campaign{}.Fields()
	_ = campaignFields
	// campaignDescName is the schema descriptor for name field.
	campaignDescName := campaignFields[1].Descriptor()
	// campaign.NameValidator is a validator for the "name" field. It is called by the builders before save.
	campaign.NameValidator = campaignDescName.Validators[0].(func(string) error)
	// campaignDescBaseCommissionRate is the schema descriptor for base_commission_rate field.
	campaignDescBaseCommissionRate := campaignFields[6].Descriptor()
	// campaign.DefaultBaseCommissionRate holds the default value on creation for the base_commission_rate field.
	campaign.DefaultBaseCommissionRate = campaignDescBaseCommissionRate.Default.(float64)
	// campaignDescTotalClicks is the schema descriptor for total_clicks field.
	campaignDescTotalClicks := campaignFields[14].Descriptor()
	// campaign.DefaultTotalClicks holds the default value on creation for the total_clicks field.
	campaign.DefaultTotalClicks = campaignDescTotalClicks.Default.(int)
	// campaignDescTotalConversions is the schema descriptor for total_conversions field.
	campaignDescTotalConversions := campaignFields[15].Descriptor()
	// campaign.DefaultTotalConversions holds the default value on creation for the total_conversions field.
	campaign.DefaultTotalConversions = campaignDescTotalConversions.Default.(int)
	// campaignDescTotalRevenue is the schema descriptor for total_revenue field.
	campaignDescTotalRevenue := campaignFields[16].Descriptor()
	// campaign.DefaultTotalRevenue holds the default value on creation for the total_revenue field.
	campaign.DefaultTotalRevenue = campaignDescTotalRevenue.Default.(float64)
	// campaignDescConversionRate is the schema descriptor for conversion_rate field.
	campaignDescConversionRate := campaignFields[17].Descriptor()
	// campaign.DefaultConversionRate holds the default value on creation for the conversion_rate field.
	campaign.DefaultConversionRate = campaignDescConversionRate.Default.(float64)
	// campaignDescCreatedAt is the schema descriptor for created_at field.
	campaignDescCreatedAt := campaignFields[18].Descriptor()
	// campaign.DefaultCreatedAt holds the default value on creation for the created_at field.
	campaign.DefaultCreatedAt = campaignDescCreatedAt.Default.(func() time.Time)
	// campaignDescUpdatedAt is the schema descriptor for updated_at field.
	campaignDescUpdatedAt := campaignFields[19].Descriptor()
	// campaign.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	campaign.DefaultUpdatedAt = campaignDescUpdatedAt.Default.(func() time.Time)
	// campaign.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	campaign.UpdateDefaultUpdatedAt = campaignDescUpdatedAt.UpdateDefault.(func() time.Time)
	// campaignDescID is the schema descriptor for id field.
	campaignDescID := campaignFields[0].Descriptor()
	// campaign.IDValidator is a validator for the "id" field. It is called by the builders before save.
	campaign.IDValidator = campaignDescID.Validators[0].(func(int64) error)
	campaignlinkFields := schema.CampaignLink{}.Fields()
	_ = campaignlinkFields
	// campaignlinkDescCreatedAt is the schema descriptor for created_at field.
	campaignlinkDescCreatedAt := campaignlinkFields[4].Descriptor()
	// campaignlink.DefaultCreatedAt holds the default value on creation for the created_at field.
	campaignlink.DefaultCreatedAt = campaignlinkDescCreatedAt.Default.(func() time.Time)
	// campaignlinkDescIsActive is the schema descriptor for is_active field.
	campaignlinkDescIsActive := campaignlinkFields[5].Descriptor()
	// campaignlink.DefaultIsActive holds the default value on creation for the is_active field.
	campaignlink.DefaultIsActive = campaignlinkDescIsActive.Default.(bool)
	// campaignlinkDescID is the schema descriptor for id field.
	campaignlinkDescID := campaignlinkFields[0].Descriptor()
	// campaignlink.IDValidator is a validator for the "id" field. It is called by the builders before save.
	campaignlink.IDValidator = campaignlinkDescID.Validators[0].(func(int64) error)
	referralFields := schema.Referral{}.Fields()
	_ = referralFields
	// referralDescCreatedAt is the schema descriptor for created_at field.
	referralDescCreatedAt := referralFields[3].Descriptor()
	// referral.DefaultCreatedAt holds the default value on creation for the created_at field.
	referral.DefaultCreatedAt = referralDescCreatedAt.Default.(func() time.Time)
	// referralDescID is the schema descriptor for id field.
	referralDescID := referralFields[0].Descriptor()
	// referral.IDValidator is a validator for the "id" field. It is called by the builders before save.
	referral.IDValidator = referralDescID.Validators[0].(func(int64) error)
	trackFields := schema.Track{}.Fields()
	_ = trackFields
	// trackDescCreatedAt is the schema descriptor for created_at field.
	trackDescCreatedAt := trackFields[7].Descriptor()
	// track.DefaultCreatedAt holds the default value on creation for the created_at field.
	track.DefaultCreatedAt = trackDescCreatedAt.Default.(func() time.Time)
	// trackDescIsUniqueClick is the schema descriptor for is_unique_click field.
	trackDescIsUniqueClick := trackFields[8].Descriptor()
	// track.DefaultIsUniqueClick holds the default value on creation for the is_unique_click field.
	track.DefaultIsUniqueClick = trackDescIsUniqueClick.Default.(bool)
	// trackDescID is the schema descriptor for id field.
	trackDescID := trackFields[0].Descriptor()
	// track.IDValidator is a validator for the "id" field. It is called by the builders before save.
	track.IDValidator = trackDescID.Validators[0].(func(int64) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[7].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[8].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescIsActive is the schema descriptor for is_active field.
	userDescIsActive := userFields[9].Descriptor()
	// user.DefaultIsActive holds the default value on creation for the is_active field.
	user.DefaultIsActive = userDescIsActive.Default.(bool)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(int64) error)
}
