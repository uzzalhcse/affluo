// internal/handler/container.go
package handler

import (
	"affluo/internal/service"
)

type Container struct {
	User       *UserHandler
	Campaign   *CampaignHandler
	Tracking   *TrackingHandler
	Auth       *AuthHandler // Add Auth Handler
	Banner     *BannerHandler
	Commission *CommissionHandler
	Affiliate  *AffiliateHandler
}

func NewContainer(services *service.Container) *Container {
	return &Container{
		User:       NewUserHandler(services.User),
		Campaign:   NewCampaignHandler(services.Campaign),
		Tracking:   NewTrackingHandler(services.Tracking, services.Affiliate),
		Auth:       NewAuthHandler(services.Auth), // Initialize Auth Handler
		Banner:     NewBannerHandler(services.Banner),
		Commission: NewCommissionHandler(services.Commission),
		Affiliate:  NewAffiliateHandler(services.Affiliate),
	}
}
