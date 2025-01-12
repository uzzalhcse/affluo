// internal/service/container.go
package service

import (
	"affluo/ent"

	"github.com/redis/go-redis/v9"
)

type Container struct {
	User       *UserService
	Campaign   *CampaignService
	Tracking   *TrackingService
	Auth       *AuthService // Add Auth Service
	Banner     *BannerService
	AntiFraud  *AntiFraudService
	Commission *CommissionService
}

func NewContainer(client *ent.Client, redisClient *redis.Client) *Container {
	// Add JWT secret key configuration - ideally from environment or config
	jwtSecretKey := "your-secret-key-here" // In real app, use a secure method to generate and store this

	return &Container{
		User:       NewUserService(client),
		Campaign:   NewCampaignService(client),
		Tracking:   NewTrackingService(client),
		Auth:       NewAuthService(client, redisClient, jwtSecretKey), // Initialize Auth Service
		Banner:     NewBannerService(client),
		AntiFraud:  NewAntiFraudService(client, redisClient),
		Commission: NewCommissionService(client),
	}
}
