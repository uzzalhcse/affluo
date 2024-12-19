package dto

import (
	"affluo/ent"
	"time"
)

type CreateBannerRequest struct {
	Name             string   `json:"name" validate:"required"`
	Description      string   `json:"description,omitempty"`
	ClickURL         string   `json:"click_url,omitempty"`
	Type             string   `json:"type" validate:"oneof=static dynamic"`
	Size             string   `json:"size" validate:"required"`
	Status           string   `json:"status" validate:"oneof=draft active inactive"`
	AllowedCountries []string `json:"allowed_countries,omitempty"`
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
			Creatives:        banner.Edges.Creatives,
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

type BannerCreativeResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	ImageURL  string    `json:"image_url"`
	Size      string    `json:"size"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
}

type BannerCreativesResponse struct {
	BannerCreatives []*BannerCreativeResponse `json:"items"`
}

func NewBannerCreativesResponse(bannerCreatives []*ent.BannerCreative) *BannerCreativesResponse {
	res := &BannerCreativesResponse{}
	for _, bannerCreative := range bannerCreatives {
		res.BannerCreatives = append(res.BannerCreatives, &BannerCreativeResponse{
			ID:        bannerCreative.ID,
			Name:      bannerCreative.Name,
			ImageURL:  bannerCreative.ImageURL,
			Enabled:   bannerCreative.Enabled,
			Size:      bannerCreative.Size,
			CreatedAt: bannerCreative.CreatedAt,
		})
	}
	return res
}
