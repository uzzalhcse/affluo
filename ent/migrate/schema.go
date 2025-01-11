// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BannersColumns holds the columns for the "banners" table.
	BannersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"static", "dynamic"}, Default: "static"},
		{Name: "click_url", Type: field.TypeString, Nullable: true},
		{Name: "size", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"draft", "active", "inactive"}, Default: "draft"},
		{Name: "allowed_countries", Type: field.TypeJSON, Nullable: true},
		{Name: "weight", Type: field.TypeInt, Default: 1},
		{Name: "smart_weight", Type: field.TypeFloat64, Nullable: true},
		{Name: "last_impression", Type: field.TypeTime, Nullable: true},
		{Name: "start_date", Type: field.TypeTime, Nullable: true},
		{Name: "end_date", Type: field.TypeTime, Nullable: true},
		{Name: "allowed_devices", Type: field.TypeJSON, Nullable: true},
		{Name: "allowed_browsers", Type: field.TypeJSON, Nullable: true},
		{Name: "allowed_os", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// BannersTable holds the schema information for the "banners" table.
	BannersTable = &schema.Table{
		Name:       "banners",
		Columns:    BannersColumns,
		PrimaryKey: []*schema.Column{BannersColumns[0]},
	}
	// BannerCreativesColumns holds the columns for the "banner_creatives" table.
	BannerCreativesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "is_primary", Type: field.TypeBool, Default: false},
		{Name: "display_order", Type: field.TypeInt, Nullable: true},
		{Name: "banner_id", Type: field.TypeInt64},
		{Name: "creative_id", Type: field.TypeInt64},
	}
	// BannerCreativesTable holds the schema information for the "banner_creatives" table.
	BannerCreativesTable = &schema.Table{
		Name:       "banner_creatives",
		Columns:    BannerCreativesColumns,
		PrimaryKey: []*schema.Column{BannerCreativesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "banner_creatives_banners_banner",
				Columns:    []*schema.Column{BannerCreativesColumns[5]},
				RefColumns: []*schema.Column{BannersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "banner_creatives_creatives_creative",
				Columns:    []*schema.Column{BannerCreativesColumns[6]},
				RefColumns: []*schema.Column{CreativesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "bannercreative_banner_id_creative_id",
				Unique:  true,
				Columns: []*schema.Column{BannerCreativesColumns[5], BannerCreativesColumns[6]},
			},
		},
	}
	// BannerStatsColumns holds the columns for the "banner_stats" table.
	BannerStatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "date", Type: field.TypeTime},
		{Name: "impressions", Type: field.TypeInt64, Default: 0},
		{Name: "clicks", Type: field.TypeInt64, Default: 0},
		{Name: "leads", Type: field.TypeInt64, Default: 0},
		{Name: "earnings", Type: field.TypeFloat64, Default: 0},
		{Name: "ctr", Type: field.TypeFloat64, Nullable: true},
		{Name: "conversion_rate", Type: field.TypeFloat64, Nullable: true},
		{Name: "device_type", Type: field.TypeString, Nullable: true},
		{Name: "browser", Type: field.TypeString, Nullable: true},
		{Name: "os", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "banner_stats", Type: field.TypeInt64, Nullable: true},
		{Name: "user_stats", Type: field.TypeInt64, Nullable: true},
	}
	// BannerStatsTable holds the schema information for the "banner_stats" table.
	BannerStatsTable = &schema.Table{
		Name:       "banner_stats",
		Columns:    BannerStatsColumns,
		PrimaryKey: []*schema.Column{BannerStatsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "banner_stats_banners_stats",
				Columns:    []*schema.Column{BannerStatsColumns[13]},
				RefColumns: []*schema.Column{BannersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "banner_stats_users_stats",
				Columns:    []*schema.Column{BannerStatsColumns[14]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "bannerstats_date_banner_stats_user_stats",
				Unique:  true,
				Columns: []*schema.Column{BannerStatsColumns[1], BannerStatsColumns[13], BannerStatsColumns[14]},
			},
			{
				Name:    "bannerstats_banner_stats",
				Unique:  false,
				Columns: []*schema.Column{BannerStatsColumns[13]},
			},
			{
				Name:    "bannerstats_user_stats",
				Unique:  false,
				Columns: []*schema.Column{BannerStatsColumns[14]},
			},
		},
	}
	// CampaignsColumns holds the columns for the "campaigns" table.
	CampaignsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "unique_code", Type: field.TypeString, Unique: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"sale", "lead", "click", "subscription"}, Default: "sale"},
		{Name: "commission_type", Type: field.TypeEnum, Enums: []string{"flat_rate", "percentage", "tiered"}, Default: "percentage"},
		{Name: "base_commission_rate", Type: field.TypeFloat64, Default: 0},
		{Name: "commission_tiers", Type: field.TypeJSON, Nullable: true},
		{Name: "target_geography", Type: field.TypeString, Nullable: true},
		{Name: "target_demographics", Type: field.TypeJSON, Nullable: true},
		{Name: "start_date", Type: field.TypeTime},
		{Name: "end_date", Type: field.TypeTime, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"draft", "active", "paused", "completed"}, Default: "draft"},
		{Name: "tracking_url", Type: field.TypeString, Nullable: true},
		{Name: "total_clicks", Type: field.TypeInt, Default: 0},
		{Name: "total_conversions", Type: field.TypeInt, Default: 0},
		{Name: "total_revenue", Type: field.TypeFloat64, Default: 0},
		{Name: "conversion_rate", Type: field.TypeFloat64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_campaigns", Type: field.TypeInt64, Nullable: true},
	}
	// CampaignsTable holds the schema information for the "campaigns" table.
	CampaignsTable = &schema.Table{
		Name:       "campaigns",
		Columns:    CampaignsColumns,
		PrimaryKey: []*schema.Column{CampaignsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "campaigns_users_campaigns",
				Columns:    []*schema.Column{CampaignsColumns[20]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "campaign_name",
				Unique:  false,
				Columns: []*schema.Column{CampaignsColumns[1]},
			},
			{
				Name:    "campaign_unique_code",
				Unique:  false,
				Columns: []*schema.Column{CampaignsColumns[3]},
			},
			{
				Name:    "campaign_type",
				Unique:  false,
				Columns: []*schema.Column{CampaignsColumns[4]},
			},
			{
				Name:    "campaign_status",
				Unique:  false,
				Columns: []*schema.Column{CampaignsColumns[12]},
			},
		},
	}
	// CampaignLinksColumns holds the columns for the "campaign_links" table.
	CampaignLinksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "unique_code", Type: field.TypeString, Unique: true},
		{Name: "original_url", Type: field.TypeString},
		{Name: "tracking_url", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "campaign_links", Type: field.TypeInt64, Nullable: true},
	}
	// CampaignLinksTable holds the schema information for the "campaign_links" table.
	CampaignLinksTable = &schema.Table{
		Name:       "campaign_links",
		Columns:    CampaignLinksColumns,
		PrimaryKey: []*schema.Column{CampaignLinksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "campaign_links_campaigns_links",
				Columns:    []*schema.Column{CampaignLinksColumns[6]},
				RefColumns: []*schema.Column{CampaignsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CreativesColumns holds the columns for the "creatives" table.
	CreativesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "image_url", Type: field.TypeString, Nullable: true},
		{Name: "size", Type: field.TypeString, Nullable: true},
		{Name: "enabled", Type: field.TypeBool, Default: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// CreativesTable holds the schema information for the "creatives" table.
	CreativesTable = &schema.Table{
		Name:       "creatives",
		Columns:    CreativesColumns,
		PrimaryKey: []*schema.Column{CreativesColumns[0]},
	}
	// GigTrackingsColumns holds the columns for the "gig_trackings" table.
	GigTrackingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "date", Type: field.TypeTime},
		{Name: "type", Type: field.TypeString, Default: "services"},
		{Name: "utm_query", Type: field.TypeString, Nullable: true},
		{Name: "lp", Type: field.TypeString, Nullable: true},
		{Name: "track_id", Type: field.TypeString, Nullable: true},
		{Name: "revenue", Type: field.TypeFloat64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "gig_tracking_publisher", Type: field.TypeInt64},
	}
	// GigTrackingsTable holds the schema information for the "gig_trackings" table.
	GigTrackingsTable = &schema.Table{
		Name:       "gig_trackings",
		Columns:    GigTrackingsColumns,
		PrimaryKey: []*schema.Column{GigTrackingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "gig_trackings_users_publisher",
				Columns:    []*schema.Column{GigTrackingsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "gigtracking_date_lp_type_track_id_gig_tracking_publisher",
				Unique:  true,
				Columns: []*schema.Column{GigTrackingsColumns[1], GigTrackingsColumns[4], GigTrackingsColumns[2], GigTrackingsColumns[5], GigTrackingsColumns[9]},
			},
			{
				Name:    "gigtracking_gig_tracking_publisher",
				Unique:  false,
				Columns: []*schema.Column{GigTrackingsColumns[9]},
			},
		},
	}
	// LeadsColumns holds the columns for the "leads" table.
	LeadsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "reference_id", Type: field.TypeString, Nullable: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"product", "service"}},
		{Name: "amount", Type: field.TypeFloat64, Nullable: true},
		{Name: "currency", Type: field.TypeString, Default: "USD"},
		{Name: "ip_address", Type: field.TypeString, Nullable: true},
		{Name: "user_agent", Type: field.TypeString, Nullable: true},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "banner_leads", Type: field.TypeInt64, Nullable: true},
	}
	// LeadsTable holds the schema information for the "leads" table.
	LeadsTable = &schema.Table{
		Name:       "leads",
		Columns:    LeadsColumns,
		PrimaryKey: []*schema.Column{LeadsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "leads_banners_leads",
				Columns:    []*schema.Column{LeadsColumns[9]},
				RefColumns: []*schema.Column{BannersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PayoutsColumns holds the columns for the "payouts" table.
	PayoutsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "amount", Type: field.TypeFloat64},
		{Name: "paid_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"pending", "completed", "failed"}},
		{Name: "transaction_id", Type: field.TypeString, Nullable: true},
		{Name: "user_payouts", Type: field.TypeInt64, Nullable: true},
	}
	// PayoutsTable holds the schema information for the "payouts" table.
	PayoutsTable = &schema.Table{
		Name:       "payouts",
		Columns:    PayoutsColumns,
		PrimaryKey: []*schema.Column{PayoutsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "payouts_users_payouts",
				Columns:    []*schema.Column{PayoutsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "title", Type: field.TypeString},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_posts", Type: field.TypeInt64, Nullable: true},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_users_posts",
				Columns:    []*schema.Column{PostsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ReferralsColumns holds the columns for the "referrals" table.
	ReferralsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"pending", "approved", "rejected"}},
		{Name: "commission_amount", Type: field.TypeFloat64},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "processed_at", Type: field.TypeTime, Nullable: true},
		{Name: "campaign_referrals", Type: field.TypeInt64, Nullable: true},
		{Name: "user_referrals", Type: field.TypeInt64, Nullable: true},
	}
	// ReferralsTable holds the schema information for the "referrals" table.
	ReferralsTable = &schema.Table{
		Name:       "referrals",
		Columns:    ReferralsColumns,
		PrimaryKey: []*schema.Column{ReferralsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "referrals_campaigns_referrals",
				Columns:    []*schema.Column{ReferralsColumns[5]},
				RefColumns: []*schema.Column{CampaignsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "referrals_users_referrals",
				Columns:    []*schema.Column{ReferralsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TestsColumns holds the columns for the "tests" table.
	TestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// TestsTable holds the schema information for the "tests" table.
	TestsTable = &schema.Table{
		Name:       "tests",
		Columns:    TestsColumns,
		PrimaryKey: []*schema.Column{TestsColumns[0]},
	}
	// TracksColumns holds the columns for the "tracks" table.
	TracksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "ip_address", Type: field.TypeString},
		{Name: "user_agent", Type: field.TypeString},
		{Name: "device_fingerprint", Type: field.TypeString},
		{Name: "referrer", Type: field.TypeString, Nullable: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"click", "impression", "conversion"}},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"valid", "suspected_fraud", "blacklisted"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "is_unique_click", Type: field.TypeBool, Default: false},
		{Name: "additional_metadata", Type: field.TypeJSON, Nullable: true},
		{Name: "campaign_tracks", Type: field.TypeInt64, Nullable: true},
		{Name: "campaign_link_tracks", Type: field.TypeInt64, Nullable: true},
		{Name: "user_tracks", Type: field.TypeInt64, Nullable: true},
	}
	// TracksTable holds the schema information for the "tracks" table.
	TracksTable = &schema.Table{
		Name:       "tracks",
		Columns:    TracksColumns,
		PrimaryKey: []*schema.Column{TracksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tracks_campaigns_tracks",
				Columns:    []*schema.Column{TracksColumns[10]},
				RefColumns: []*schema.Column{CampaignsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "tracks_campaign_links_tracks",
				Columns:    []*schema.Column{TracksColumns[11]},
				RefColumns: []*schema.Column{CampaignLinksColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "tracks_users_tracks",
				Columns:    []*schema.Column{TracksColumns[12]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password_hash", Type: field.TypeString},
		{Name: "first_name", Type: field.TypeString, Nullable: true},
		{Name: "last_name", Type: field.TypeString, Nullable: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"admin", "affiliate", "manager"}, Default: "affiliate"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "last_login", Type: field.TypeTime, Nullable: true},
		{Name: "reset_token", Type: field.TypeString, Nullable: true},
		{Name: "reset_token_expires_at", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_email",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[2]},
			},
		},
	}
	// CampaignBannersColumns holds the columns for the "campaign_banners" table.
	CampaignBannersColumns = []*schema.Column{
		{Name: "campaign_id", Type: field.TypeInt64},
		{Name: "banner_id", Type: field.TypeInt64},
	}
	// CampaignBannersTable holds the schema information for the "campaign_banners" table.
	CampaignBannersTable = &schema.Table{
		Name:       "campaign_banners",
		Columns:    CampaignBannersColumns,
		PrimaryKey: []*schema.Column{CampaignBannersColumns[0], CampaignBannersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "campaign_banners_campaign_id",
				Columns:    []*schema.Column{CampaignBannersColumns[0]},
				RefColumns: []*schema.Column{CampaignsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "campaign_banners_banner_id",
				Columns:    []*schema.Column{CampaignBannersColumns[1]},
				RefColumns: []*schema.Column{BannersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BannersTable,
		BannerCreativesTable,
		BannerStatsTable,
		CampaignsTable,
		CampaignLinksTable,
		CreativesTable,
		GigTrackingsTable,
		LeadsTable,
		PayoutsTable,
		PostsTable,
		ReferralsTable,
		TestsTable,
		TracksTable,
		UsersTable,
		CampaignBannersTable,
	}
)

func init() {
	BannerCreativesTable.ForeignKeys[0].RefTable = BannersTable
	BannerCreativesTable.ForeignKeys[1].RefTable = CreativesTable
	BannerStatsTable.ForeignKeys[0].RefTable = BannersTable
	BannerStatsTable.ForeignKeys[1].RefTable = UsersTable
	CampaignsTable.ForeignKeys[0].RefTable = UsersTable
	CampaignLinksTable.ForeignKeys[0].RefTable = CampaignsTable
	GigTrackingsTable.ForeignKeys[0].RefTable = UsersTable
	LeadsTable.ForeignKeys[0].RefTable = BannersTable
	PayoutsTable.ForeignKeys[0].RefTable = UsersTable
	PostsTable.ForeignKeys[0].RefTable = UsersTable
	ReferralsTable.ForeignKeys[0].RefTable = CampaignsTable
	ReferralsTable.ForeignKeys[1].RefTable = UsersTable
	TracksTable.ForeignKeys[0].RefTable = CampaignsTable
	TracksTable.ForeignKeys[1].RefTable = CampaignLinksTable
	TracksTable.ForeignKeys[2].RefTable = UsersTable
	CampaignBannersTable.ForeignKeys[0].RefTable = CampaignsTable
	CampaignBannersTable.ForeignKeys[1].RefTable = BannersTable
}
