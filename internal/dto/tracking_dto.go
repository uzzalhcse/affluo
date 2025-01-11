// internal/dto/tracking.go
package dto

import (
	"affluo/ent"
	"time"
)

type ImpressionRequest struct {
	IPAddress  string            `json:"ip_address"`
	UserAgent  string            `json:"user_agent"`
	Referer    string            `json:"referer"`
	Metadata   map[string]string `json:"metadata"`
	DeviceType string            `json:"device_type"`
	Browser    string            `json:"browser"`
	OS         string            `json:"os"`
}

type ClickRequest struct {
	IPAddress  string            `json:"ip_address"`
	UserAgent  string            `json:"user_agent"`
	Referer    string            `json:"referer"`
	Metadata   map[string]string `json:"metadata"`
	DeviceType string            `json:"device_type"`
	Browser    string            `json:"browser"`
	OS         string            `json:"os"`
}

type LeadRequest struct {
	Type        string                 `json:"type"`
	Amount      float64                `json:"amount"`
	Currency    string                 `json:"currency"`
	ReferenceID string                 `json:"reference_id"`
	IPAddress   string                 `json:"ip_address"`
	UserAgent   string                 `json:"user_agent"`
	Metadata    map[string]interface{} `json:"metadata"`
}

type StatsResponse struct {
	BannerName     string    `json:"banner_name"`
	Date           time.Time `json:"date"`
	Impressions    int64     `json:"impressions"`
	Clicks         int64     `json:"clicks"`
	Leads          int64     `json:"leads"`
	CTR            float64   `json:"ctr"`
	Earning        float64   `json:"earning"`
	ConversionRate float64   `json:"conversion_rate"`
}

type StatsAggregateResponse struct {
	TotalImpressions int64           `json:"total_impressions"`
	TotalClicks      int64           `json:"total_clicks"`
	TotalLeads       int64           `json:"total_leads"`
	AverageCTR       float64         `json:"average_ctr"`
	AverageConvRate  float64         `json:"average_conv_rate"`
	TotalEarning     float64         `json:"total_earning"`
	Items            []StatsResponse `json:"items"`
}
type GigStats struct {
	Items           []*ent.GigTracking `json:"items"`
	TotalRevenue    float64            `json:"total_revenue"`
	TotalClicks     int                `json:"total_clicks"`
	TotalLeads      int                `json:"total_leads"`
	AverageRevenue  float64            `json:"average_revenue"`
	ConversionRate  float64            `json:"conversion_rate"`
	TopPerformingLP string             `json:"top_performing_lp"`
}
