// dto/commission.go
package dto

type CreateCommissionPlanRequest struct {
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Type                 string  `json:"type"`
	ClickCommission      float64 `json:"click_commission"`
	ImpressionCommission float64 `json:"impression_commission"`
	FirstLeadCommission  float64 `json:"first_lead_commission"`
	RepeatLeadCommission float64 `json:"repeat_lead_commission"`
	ValidMonths          int     `json:"valid_months"`
	MinimumPayout        float64 `json:"minimum_payout"`
	IsActive             bool    `json:"is_active"`
	IsDefault            bool    `json:"is_default"`
}

type CommissionPlanResponse struct {
	ID                   int64   `json:"id"`
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Type                 string  `json:"type"`
	ClickCommission      float64 `json:"click_commission"`
	ImpressionCommission float64 `json:"impression_commission"`
	FirstLeadCommission  float64 `json:"first_lead_commission"`
	RepeatLeadCommission float64 `json:"repeat_lead_commission"`
	ValidMonths          int     `json:"valid_months"`
	MinimumPayout        float64 `json:"minimum_payout"`
	IsActive             bool    `json:"is_active"`
	IsDefault            bool    `json:"is_default"`
}
