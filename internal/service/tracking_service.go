// service/tracking_service.go
package service

import (
	"affluo/ent"
	"affluo/ent/affiliate"
	"affluo/ent/banner"
	"affluo/ent/bannerstats"
	"affluo/ent/campaign"
	"affluo/ent/commissionhistory"
	"affluo/ent/commissionplan"
	"affluo/ent/earninghistory"
	"affluo/ent/gigtracking"
	"affluo/ent/lead"
	"affluo/ent/user"
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
func (s *TrackingService) RecordImpression(ctx context.Context, bannerID, publisherId int64, req *dto.ImpressionRequest) error {
	date := time.Now().UTC().Truncate(24 * time.Hour)

	// Use transaction to ensure atomic updates
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	// Get publisher's commission plan
	commission, err := s.getActiveCommissionPlan(ctx, tx, publisherId)
	if err != nil {
		return err
	}

	impressionRate := commission.ImpressionCommission
	earning := 0.0

	// Get or create stats for today
	stats, err := tx.BannerStats.Query().
		Where(
			bannerstats.HasBannerWith(banner.IDEQ(bannerID)),
			bannerstats.HasPublisherWith(user.IDEQ(publisherId)),
			bannerstats.DateEQ(date),
		).
		First(ctx)

	if ent.IsNotFound(err) {
		stats, err = tx.BannerStats.Create().
			SetBannerID(bannerID).
			SetPublisherID(publisherId).
			SetDate(date).
			Save(ctx)
	}
	if err != nil {
		return err
	}

	earning = float64(stats.Impressions+1)*impressionRate + float64(stats.Clicks)*commission.ClickCommission
	// Update impression count
	_, err = tx.BannerStats.UpdateOne(stats).
		AddImpressions(1).
		SetEarnings(earning).
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
func (s *TrackingService) RecordClick(ctx context.Context, bannerID, publisherID int64) (string, error) {
	date := time.Now().UTC().Truncate(24 * time.Hour)

	tx, err := s.client.Tx(ctx)
	if err != nil {
		return "", fmt.Errorf("starting transaction: %w", err)
	}
	defer tx.Rollback()

	// Add timeout context to prevent long-running queries
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Add Limit(1) to ensure we only get one result
	stats, err := tx.BannerStats.Query().
		Where(
			bannerstats.HasBannerWith(banner.IDEQ(bannerID)),
			//bannerstats.DateEQ(date),
			bannerstats.HasPublisherWith(user.IDEQ(publisherID)),
		).
		WithBanner().
		Limit(1).
		First(ctxWithTimeout)

	if ent.IsNotFound(err) {
		stats, err = tx.BannerStats.Create().
			SetBannerID(bannerID).
			SetPublisherID(publisherID).
			SetDate(date).
			Save(ctxWithTimeout)
	}
	if err != nil {
		return "", fmt.Errorf("querying banner stats: %w", err)
	}

	commission, err := s.getActiveCommissionPlan(ctxWithTimeout, tx, publisherID)
	if err != nil {
		return "", fmt.Errorf("getting commission plan: %w", err)
	}

	clickRate := commission.ClickCommission
	impressionRate := commission.ImpressionCommission

	// Calculate stats
	newClicks := stats.Clicks + 1
	impressions := stats.Impressions
	newCTR := float64(0)
	if impressions > 0 {
		newCTR = float64(newClicks) / float64(impressions)
	}

	// Calculate earnings
	clickEarnings := float64(newClicks) * clickRate
	impressionEarnings := float64(impressions) * impressionRate
	totalEarnings := clickEarnings + impressionEarnings

	// Update stats in a single query
	_, err = tx.BannerStats.UpdateOne(stats).
		SetClicks(newClicks).
		SetCtr(newCTR).
		SetEarnings(totalEarnings).
		Save(ctxWithTimeout)
	if err != nil {
		return "", fmt.Errorf("updating banner stats: %w", err)
	}

	// Update smart weight
	if err := s.updateSmartWeight(ctxWithTimeout, tx, bannerID); err != nil {
		return "", fmt.Errorf("updating smart weight: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return "", fmt.Errorf("committing transaction: %w", err)
	}

	return stats.Edges.Banner.ClickURL, nil
}

// Separate function for getting commission plan
func (s *TrackingService) getActiveCommissionPlan(ctx context.Context, tx *ent.Tx, publisherID int64) (*ent.CommissionPlan, error) {
	// First try to get personal commission plan
	personalPlan, err := tx.CommissionPlan.Query().
		Where(
			commissionplan.HasPublishersWith(user.IDEQ(publisherID)),
			commissionplan.IsActive(true),
		).
		First(ctx)

	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("error querying personal commission plan: %w", err)
	}

	if personalPlan != nil {
		return personalPlan, nil
	}

	// If no personal plan found, get default plan
	defaultPlan, err := tx.CommissionPlan.Query().
		Where(
			commissionplan.IsDefault(true),
			commissionplan.IsActive(true),
		).
		First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("no active commission plan found for publisher %d", publisherID)
		}
		return nil, fmt.Errorf("error querying default commission plan: %w", err)
	}

	return defaultPlan, nil
}
func (s *TrackingService) SyncVisit(ctx context.Context, pubId int64, landingPage, term, utm_query, trackingHash string) error {
	date := time.Now().UTC().Truncate(24 * time.Hour)
	// Check if publisher exists
	publisher, err := s.client.User.Get(ctx, pubId)
	if err != nil {
		if ent.IsNotFound(err) {
			return fmt.Errorf("publisher not found: %d", pubId)
		}
		return fmt.Errorf("error fetching publisher: %w", err)
	}

	// Try to find existing tracking for the date
	tracking, err := s.client.GigTracking.Query().
		Where(
			gigtracking.HasPublisherWith(user.ID(pubId)),
			gigtracking.Type(term),
			gigtracking.Lp(landingPage),
			gigtracking.DateEQ(date),
		).
		Only(ctx)

	if err != nil {
		if !ent.IsNotFound(err) {
			return fmt.Errorf("error querying tracking: %w", err)
		}

		// Create new tracking if not found
		tracking, err = s.client.GigTracking.Create().
			SetPublisher(publisher).
			SetDate(date).
			SetLp(landingPage).
			SetType(term).
			SetUtmQuery(utm_query).
			SetTrackID(trackingHash).
			Save(ctx)

		if err != nil {
			return fmt.Errorf("error creating tracking: %w", err)
		}
	} else {
		fmt.Printf("Found existing tracking: %+v\n", tracking)
		// Update existing tracking
		tracking, err = tracking.Update().
			SetLp(landingPage).
			SetUtmQuery(utm_query).
			Save(ctx)

		fmt.Printf("after update: %+v\n", tracking)
		if err != nil {
			return fmt.Errorf("error updating tracking: %w", err)
		}
	}

	return nil
}

