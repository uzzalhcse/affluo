package dto

import (
	"affluo/ent/schema"
	"time"
)

type CreateCampaignRequest struct {
	Name               string                  `json:"name" validate:"required"`
	Description        string                  `json:"description,omitempty"`
	Type               string                  `json:"type" validate:"oneof=sale lead click subscription"`
	CommissionType     string                  `json:"commission_type" validate:"oneof=flat_rate percentage tiered"`
	BaseCommissionRate float64                 `json:"base_commission_rate,omitempty"`
	CommissionTiers    []schema.CommissionTier `json:"commission_tiers,omitempty"`
	TargetGeography    string                  `json:"target_geography,omitempty"`
	TargetDemographics map[string]interface{}  `json:"target_demographics,omitempty"`
	StartDate          time.Time               `json:"start_date" validate:"required"`
	EndDate            *time.Time              `json:"end_date,omitempty"`
	TrackingURL        string                  `json:"tracking_url,omitempty"`
}

type UpdateCampaignRequest struct {
	Name               *string                  `json:"name,omitempty"`
	Description        *string                  `json:"description,omitempty"`
	Type               *string                  `json:"type,omitempty"`
	CommissionType     *string                  `json:"commission_type,omitempty"`
	BaseCommissionRate *float64                 `json:"base_commission_rate,omitempty"`
	CommissionTiers    *[]schema.CommissionTier `json:"commission_tiers,omitempty"`
	TargetGeography    *string                  `json:"target_geography,omitempty"`
	TargetDemographics map[string]interface{}   `json:"target_demographics,omitempty"`
	StartDate          *time.Time               `json:"start_date,omitempty"`
	EndDate            *time.Time               `json:"end_date,omitempty"`
	Status             *string                  `json:"status,omitempty"`
	TrackingURL        *string                  `json:"tracking_url,omitempty"`
}

type CampaignResponse struct {
	ID                 int64                   `json:"id"`
	Name               string                  `json:"name"`
	Description        string                  `json:"description,omitempty"`
	Type               string                  `json:"type"`
	CommissionType     string                  `json:"commission_type"`
	BaseCommissionRate float64                 `json:"base_commission_rate"`
	CommissionTiers    []schema.CommissionTier `json:"commission_tiers,omitempty"`
	TargetGeography    string                  `json:"target_geography,omitempty"`
	TargetDemographics map[string]interface{}  `json:"target_demographics,omitempty"`
	StartDate          time.Time               `json:"start_date"`
	EndDate            *time.Time              `json:"end_date,omitempty"`
	Status             string                  `json:"status"`
	TrackingURL        string                  `json:"tracking_url,omitempty"`
	TotalClicks        int                     `json:"total_clicks"`
	TotalConversions   int                     `json:"total_conversions"`
	TotalRevenue       float64                 `json:"total_revenue"`
	ConversionRate     float64                 `json:"conversion_rate"`
}

type CampaignPerformanceResponse struct {
	CampaignID        int64   `json:"campaign_id"`
	TotalClicks       int     `json:"total_clicks"`
	TotalConversions  int     `json:"total_conversions"`
	TotalRevenue      float64 `json:"total_revenue"`
	ConversionRate    float64 `json:"conversion_rate"`
	AverageCommission float64 `json:"average_commission"`
}
type CampaignFilter struct {
	Status        string    `json:"status,omitempty"`
	Type          string    `json:"type,omitempty"`
	StartDateFrom time.Time `json:"start_date_from,omitempty"`
	StartDateTo   time.Time `json:"start_date_to,omitempty"`
}
