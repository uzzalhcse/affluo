// internal/app/routes.go
package app

import "affluo/internal/middleware"

func (a *Application) SetupRoutes() {

	campaigns := a.App.Group("/api/campaigns")
	campaigns.Use(a.Middleware.Auth.Authenticate())
	// Standard CRUD Operations
	campaigns.Post("/", a.Handlers.Campaign.CreateCampaign)      // Create a new campaign
	campaigns.Get("/", a.Handlers.Campaign.ListCampaigns)        // List all campaigns with optional filtering
	campaigns.Get("/:id", a.Handlers.Campaign.GetCampaign)       // Retrieve a specific campaign
	campaigns.Put("/:id", a.Handlers.Campaign.UpdateCampaign)    // Update an existing campaign
	campaigns.Delete("/:id", a.Handlers.Campaign.DeleteCampaign) // Soft delete a campaign

	// Additional Performance and Tracking Endpoints
	campaigns.Get("/:id/performance", a.Handlers.Campaign.GetCampaignPerformance)  // Get detailed performance metrics
	campaigns.Post("/:id/tracking-link", a.Handlers.Campaign.GenerateTrackingLink) // Generate a unique tracking link

	// New Authentication Routes
	auth := a.App.Group("/api/auth")
	auth.Post("/register", a.Handlers.Auth.Register)
	auth.Post("/login", a.Handlers.Auth.Login)
	auth.Post("/reset-password", a.Handlers.Auth.InitiatePasswordReset)
	auth.Post("/confirm-reset", a.Handlers.Auth.ResetPassword)
	protectedAuth := auth.Use(a.Middleware.Auth.Authenticate())
	protectedAuth.Get("/profile", a.Handlers.Auth.Profile)

	// Admin Api

	adminGroup := a.App.Group("/api/admin")
	{
		// Banner routes
		bannerGroup := adminGroup.Group("/banners")
		{
			bannerGroup.Post("/", a.Middleware.Auth.Authenticate("admin"), a.Handlers.Banner.CreateBanner)
			bannerGroup.Get("/", a.Middleware.Auth.Authenticate("admin"), a.Handlers.Banner.GetAllBanners)
			bannerGroup.Get("/:id", a.Middleware.Auth.Authenticate("admin"), a.Handlers.Banner.GetBanner)
			bannerGroup.Put("/:id", a.Middleware.Auth.Authenticate("admin"), a.Handlers.Banner.UpdateBanner)
		}

		// Creative routes
		creativeGroup := adminGroup.Group("/creatives")
		{
			creativeGroup.Get("/", a.Middleware.Auth.Authenticate("admin"), a.Handlers.Banner.GetAllCreatives)
			creativeGroup.Post("/", a.Middleware.Auth.Authenticate("admin"), a.Handlers.Banner.CreateBannerCreative)
		}

		users := adminGroup.Group("/users")
		users.Get("/", a.Handlers.User.ListUsers)
		users.Get("/toggle-active/:id", a.Handlers.User.ToggleActive)
		users.Post("/", a.Handlers.User.CreateUser)
		users.Get("/:id", a.Handlers.User.GetUser)
		users.Put("/:id", a.Handlers.User.UpdateUser)
		users.Delete("/:id", a.Handlers.User.DeleteUser)

		commissionGroup := adminGroup.Group("/commission-plans")
		{
			commissionGroup.Post("/", a.Middleware.Auth.Authenticate("admin"), a.Handlers.Commission.CreateCommissionPlan)
			commissionGroup.Get("/", a.Middleware.Auth.Authenticate("admin"), a.Handlers.Commission.ListCommissionPlans)
			commissionGroup.Get("/:id", a.Middleware.Auth.Authenticate("admin"), a.Handlers.Commission.GetCommissionPlan)
			commissionGroup.Put("/:id", a.Middleware.Auth.Authenticate("admin"), a.Handlers.Commission.UpdateCommissionPlan)
			commissionGroup.Post("/:id/assign/:publisherId", a.Middleware.Auth.Authenticate("admin"), a.Handlers.Commission.AssignToPublisher)
			commissionGroup.Get("/publisher/:publisherId", a.Middleware.Auth.Authenticate("admin"), a.Handlers.Commission.GetPublisherCommission)
		}
	}

	// Publisher routes remain the same
	publisherGroup := a.App.Group("/api/publisher")
	{
		publisherGroup.Get("/banners", a.Middleware.Auth.Authenticate(), a.Handlers.Banner.GetPublisherBanners)
		publisherGroup.Post("/affiliates", a.Middleware.Auth.Authenticate(), a.Handlers.Affiliate.CreateAffiliate)
		publisherGroup.Get("/affiliates", a.Middleware.Auth.Authenticate(), a.Handlers.Affiliate.GetAffiliate)
	}
	trackingGroup := a.App.Group("/api/tracking")
	{
		trackingGroup.Use(middleware.AntiFraud(a.Services.AntiFraud))
		trackingGroup.Post("/impression/:id", a.Handlers.Tracking.RecordImpression)
		trackingGroup.Get("/click", a.Handlers.Tracking.RecordClick)
		trackingGroup.Post("/lead/:id", a.Handlers.Tracking.RecordLead)

		trackingGroup.Get("/pixel/:id", a.Handlers.Tracking.PixelTracking)
		trackingGroup.Get("/visit", a.Handlers.Tracking.VisitTracking)

		// Banner selection endpoint
		trackingGroup.Get("/select/:campaign_id", a.Handlers.Tracking.SelectBanner)
	}

	// Stats endpoints (require authentication)
	statsGroup := a.App.Group("/api/stats")
	{
		statsGroup.Get("/banner", a.Middleware.Auth.Authenticate(), a.Handlers.Tracking.GetStats)
		statsGroup.Get("/gigs", a.Middleware.Auth.Authenticate(), a.Handlers.Tracking.GetGigReports)
	}
	webhookGroup := a.App.Group("/api")
	{
		webhookGroup.Get("/callback", a.Middleware.Auth.Authenticate(), a.Handlers.Tracking.GigLeadCallBack)
	}

}
