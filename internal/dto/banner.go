package dto

import (
	"affluo/ent"
	"time"
)

type AssignCreativeRequest struct {
	BannerID     int64 `json:"banner_id" validate:"required"`
	CreativeID   int64 `json:"creative_id" validate:"required"`
	DisplayOrder int   `json:"display_order"`
	IsPrimary    bool  `json:"is_primary"`
}

type CreativeOrder struct {
	CreativeID int64 `json:"creative_id" validate:"required"`
	Order      int   `json:"order" validate:"required"`
}

// Modify existing CreateBannerRequest to include creative IDs
type CreateBannerRequest struct {
	Name             string   `json:"name" validate:"required"`
	Description      string   `json:"description,omitempty"`
	ClickURL         string   `json:"click_url,omitempty"`
	Type             string   `json:"type" validate:"oneof=static dynamic"`
	Size             string   `json:"size" validate:"required"`
	Status           string   `json:"status" validate:"oneof=draft active inactive"`
	AllowedCountries []string `json:"allowed_countries,omitempty"`
	CreativeIDs      []int64  `json:"creative_ids,omitempty"` // Add this field
}

// Modify BannerCreativeResponse to include relationship metadata
type BannerCreativeResponse struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	ImageURL     string    `json:"image_url"`
	Size         string    `json:"size"`
	Enabled      bool      `json:"enabled"`
	DisplayOrder int       `json:"display_order,omitempty"`
	IsPrimary    bool      `json:"is_primary,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}

type BannerResponse struct {
	ID               int64                 `json:"id"`
	Name             string                `json:"name"`
	Description      string                `json:"description,omitempty"`
	Type             string                `json:"type"`
	Size             string                `json:"size"`
	HTMLCode         string                `json:"html_code,omitempty"`
	TrackingURL      string                `json:"tracking_url,omitempty"`
	AllowedCountries []string              `json:"allowed_countries,omitempty"`
	Status           string                `json:"status"`
	Creatives        []*ent.BannerCreative `json:"creatives,omitempty"`
	CreatedAt        time.Time             `json:"created_at"`
}
type BannersResponse struct {
	Banners []*BannerResponse `json:"items"`
}

func NewBannersResponse(banners []*ent.Banner) *BannersResponse {
	res := &BannersResponse{}
	for _, banner := range banners {
		res.Banners = append(res.Banners, &BannerResponse{
			ID:               banner.ID,
			Name:             banner.Name,
			Description:      banner.Description,
			Type:             banner.Type.String(),
			Size:             banner.Size,
			AllowedCountries: banner.AllowedCountries,
			Status:           banner.Status.String(),
			CreatedAt:        banner.CreatedAt,
			Creatives:        banner.Edges.BannerCreatives,
		})
	}
	return res
}

type CreateBannerCreativeRequest struct {
	Name     string `json:"name,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
	Enabled  bool   `json:"enabled,omitempty"`
	Size     string `json:"size,omitempty"`
}
