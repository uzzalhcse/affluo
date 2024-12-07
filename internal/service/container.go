package service

import (
	"affluo/ent"
	"github.com/redis/go-redis/v9"
)

type Container struct {
	Tracking  *TrackingService
	Reporting *ReportingService
	User      *UserService
	Campaign  *CampaignService
}

func NewContainer(db *ent.Client, redis *redis.Client) *Container {
	return &Container{
		Tracking:  NewTrackingService(db, redis),
		Reporting: NewReportingService(db),
		User:      NewUserService(db),
		Campaign:  NewCampaignService(db),
	}
}
