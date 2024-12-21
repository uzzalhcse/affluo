// service/tracking_service.go
package service

import (
	"affluo/ent"
	"affluo/ent/banner"
	"affluo/ent/bannerstats"
	"affluo/ent/campaign"
	"affluo/ent/lead"
	"affluo/internal/dto"
	"context"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

type TrackingService struct {
	client *ent.Client
	// Cache for banner weights to avoid frequent DB calls
	weightCache     map[int64]float64
	weightCacheMux  sync.RWMutex
	weightCacheTime time.Time
}

func NewTrackingService(client *ent.Client) *TrackingService {
	return &TrackingService{
		client:      client,
		weightCache: make(map[int64]float64),
	}
}

// RecordImpression records a banner impression
func (s *TrackingService) RecordImpression(ctx context.Context, bannerID int64, req *dto.ImpressionRequest) error {
	date := time.Now().UTC().Truncate(24 * time.Hour)

	// Use transaction to ensure atomic updates
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Get or create stats for today
	stats, err := tx.BannerStats.Query().
		Where(
			bannerstats.HasBannerWith(banner.IDEQ(bannerID)),
			bannerstats.DateEQ(date),
		).
		First(ctx)

	if ent.IsNotFound(err) {
		stats, err = tx.BannerStats.Create().
			SetBannerID(bannerID).
			SetDate(date).
			Save(ctx)
	}
	if err != nil {
		return err
	}

	// Update impression count
	_, err = tx.BannerStats.UpdateOne(stats).
		AddImpressions(1).
		Save(ctx)
	if err != nil {
		return err
	}

	// Update banner last impression time
	_, err = tx.Banner.UpdateOne(
		tx.Banner.Query().Where(banner.IDEQ(bannerID)).OnlyX(ctx),
	).
		SetLastImpression(time.Now()).
		Save(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// RecordClick records a banner click
func (s *TrackingService) RecordClick(ctx context.Context, bannerID int64, req *dto.ClickRequest) error {
	date := time.Now().UTC().Truncate(24 * time.Hour)

	tx, err := s.client.Tx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stats, err := tx.BannerStats.Query().
		Where(
			bannerstats.HasBannerWith(banner.IDEQ(bannerID)),
			bannerstats.DateEQ(date),
		).
		First(ctx)

	if ent.IsNotFound(err) {
		stats, err = tx.BannerStats.Create().
			SetBannerID(bannerID).
			SetDate(date).
			Save(ctx)
	}
	if err != nil {
		return err
	}

	// Update click count and CTR
	impressions := stats.Impressions
	newClicks := stats.Clicks + 1
	newCTR := float64(newClicks) / float64(impressions)

	_, err = tx.BannerStats.UpdateOne(stats).
		SetClicks(newClicks).
		SetCtr(newCTR).
		Save(ctx)
	if err != nil {
		return err
	}

	// Update smart weight based on performance
	err = s.updateSmartWeight(ctx, tx, bannerID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// RecordLead records a conversion/lead
func (s *TrackingService) RecordLead(ctx context.Context, bannerID int64, req *dto.LeadRequest) error {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Create lead record
	_, err = tx.Lead.Create().
		SetBannerID(bannerID).
		SetType(lead.Type(req.Type)).
		SetAmount(req.Amount).
		SetReferenceID(req.ReferenceID).
		SetIPAddress(req.IPAddress).
		SetUserAgent(req.UserAgent).
		SetMetadata(req.Metadata).
		Save(ctx)
	if err != nil {
		return err
	}

	// Update stats
	date := time.Now().UTC().Truncate(24 * time.Hour)
	stats, err := tx.BannerStats.Query().
		Where(
			bannerstats.HasBannerWith(banner.IDEQ(bannerID)),
			bannerstats.DateEQ(date),
		).
		First(ctx)

	if err != nil {
		return err
	}

	newLeads := stats.Leads + 1
	convRate := float64(newLeads) / float64(stats.Clicks)

	_, err = tx.BannerStats.UpdateOne(stats).
		SetLeads(newLeads).
		SetConversionRate(convRate).
		Save(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// SelectBanner implements smart banner rotation
func (s *TrackingService) SelectBanner(ctx context.Context, campaignID int64) (*ent.Banner, error) {
	// Get active banners for campaign
	banners, err := s.client.Banner.Query().
		Where(
			banner.HasCampaignsWith(campaign.IDEQ(campaignID)),
			banner.StatusEQ("active"),
			banner.Or(
				banner.StartDateIsNil(),
				banner.StartDateLTE(time.Now()),
			),
			banner.Or(
				banner.EndDateIsNil(),
				banner.EndDateGT(time.Now()),
			),
		).
		All(ctx)

	if err != nil {
		return nil, err
	}

	if len(banners) == 0 {
		return nil, errors.New("no active banners found")
	}

	// Get cached weights or calculate new ones
	s.weightCacheMux.RLock()
	if time.Since(s.weightCacheTime) > 5*time.Minute {
		s.weightCacheMux.RUnlock()
		s.updateWeightCache(ctx, banners)
	} else {
		s.weightCacheMux.RUnlock()
	}

	// Calculate total weight
	var totalWeight float64
	s.weightCacheMux.RLock()
	for _, b := range banners {
		totalWeight += s.weightCache[b.ID]
	}
	s.weightCacheMux.RUnlock()

	// Select banner using weighted random selection
	r := rand.Float64() * totalWeight
	var currentWeight float64

	s.weightCacheMux.RLock()
	defer s.weightCacheMux.RUnlock()

	for _, b := range banners {
		currentWeight += s.weightCache[b.ID]
		if r <= currentWeight {
			return b, nil
		}
	}

	// Fallback to first banner if something goes wrong
	return banners[0], nil
}

// updateSmartWeight updates a banner's smart weight based on performance
func (s *TrackingService) updateSmartWeight(ctx context.Context, tx *ent.Tx, bannerID int64) error {
	// Get last 7 days of stats
	stats, err := tx.BannerStats.Query().
		Where(
			bannerstats.HasBannerWith(banner.IDEQ(bannerID)),
			bannerstats.DateGTE(time.Now().AddDate(0, 0, -7)),
		).
		All(ctx)
	if err != nil {
		return err
	}

	// Calculate weighted average CTR and conversion rate
	var totalCTR, totalConvRate, weightSum float64
	for i, stat := range stats {
		weight := math.Pow(0.8, float64(len(stats)-i-1)) // More recent days have higher weight
		totalCTR += stat.Ctr * weight
		totalConvRate += stat.ConversionRate * weight
		weightSum += weight
	}

	avgCTR := totalCTR / weightSum
	avgConvRate := totalConvRate / weightSum

	// Calculate smart weight (simplified version)
	smartWeight := (avgCTR + avgConvRate) * 100

	// Update banner smart weight
	_, err = tx.Banner.UpdateOne(
		tx.Banner.Query().Where(banner.IDEQ(bannerID)).OnlyX(ctx),
	).
		SetSmartWeight(smartWeight).
		Save(ctx)

	return err
}

// updateWeightCache updates the cached weights for all banners
func (s *TrackingService) updateWeightCache(ctx context.Context, banners []*ent.Banner) {
	s.weightCacheMux.Lock()
	defer s.weightCacheMux.Unlock()

	for _, b := range banners {
		if b.SmartWeight == 0 {
			s.weightCache[b.ID] = float64(b.Weight)
		} else {
			s.weightCache[b.ID] = b.SmartWeight
		}
	}
	s.weightCacheTime = time.Now()
}

// Add this method to the TrackingService struct
func (s *TrackingService) GetStats(ctx context.Context, bannerID int64, startDateStr, endDateStr string) (*dto.StatsAggregateResponse, error) {
	// Parse date strings
	var startDate, endDate time.Time
	var err error

	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			return nil, fmt.Errorf("invalid start date format: %w", err)
		}
	} else {
		// Default to last 30 days if no start date provided
		startDate = time.Now().AddDate(0, 0, -30)
	}

	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			return nil, fmt.Errorf("invalid end date format: %w", err)
		}
	} else {
		endDate = time.Now()
	}

	// Query stats for the date range
	stats, err := s.client.BannerStats.Query().
		Where(
			bannerstats.HasBannerWith(banner.IDEQ(bannerID)),
			bannerstats.DateGTE(startDate),
			bannerstats.DateLTE(endDate),
		).
		Order(ent.Asc(bannerstats.FieldDate)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch stats: %w", err)
	}

	// Prepare response
	response := &dto.StatsAggregateResponse{
		DailyStats: make([]dto.StatsResponse, len(stats)),
	}

	// Calculate totals and averages
	for i, stat := range stats {
		response.TotalImpressions += stat.Impressions
		response.TotalClicks += stat.Clicks
		response.TotalLeads += stat.Leads

		// Convert to response format
		response.DailyStats[i] = dto.StatsResponse{
			BannerID:       bannerID,
			Date:           stat.Date,
			Impressions:    stat.Impressions,
			Clicks:         stat.Clicks,
			Leads:          stat.Leads,
			CTR:            stat.Ctr,
			ConversionRate: stat.ConversionRate,
		}
	}

	// Calculate averages if we have any stats
	if len(stats) > 0 {
		response.AverageCTR = float64(response.TotalClicks) / float64(response.TotalImpressions)
		if response.TotalClicks > 0 {
			response.AverageConvRate = float64(response.TotalLeads) / float64(response.TotalClicks)
		}
	}

	return response, nil
}
