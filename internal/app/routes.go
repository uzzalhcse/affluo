// internal/app/routes.go
package app

func (a *Application) SetupRoutes() {
	// Existing routes...
	tracking := a.App.Group("/api/track")
	tracking.Post("/", a.Handlers.Tracking.HandleTrack)
	tracking.Get("/script/:campaign_id", a.Handlers.Tracking.GenerateTrackingScript)

	campaigns := a.App.Group("/api/campaigns")
	campaigns.Post("/", a.Handlers.Campaign.CreateCampaign)
	campaigns.Get("/", a.Handlers.Campaign.ListCampaigns)
	campaigns.Get("/:id", a.Handlers.Campaign.GetCampaign)
	campaigns.Put("/:id", a.Handlers.Campaign.UpdateCampaign)
	campaigns.Delete("/:id", a.Handlers.Campaign.DeleteCampaign)

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

	// Example of protected routes with role-based access
	//protectedCampaigns := a.App.Group("/api/campaigns")
	//protectedCampaigns.Use(a.Middleware.Auth.Authenticate("admin", "manager"))
	//protectedCampaigns.Post("/", a.Handlers.Campaign.CreateCampaign)
	//protectedCampaigns.Delete("/:id", a.Handlers.Campaign.DeleteCampaign)
}
