package dto

import "time"

type CreateCampaignRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Type        string    `json:"type"`
	PayoutRate  float64   `json:"payout_rate"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date,omitempty"`
}

type UpdateCampaignRequest struct {
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Type        string    `json:"type,omitempty"`
	PayoutRate  float64   `json:"payout_rate,omitempty"`
	StartDate   time.Time `json:"start_date,omitempty"`
	EndDate     time.Time `json:"end_date,omitempty"`
	Status      string    `json:"status,omitempty"`
}

type CampaignFilter struct {
	Status        string    `json:"status,omitempty"`
	Type          string    `json:"type,omitempty"`
	StartDateFrom time.Time `json:"start_date_from,omitempty"`
	StartDateTo   time.Time `json:"start_date_to,omitempty"`
}
