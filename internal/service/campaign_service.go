package service

import (
	"affluo/ent"
	"affluo/ent/campaign"
	"affluo/internal/dto"
	"context"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

type CampaignService struct {
	client *ent.Client
}

func NewCampaignService(client *ent.Client) *CampaignService {
	return &CampaignService{
		client: client,
	}
}

func (s *CampaignService) CreateCampaign(ctx context.Context, req *dto.CreateCampaignRequest, userID int64) (*dto.CampaignResponse, error) {
	// Generate unique campaign code
	uniqueCode := generateUniqueCampaignCode()

	// Create campaign
	campaignCreate := s.client.Campaign.Create().
		SetName(req.Name).
		SetDescription(req.Description).
		SetType(campaign.Type(req.Type)).
		SetCommissionType(campaign.CommissionType(req.CommissionType)).
		SetBaseCommissionRate(req.BaseCommissionRate).
		SetUniqueCode(uniqueCode).
		SetStartDate(req.StartDate).
		SetOwnerID(userID)

	// Optional fields
	if req.EndDate != nil {
		campaignCreate.SetEndDate(*req.EndDate)
	}
	if req.TargetGeography != "" {
		campaignCreate.SetTargetGeography(req.TargetGeography)
	}
	if len(req.TargetDemographics) > 0 {
		campaignCreate.SetTargetDemographics(req.TargetDemographics)
	}
	if len(req.CommissionTiers) > 0 {
		campaignCreate.SetCommissionTiers(req.CommissionTiers)
	}
	if req.TrackingURL != "" {
		campaignCreate.SetTrackingURL(req.TrackingURL)
	}

	// Persist campaign
	newCampaign, err := campaignCreate.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create campaign: %w", err)
	}

	// Convert to response DTO
	return s.convertToCampaignResponse(newCampaign), nil
}

func (s *CampaignService) UpdateCampaign(ctx context.Context, id int64, req *dto.UpdateCampaignRequest) (*dto.CampaignResponse, error) {
	// Start update builder
	update := s.client.Campaign.UpdateOneID(id)

	// Update optional fields
	if req.Name != nil {
		update.SetName(*req.Name)
	}
	if req.Description != nil {
		update.SetDescription(*req.Description)
	}
	if req.Type != nil {
		update.SetType(campaign.Type(*req.Type))
	}
	if req.CommissionType != nil {
		update.SetCommissionType(campaign.CommissionType(*req.CommissionType))
	}
	if req.BaseCommissionRate != nil {
		update.SetBaseCommissionRate(*req.BaseCommissionRate)
	}
	if req.CommissionTiers != nil {
		update.SetCommissionTiers(*req.CommissionTiers)
	}
	if req.TargetGeography != nil {
		update.SetTargetGeography(*req.TargetGeography)
	}
	if len(req.TargetDemographics) > 0 {
		update.SetTargetDemographics(req.TargetDemographics)
	}
	if req.StartDate != nil {
		update.SetStartDate(*req.StartDate)
	}
	if req.EndDate != nil {
		update.SetEndDate(*req.EndDate)
	}
	if req.Status != nil {
		update.SetStatus(campaign.Status(*req.Status))
	}
	if req.TrackingURL != nil {
		update.SetTrackingURL(*req.TrackingURL)
	}

	// Save updated campaign
	updatedCampaign, err := update.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update campaign: %w", err)
	}

	return s.convertToCampaignResponse(updatedCampaign), nil
}

func (s *CampaignService) GenerateUniqueTrackingLink(ctx context.Context, campaignID int64) (string, error) {
	// Retrieve campaign
	campaign, err := s.client.Campaign.Get(ctx, campaignID)
	if err != nil {
		return "", fmt.Errorf("campaign not found: %w", err)
	}

	// Generate unique tracking identifier
	trackingID := uuid.New().String()

	// Construct tracking URL with UTM parameters
	baseURL := campaign.TrackingURL
	if baseURL == "" {
		baseURL = "https://yourdefaulttrackingdomain.com"
	}

	trackingLink := fmt.Sprintf("%s?utm_campaign=%s&utm_source=affiliate&utm_medium=referral&track_id=%s",
		baseURL,
		campaign.UniqueCode,
		trackingID,
	)

	return trackingLink, nil
}

