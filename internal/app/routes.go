package app

func (a *Application) SetupRoutes() {
	tracking := a.App.Group("/track")
	tracking.Post("/", a.Handlers.Tracking.HandleTrack)
	tracking.Get("/script/:campaign_id", a.Handlers.Tracking.GenerateTrackingScript)

	campaigns := a.App.Group("/campaigns")
	campaigns.Post("/", a.Handlers.Campaign.CreateCampaign)
	campaigns.Get("/", a.Handlers.Campaign.ListCampaigns)
	campaigns.Get("/:id", a.Handlers.Campaign.GetCampaign)
	campaigns.Put("/:id", a.Handlers.Campaign.UpdateCampaign)
	campaigns.Delete("/:id", a.Handlers.Campaign.DeleteCampaign)

	users := a.App.Group("/users")
	users.Post("/", a.Handlers.User.CreateUser)
	users.Get("/", a.Handlers.User.ListUsers)
	users.Get("/:id", a.Handlers.User.GetUser)
	users.Put("/:id", a.Handlers.User.UpdateUser)
	users.Delete("/:id", a.Handlers.User.DeleteUser)
}
