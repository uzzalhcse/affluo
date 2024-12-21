// internal/dto/tracking.go
package dto

import "time"

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
	BannerID       int64     `json:"banner_id"`
	Date           time.Time `json:"date"`
	Impressions    int64     `json:"impressions"`
	Clicks         int64     `json:"clicks"`
	Leads          int64     `json:"leads"`
	CTR            float64   `json:"ctr"`
	ConversionRate float64   `json:"conversion_rate"`
}

type StatsAggregateResponse struct {
	TotalImpressions int64           `json:"total_impressions"`
	TotalClicks      int64           `json:"total_clicks"`
	TotalLeads       int64           `json:"total_leads"`
	AverageCTR       float64         `json:"average_ctr"`
	AverageConvRate  float64         `json:"average_conv_rate"`
	DailyStats       []StatsResponse `json:"daily_stats"`
}
