// internal/service/container.go
package service

import (
	"affluo/ent"

	"github.com/redis/go-redis/v9"
)

type Container struct {
	User      *UserService
	Campaign  *CampaignService
	Tracking  *TrackingService
	Reporting *ReportingService
	Post      *PostService
	Auth      *AuthService // Add Auth Service
}

func NewContainer(client *ent.Client, redisClient *redis.Client) *Container {
	// Add JWT secret key configuration - ideally from environment or config
	jwtSecretKey := "your-secret-key-here" // In real app, use a secure method to generate and store this

	return &Container{
		User:      NewUserService(client),
		Campaign:  NewCampaignService(client),
		Tracking:  NewTrackingService(client, redisClient),
		Reporting: NewReportingService(client),
		Post:      NewPostService(client),
		Auth:      NewAuthService(client, redisClient, jwtSecretKey), // Initialize Auth Service
	}
}