func (s *TrackingService) GigLead(ctx context.Context, publisherID int64, trackID, targetType, affiliateUserID string, price float64) error {
	// Input validation
	if err := validateInput(trackID, affiliateUserID); err != nil {
		return fmt.Errorf("input validation failed: %w", err)
	}

	// Start transaction with timeout context
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	tx, err := s.client.Tx(ctxWithTimeout)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback()

	// Execute business logic within transaction
	if err := s.processGigLead(ctxWithTimeout, tx, publisherID, trackID, targetType, affiliateUserID, price); err != nil {
		return err
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func validateInput(trackID, affiliateUserID string) error {
	if trackID == "" {
		return fmt.Errorf("track ID cannot be empty")
	}
	if affiliateUserID == "" {
		return fmt.Errorf("affiliate user ID cannot be empty")
	}
	return nil
}

func (s *TrackingService) processGigLead(ctx context.Context, tx *ent.Tx, publisherID int64, trackID, targetType, affiliateUserID string, price float64) error {
	// Query existing tracking record
	affiliateList, err := tx.Affiliate.Query().
		Where(
			affiliate.TrackingCodeEQ(trackID),
			affiliate.AffiliateUserIDEQ(affiliateUserID),
			affiliate.HasUserWith(user.IDEQ(publisherID)),
		).
		WithUser().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return fmt.Errorf("affiliate not found: %s", trackID)
		}
		return fmt.Errorf("failed to query affiliate record: %w", err)
	}

	if affiliateList.Edges.User.ID != publisherID {
		return fmt.Errorf("publisher ID does not match affiliate record") // TODO: should this be an error? we should allow but less commission for non-publisher affiliates
	}

	// Check for duplicate earning
	exists, err := tx.EarningHistory.Query().
		Where(
			earninghistory.TrackID(trackID),
			earninghistory.EventType("lead"),
			earninghistory.Source("gigs"),
			earninghistory.HasUserWith(user.IDEQ(publisherID)),
		).
		Exist(ctx)
	if err != nil {
		return fmt.Errorf("failed to check duplicate earning: %w", err)
	}
	if exists {
		return fmt.Errorf("duplicate earning detected for track ID: %s", trackID)
	}

	// Get commission plan
	commission, err := s.getActiveCommissionPlan(ctx, tx, publisherID)
	if err != nil {
		return fmt.Errorf("failed to get commission plan: %w", err)
	}

	// Update tracking record and histories atomically
	if err := s.updateTrackingAndHistories(ctx, tx, commission, targetType, affiliateUserID, trackID, publisherID, price); err != nil {
		return fmt.Errorf("failed to update tracking and histories: %w", err)
	}

	return nil
}

