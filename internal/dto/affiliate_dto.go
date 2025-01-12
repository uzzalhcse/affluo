// internal/dto/affiliate_dto.go
package dto

type CreateAffiliateRequest struct {
	Source          string  `json:"source"`
	Commission      float64 `json:"commission"`
	UserID          int64   `json:"user_id"`
	AffiliateUserID string  `json:"affiliate_user_id"`
}

type UpdateAffiliateRequest struct {
	Source     string  `json:"source"`
	Commission float64 `json:"commission"`
}
