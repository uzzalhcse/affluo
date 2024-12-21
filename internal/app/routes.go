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

	users := a.App.Group("/api/users")
	users.Post("/", a.Handlers.User.CreateUser)
	users.Get("/", a.Handlers.User.ListUsers)
	users.Get("/:id", a.Handlers.User.GetUser)
	users.Put("/:id", a.Handlers.User.UpdateUser)
	users.Delete("/:id", a.Handlers.User.DeleteUser)

	// New Authentication Routes
	auth := a.App.Group("/api/auth")
	auth.Post("/register", a.Handlers.Auth.Register)
	auth.Post("/login", a.Handlers.Auth.Login)
	auth.Post("/reset-password", a.Handlers.Auth.InitiatePasswordReset)
	auth.Post("/confirm-reset", a.Handlers.Auth.ResetPassword)
	protectedAuth := auth.Use(a.Middleware.Auth.Authenticate())
	protectedAuth.Get("/profile", a.Handlers.Auth.Profile)

	// Banner routes
	bannerGroup := a.App.Group("/api/banners")
	{
		bannerGroup.Post("/", a.Middleware.Auth.Authenticate(), a.Handlers.Banner.CreateBanner)
		bannerGroup.Get("/", a.Middleware.Auth.Authenticate(), a.Handlers.Banner.GetAllBanners)
		bannerGroup.Get("/:id", a.Middleware.Auth.Authenticate(), a.Handlers.Banner.GetBanner)
		bannerGroup.Post("/:id/creatives", a.Middleware.Auth.Authenticate(), a.Handlers.Banner.CreateBannerCreative)
		bannerGroup.Get("/:id/creatives", a.Middleware.Auth.Authenticate(), a.Handlers.Banner.GetBannerCreatives)
	}
	// Creative routes
	creativeGroup := a.App.Group("/api/creatives")
	{
		creativeGroup.Get("/", a.Middleware.Auth.Authenticate(), a.Handlers.Banner.GetAllBannerCreatives)
		creativeGroup.Post("/", a.Middleware.Auth.Authenticate(), a.Handlers.Banner.CreateBannerCreative)
		creativeGroup.Get("/:id/creatives", a.Middleware.Auth.Authenticate(), a.Handlers.Banner.GetBannerCreatives)
	}

	// Publisher routes
	publisherGroup := a.App.Group("/api/publisher")
	{
		publisherGroup.Get("/banners", a.Middleware.Auth.Authenticate(), a.Handlers.Banner.GetPublisherBanners)
	}
	trackingGroup := a.App.Group("/api/tracking")
	{
		trackingGroup.Use(middleware.AntiFraud(a.Services.AntiFraud))
		trackingGroup.Post("/impression/:id", a.Handlers.Tracking.RecordImpression)
		trackingGroup.Post("/click/:id", a.Handlers.Tracking.RecordClick)
		trackingGroup.Post("/lead/:id", a.Handlers.Tracking.RecordLead)

		trackingGroup.Get("/pixel/:id", a.Handlers.Tracking.PixelTracking)

		// Banner selection endpoint
		trackingGroup.Get("/select/:campaign_id", a.Handlers.Tracking.SelectBanner)
	}

	// Stats endpoints (require authentication)
	statsGroup := a.App.Group("/api/stats")
	{
		statsGroup.Get("/banner/:id", a.Middleware.Auth.Authenticate(), a.Handlers.Tracking.GetStats)
	}

}
