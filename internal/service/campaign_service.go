package service

import (
	"context"
	"errors"
	"fmt"

	"affluo/ent"
	"affluo/ent/campaign"
	"affluo/internal/dto"
	"github.com/google/uuid"
)

type CampaignService struct {
	client *ent.Client
}

func NewCampaignService(client *ent.Client) *CampaignService {
	return &CampaignService{
		client: client,
	}
}

// CreateCampaign creates a new marketing campaign
func (s *CampaignService) CreateCampaign(ctx context.Context, req *dto.CreateCampaignRequest) (*ent.Campaign, error) {
	// Validate input
	if err := s.validateCampaignRequest(req); err != nil {
		return nil, err
	}

	// Generate unique tracking code
	uniqueCode := s.generateCampaignCode()

	// Create campaign
	return s.client.Campaign.Create().
		SetName(req.Name).
		SetDescription(req.Description).
		SetType(campaign.Type(req.Type)).
		SetPayoutRate(req.PayoutRate).
		SetStartDate(req.StartDate).
		SetEndDate(req.EndDate).
		SetStatus("active").
		SetTrackingURL(s.generateTrackingURL(uniqueCode)).
		SetUniqueCode(uniqueCode).
		Save(ctx)
}

// UpdateCampaign updates an existing campaign
func (s *CampaignService) UpdateCampaign(ctx context.Context, id int64, req *dto.UpdateCampaignRequest) (*ent.Campaign, error) {
	// Fetch existing campaign
	_, err := s.client.Campaign.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update campaign
	update := s.client.Campaign.UpdateOneID(id)

	if req.Name != "" {
		update.SetName(req.Name)
	}
	if req.Description != "" {
		update.SetDescription(req.Description)
	}
	if req.Type != "" {
		update.SetType(campaign.Type(req.Type))
	}
	if req.PayoutRate > 0 {
		update.SetPayoutRate(req.PayoutRate)
	}
	if !req.StartDate.IsZero() {
		update.SetStartDate(req.StartDate)
	}
	if !req.EndDate.IsZero() {
		update.SetEndDate(req.EndDate)
	}
	if req.Status != "" {
		update.SetStatus(campaign.Status(req.Status))
	}

	return update.Save(ctx)
}

// GetCampaign retrieves a campaign by ID
func (s *CampaignService) GetCampaign(ctx context.Context, id int64) (*ent.Campaign, error) {
	return s.client.Campaign.Get(ctx, id)
}

// ListCampaigns retrieves campaigns with optional filtering
func (s *CampaignService) ListCampaigns(ctx context.Context, filter *dto.CampaignFilter) ([]*ent.Campaign, error) {
	query := s.client.Campaign.Query()

	if filter.Status != "" {
		query = query.Where(campaign.StatusEQ(campaign.Status(filter.Status)))
	}
	if filter.Type != "" {
		query = query.Where(campaign.TypeEQ(campaign.Type(filter.Type)))
	}
	if !filter.StartDateFrom.IsZero() {
		query = query.Where(campaign.StartDateGTE(filter.StartDateFrom))
	}
	if !filter.StartDateTo.IsZero() {
		query = query.Where(campaign.StartDateLTE(filter.StartDateTo))
	}

	return query.All(ctx)
}

// DeleteCampaign soft deletes a campaign
func (s *CampaignService) DeleteCampaign(ctx context.Context, id int64) error {
	return s.client.Campaign.UpdateOneID(id).
		SetStatus("deleted").
		Exec(ctx)
}

// Helper methods
func (s *CampaignService) validateCampaignRequest(req *dto.CreateCampaignRequest) error {
	if req.Name == "" {
		return errors.New("campaign name is required")
	}
	if req.Type == "" {
		return errors.New("campaign type is required")
	}
	if req.PayoutRate < 0 {
		return errors.New("payout rate cannot be negative")
	}
	return nil
}

// generateCampaignCode creates a unique tracking code
func (s *CampaignService) generateCampaignCode() string {
	return uuid.New().String()[:8]
}

// generateTrackingURL creates a unique tracking URL
func (s *CampaignService) generateTrackingURL(uniqueCode string) string {
	return fmt.Sprintf("https://track.yourdomain.com/%s", uniqueCode)
}