// Updated service code
func (s *TrackingService) updateTrackingAndHistories(ctx context.Context, tx *ent.Tx, commission *ent.CommissionPlan,
	targetType, affiliateUserID, trackID string, publisherID int64, price float64) error {

	// Check if this is the first order for this affiliate-publisher pair
	isFirstOrder, err := tx.CommissionHistory.Query().
		Where(
			commissionhistory.AffiliateUserID(affiliateUserID),
			commissionhistory.HasUserWith(user.IDEQ(publisherID)),
		).
		Exist(ctx)
	if err != nil {
		return fmt.Errorf("failed to check order history: %w", err)
	}

	// Get the commission rate based on order history
	var commissionRate float64
	if !isFirstOrder {
		commissionRate = commission.FirstLeadCommission
	} else {
		// Check if within valid months period
		firstOrder, err := tx.CommissionHistory.Query().
			Where(
				commissionhistory.AffiliateUserID(affiliateUserID),
				commissionhistory.HasUserWith(user.IDEQ(publisherID)),
			).
			Order(ent.Desc(commissionhistory.FieldDate)).
			First(ctx)
		if err != nil {
			return fmt.Errorf("failed to get first order date: %w", err)
		}

		// Parse the date string to time.Time
		firstOrderDate, err := time.Parse("2006-01-02", firstOrder.Date)
		if err != nil {
			return fmt.Errorf("failed to parse first order date: %w", err)
		}

		// Check if within valid period
		monthsSinceFirst := time.Since(firstOrderDate).Hours() / 24 / 30
		if monthsSinceFirst <= float64(commission.ValidMonths) {
			commissionRate = commission.RepeatLeadCommission
		} else {
			commissionRate = 0 // Or set to a default rate for expired period
		}
	}

	// Calculate actual commission amount
	commissionAmount := (commissionRate / 100.0) * price

	// Update affiliate commission
	if err := tx.Affiliate.Update().
		Where(
			affiliate.TrackingCode(trackID),
			affiliate.AffiliateUserID(affiliateUserID),
		).
		AddCommission(commissionAmount).
		Exec(ctx); err != nil {
		return fmt.Errorf("failed to update affiliate commission: %w", err)
	}

	// Create earning history
	if _, err := tx.EarningHistory.Create().
		SetAmount(commissionAmount).
		SetEventType("lead").
		SetSource("gigs").
		SetTrackID(trackID).
		SetUserID(publisherID).
		Save(ctx); err != nil {
		return fmt.Errorf("failed to create earning history: %w", err)
	}

	// Create commission history with additional metadata
	if _, err := tx.CommissionHistory.Create().
		SetAmount(commissionAmount).
		SetAffiliateUserID(affiliateUserID).
		SetTrackID(trackID).
		SetUserID(publisherID).
		SetCommissionRate(commissionRate).
		SetIsFirstOrder(!isFirstOrder).
		Save(ctx); err != nil {
		return fmt.Errorf("failed to create commission history: %w", err)
	}

	return nil
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
			//bannerstats.HasBannerWith(banner.IDEQ(bannerID)),
			bannerstats.DateGTE(startDate),
			bannerstats.DateLTE(endDate),
		).
		WithBanner().
		Order(ent.Asc(bannerstats.FieldDate)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch stats: %w", err)
	}

	// Prepare response
	response := &dto.StatsAggregateResponse{
		Items: make([]dto.StatsResponse, len(stats)),
	}

	// Calculate totals and averages
	for i, stat := range stats {
		response.TotalImpressions += stat.Impressions
		response.TotalClicks += stat.Clicks
		response.TotalLeads += stat.Leads
		response.TotalEarning += stat.Earnings

		// Convert to response format
		response.Items[i] = dto.StatsResponse{
			BannerName:     stat.Edges.Banner.Name,
			Date:           stat.Date,
			Impressions:    stat.Impressions,
			Clicks:         stat.Clicks,
			Leads:          stat.Leads,
			CTR:            stat.Ctr,
			Earning:        stat.Earnings,
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

func (s *TrackingService) GigReports(ctx context.Context, publisherId int64, startDateStr, endDateStr string) (*dto.GigStats, error) {
	startDate, endDate, err := parseDateRange(startDateStr, endDateStr)
	if err != nil {
		return nil, err
	}

	// Query stats for the date range
	items, err := s.client.GigTracking.Query().
		Where(
			gigtracking.DateGTE(startDate),
			gigtracking.DateLTE(endDate),
			gigtracking.HasPublisherWith(user.ID(publisherId)),
			gigtracking.RevenueGT(0),
		).
		Order(ent.Asc(gigtracking.FieldDate)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch stats: %w", err)
	}

	// Calculate summary metrics
	stats := calculateSummaryMetrics(items)
	stats.Items = items

	return stats, nil
}

func parseDateRange(startDateStr, endDateStr string) (time.Time, time.Time, error) {
	var startDate, endDate time.Time
	var err error

	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			return startDate, endDate, fmt.Errorf("invalid start date format: %w", err)
		}
	} else {
		startDate = time.Now().AddDate(0, 0, -30)
	}

	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			return startDate, endDate, fmt.Errorf("invalid end date format: %w", err)
		}
	} else {
		endDate = time.Now()
	}

	return startDate, endDate, nil
}

func calculateSummaryMetrics(items []*ent.GigTracking) *dto.GigStats {
	stats := &dto.GigStats{}
	totalClicks := 0

	for _, item := range items {
		stats.TotalRevenue += item.Revenue
		totalClicks++

		// Track leads (assuming revenue > 0 indicates a conversion)
		if item.Revenue > 0 {
			stats.TotalLeads++
		}
	}

	// Calculate averages and rates
	if len(items) > 0 {
		stats.AverageRevenue = stats.TotalRevenue / float64(len(items))
	}
	if totalClicks > 0 {
		stats.ConversionRate = float64(stats.TotalLeads) / float64(totalClicks)
	}

	stats.TotalClicks = totalClicks

	return stats
}
