package handler

import "affluo/internal/service"

type Container struct {
	Tracking *TrackingHandler
	User     *UserHandler
	Campaign *CampaignHandler
}

func NewContainer(services *service.Container) *Container {
	return &Container{
		Tracking: NewTrackingHandler(services.Tracking),
		User:     NewUserHandler(services.User),
		Campaign: NewCampaignHandler(services.Campaign),
	}
}
