package service

import (
	"affluo/ent"
	"affluo/ent/banner"
	"affluo/ent/bannercreative"
	"affluo/internal/dto"
	"context"
	"fmt"
)

type BannerService struct {
	client *ent.Client
}

func NewBannerService(client *ent.Client) *BannerService {
	return &BannerService{client: client}
}
func (s *BannerService) GetAllBanners(ctx context.Context) ([]*ent.Banner, error) {
	return s.client.Banner.Query().All(ctx)
}

func (s *BannerService) GetAllBannerCreatives(ctx context.Context) ([]*ent.BannerCreative, error) {
	return s.client.BannerCreative.Query().All(ctx)
}
func (s *BannerService) CreateBanner(ctx context.Context, req *dto.CreateBannerRequest) (*dto.BannerResponse, error) {
	// Begin a transaction
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	// Create banner
	b, err := tx.Banner.Create().
		SetName(req.Name).
		SetDescription(req.Description).
		SetClickURL(req.ClickURL).
		SetType(banner.Type(req.Type)).
		SetSize(req.Size).
		SetStatus(banner.Status(req.Status)).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed creating banner: %w", err)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	// Convert to response
	return &dto.BannerResponse{
		ID:               b.ID,
		Name:             b.Name,
		Description:      b.Description,
		Type:             string(b.Type),
		Size:             b.Size,
		AllowedCountries: b.AllowedCountries,
		Status:           string(b.Status),
		CreatedAt:        b.CreatedAt,
	}, nil
}

func (s *BannerService) GetBannerByID(ctx context.Context, id int64) (*dto.BannerResponse, error) {
	b, err := s.client.Banner.Query().
		Where(banner.IDEQ(id)).
		First(ctx)
	if err != nil {
		return nil, err
	}

	// Fetch associated creatives
	creatives, err := b.QueryCreatives().All(ctx)
	if err != nil {
		return nil, err
	}

	response := &dto.BannerResponse{
		ID:               b.ID,
		Name:             b.Name,
		Description:      b.Description,
		Type:             string(b.Type),
		Size:             b.Size,
		AllowedCountries: b.AllowedCountries,
		Status:           string(b.Status),
		Creatives:        creatives,
		CreatedAt:        b.CreatedAt,
	}

	return response, nil
}

func (s *BannerService) CreateBannerCreative(ctx context.Context, req *dto.CreateBannerCreativeRequest) (*dto.BannerCreativeResponse, error) {
	// Begin a transaction
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	// Create banner creative
	bc, err := tx.BannerCreative.Create().
		SetNillableName(func() *string {
			if req.Name != "" {
				return &req.Name
			}
			return nil
		}()).
		SetNillableImageURL(func() *string {
			if req.ImageURL != "" {
				return &req.ImageURL
			}
			return nil
		}()).
		SetNillableSize(func() *string {
			if req.Size != "" {
				return &req.Size
			}
			return nil
		}()).
		SetEnabled(req.Enabled).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed creating banner creative: %w", err)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	// Convert to response
	return &dto.BannerCreativeResponse{
		ID:        bc.ID,
		Name:      bc.Name,
		ImageURL:  bc.ImageURL,
		Enabled:   bc.Enabled,
		CreatedAt: bc.CreatedAt,
	}, nil
}

func (s *BannerService) GetBannerCreativesByBannerID(ctx context.Context, bannerID int64) ([]dto.BannerCreativeResponse, error) {
	// Fetch banner creatives
	creatives, err := s.client.BannerCreative.Query().
		Where(bannercreative.HasBannerWith(banner.IDEQ(bannerID))).
		All(ctx)
	if err != nil {
		return nil, err
	}

	// Convert to response
	responses := make([]dto.BannerCreativeResponse, len(creatives))
	for i, bc := range creatives {
		responses[i] = dto.BannerCreativeResponse{
			ID:        bc.ID,
			Name:      bc.Name,
			ImageURL:  bc.ImageURL,
			Enabled:   bc.Enabled,
			Size:      bc.Size,
			CreatedAt: bc.CreatedAt,
		}
	}

	return responses, nil
}
