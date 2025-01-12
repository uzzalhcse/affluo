// dto/commission.go
package dto

type CreateCommissionPlanRequest struct {
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Type                 string  `json:"type"`
	ClickCommission      float64 `json:"click_commission"`
	ImpressionCommission float64 `json:"impression_commission"`
	LeadCommission       float64 `json:"lead_commission"`
	SaleCommissionRate   float64 `json:"sale_commission_rate"`
	MinimumPayout        float64 `json:"minimum_payout"`
	ValidFrom            string  `json:"valid_from"`
	ValidUntil           string  `json:"valid_until"`
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
	LeadCommission       float64 `json:"lead_commission"`
	SaleCommissionRate   float64 `json:"sale_commission_rate"`
	MinimumPayout        float64 `json:"minimum_payout"`
	ValidFrom            string  `json:"valid_from"`
	ValidUntil           string  `json:"valid_until"`
	IsActive             bool    `json:"is_active"`
	IsDefault            bool    `json:"is_default"`
}
