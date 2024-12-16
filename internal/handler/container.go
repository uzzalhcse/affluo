// internal/handler/container.go
package handler

import (
	"affluo/internal/service"
)

type Container struct {
	User     *UserHandler
	Campaign *CampaignHandler
	Tracking *TrackingHandler
	Post     *PostHandler
	Auth     *AuthHandler // Add Auth Handler
	Banner   *BannerHandler
}

func NewContainer(services *service.Container) *Container {
	return &Container{
		User:     NewUserHandler(services.User),
		Campaign: NewCampaignHandler(services.Campaign),
		Tracking: NewTrackingHandler(services.Tracking),
		//Post:       NewPostHandler(services.Post),
		Auth:   NewAuthHandler(services.Auth), // Initialize Auth Handler
		Banner: NewBannerHandler(services.Banner),
	}
}