func (s *CampaignService) GetCampaignPerformance(ctx context.Context, campaignID int64) (*dto.CampaignPerformanceResponse, error) {
	// Retrieve campaign
	campaign, err := s.client.Campaign.Get(ctx, campaignID)
	if err != nil {
		return nil, fmt.Errorf("campaign not found: %w", err)
	}

	// Calculate performance metrics
	conversionRate := 0.0
	if campaign.TotalClicks > 0 {
		conversionRate = float64(campaign.TotalConversions) / float64(campaign.TotalClicks) * 100
	}

	averageCommission := 0.0
	if campaign.TotalConversions > 0 {
		averageCommission = campaign.TotalRevenue / float64(campaign.TotalConversions)
	}

	return &dto.CampaignPerformanceResponse{
		CampaignID:        campaign.ID,
		TotalClicks:       campaign.TotalClicks,
		TotalConversions:  campaign.TotalConversions,
		TotalRevenue:      campaign.TotalRevenue,
		ConversionRate:    conversionRate,
		AverageCommission: averageCommission,
	}, nil
}

// Helper function to generate unique campaign code
func generateUniqueCampaignCode() string {
	rand.Seed(time.Now().UnixNano())

	// Generate a random string of 8 characters
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 8)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}

	return string(b)
}

// Helper function to convert Ent Campaign to Response DTO
func (s *CampaignService) convertToCampaignResponse(c *ent.Campaign) *dto.CampaignResponse {
	response := &dto.CampaignResponse{
		ID:                 c.ID,
		Name:               c.Name,
		Description:        c.Description,
		Type:               c.Type.String(),
		CommissionType:     c.CommissionType.String(),
		BaseCommissionRate: c.BaseCommissionRate,
		StartDate:          c.StartDate,
		Status:             c.Status.String(),
		TotalClicks:        c.TotalClicks,
		TotalConversions:   c.TotalConversions,
		TotalRevenue:       c.TotalRevenue,
		ConversionRate:     c.ConversionRate,
		EndDate:            &c.EndDate,
	}

	if c.TargetGeography != "" {
		response.TargetGeography = c.TargetGeography
	}
	if c.TrackingURL != "" {
		response.TrackingURL = c.TrackingURL
	}
	if len(c.CommissionTiers) > 0 {
		response.CommissionTiers = c.CommissionTiers
	}
	if len(c.TargetDemographics) > 0 {
		response.TargetDemographics = c.TargetDemographics
	}

	return response
}
func (s *CampaignService) GetCampaign(ctx context.Context, id int64) (*dto.CampaignResponse, error) {
	// Retrieve campaign
	campaign, err := s.client.Campaign.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("campaign not found: %w", err)
	}

	// Convert to response DTO
	return s.convertToCampaignResponse(campaign), nil
}

func (s *CampaignService) ListCampaigns(ctx context.Context, filter *dto.CampaignFilter) ([]*dto.CampaignResponse, error) {
	// Start query builder
	query := s.client.Campaign.Query()

	// Apply filters
	if filter.Status != "" {
		query.Where(campaign.StatusEQ(campaign.Status(filter.Status)))
	}
	if filter.Type != "" {
		query.Where(campaign.TypeEQ(campaign.Type(filter.Type)))
	}

	// Date range filters
	if !filter.StartDateFrom.IsZero() {
		query.Where(campaign.StartDateGTE(filter.StartDateFrom))
	}
	if !filter.StartDateTo.IsZero() {
		query.Where(campaign.StartDateLTE(filter.StartDateTo))
	}

	// Execute query
	campaigns, err := query.All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve campaigns: %w", err)
	}

	// Convert to response DTOs
	responses := make([]*dto.CampaignResponse, len(campaigns))
	for i, c := range campaigns {
		responses[i] = s.convertToCampaignResponse(c)
	}

	return responses, nil
}

func (s *CampaignService) DeleteCampaign(ctx context.Context, id int64) error {
	// Soft delete campaign by updating its status
	_, err := s.client.Campaign.UpdateOneID(id).
		SetStatus("completed"). // Or use a soft delete status like 'deleted'
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete campaign: %w", err)
	}

	return nil
}
