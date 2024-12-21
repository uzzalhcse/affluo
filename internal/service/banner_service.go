package service

import (
	"affluo/constant"
	"affluo/ent"
	"affluo/ent/banner"
	"affluo/ent/bannercreative"
	"affluo/ent/creative"
	"affluo/internal/dto"
	"context"
	"fmt"
	"strings"
	"time"
)

type BannerService struct {
	client *ent.Client
}

func NewBannerService(client *ent.Client) *BannerService {
	return &BannerService{client: client}
}

func (s *BannerService) GetAllPublisherBanners(ctx context.Context) ([]*dto.BannerResponse, error) {
	banners, err := s.client.Banner.Query().
		Where(
			banner.StatusEQ(banner.StatusActive),
		).
		Order(ent.Asc(banner.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var responses []*dto.BannerResponse
	for _, b := range banners {
		response, err := s.enrichBannerResponse(ctx, b)
		if err != nil {
			return nil, err
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (s *BannerService) GetAllCreatives(ctx context.Context) ([]*ent.Creative, error) {
	return s.client.Creative.Query().
		Order(ent.Asc(creative.FieldCreatedAt)).
		All(ctx)
}
func (s *BannerService) CreateBannerCreative(ctx context.Context, req *dto.CreateBannerCreativeRequest) (*ent.Creative, error) {
	return s.client.Creative.Create().
		SetName(req.Name).
		SetImageURL(req.ImageURL).
		SetSize(req.Size).
		Save(ctx)
}

func (s *BannerService) GetBannerCreativesByBannerID(ctx context.Context, bannerID int64) ([]*ent.Creative, error) {
	return s.client.Creative.Query().
		Where(
			creative.HasBannerCreativesWith(
				bannercreative.BannerIDEQ(bannerID),
			),
		).
		Order(ent.Asc(creative.FieldCreatedAt)).
		All(ctx)
}

func (s *BannerService) UpdateBannerCreativeOrder(ctx context.Context, bannerID int64, orders []dto.CreativeOrder) error {
	// Start a transaction
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return err
	}

	for _, order := range orders {
		err := tx.BannerCreative.Update().
			Where(
				bannercreative.BannerIDEQ(bannerID),
				bannercreative.CreativeIDEQ(order.CreativeID),
			).
			SetDisplayOrder(order.Order).
			Exec(ctx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (s *BannerService) RemoveCreativeFromBanner(ctx context.Context, bannerID, creativeID int64) error {
	_, err := s.client.BannerCreative.Delete().
		Where(
			bannercreative.BannerIDEQ(bannerID),
			bannercreative.CreativeIDEQ(creativeID),
		).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *BannerService) SetPrimaryCreative(ctx context.Context, bannerID, creativeID int64) error {
	// Start a transaction
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return err
	}

	// First, set all creatives for this banner as non-primary
	err = tx.BannerCreative.Update().
		Where(bannercreative.BannerIDEQ(bannerID)).
		SetIsPrimary(false).
		Exec(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Then set the specified creative as primary
	err = tx.BannerCreative.Update().
		Where(
			bannercreative.BannerIDEQ(bannerID),
			bannercreative.CreativeIDEQ(creativeID),
		).
		SetIsPrimary(true).
		Exec(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// CreateBanner creates a banner with associated creatives
func (s *BannerService) CreateBanner(ctx context.Context, req *dto.CreateBannerRequest) (*dto.BannerResponse, error) {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	// Create banner with all available fields from schema
	b, err := tx.Banner.Create().
		SetName(req.Name).
		SetDescription(req.Description).
		SetClickURL(req.ClickURL).
		SetType(banner.Type(req.Type)).
		SetSize(req.Size).
		SetStatus(banner.Status(req.Status)).
		SetAllowedCountries(req.AllowedCountries).
		SetWeight(1). // Default weight
		SetNillableStartDate(nil).
		SetNillableEndDate(nil).
		SetAllowedDevices([]string{}).
		SetAllowedBrowsers([]string{}).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed creating banner: %w", err)
	}

	// Create banner-creative relationships if creatives are provided
	if len(req.CreativeIDs) > 0 {
		for i, creativeID := range req.CreativeIDs {
			_, err := tx.BannerCreative.Create().
				SetBannerID(b.ID).
				SetCreativeID(creativeID).
				SetDisplayOrder(i + 1).
				SetIsPrimary(i == 0).
				Save(ctx)
			if err != nil {
				tx.Rollback()
				return nil, fmt.Errorf("failed creating banner-creative relationship: %w", err)
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return s.GetBannerByID(ctx, b.ID)
}

// GetBannerByID retrieves a banner with all its details
func (s *BannerService) GetBannerByID(ctx context.Context, id int64) (*dto.BannerResponse, error) {
	b, err := s.client.Banner.Query().
		Where(banner.IDEQ(id)).
		WithCreatives().
		First(ctx)
	if err != nil {
		return nil, err
	}

	return s.enrichBannerResponse(ctx, b)
}

// GetAllBanners retrieves all banners with their creatives
func (s *BannerService) GetAllBanners(ctx context.Context) (*dto.BannersResponse, error) {
	banners, err := s.client.Banner.Query().
		WithCreatives().
		All(ctx)
	if err != nil {
		return nil, err
	}

	return dto.NewBannersResponse(banners), nil
}

// GetActiveBanners retrieves all active banners
func (s *BannerService) GetActiveBanners(ctx context.Context) ([]*dto.BannerResponse, error) {
	now := time.Now()
	banners, err := s.client.Banner.Query().
		Where(
			banner.And(
				banner.StatusEQ(banner.StatusActive),
				banner.Or(
					banner.StartDateIsNil(),
					banner.StartDateLTE(now),
				),
				banner.Or(
					banner.EndDateIsNil(),
					banner.EndDateGT(now),
				),
			),
		).
		WithCreatives().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var response []*dto.BannerResponse
	for _, b := range banners {
		enriched, err := s.enrichBannerResponse(ctx, b)
		if err != nil {
			return nil, err
		}
		response = append(response, enriched)
	}

	return response, nil
}

// UpdateBanner updates a banner's details
func (s *BannerService) UpdateBanner(ctx context.Context, id int64, req *dto.CreateBannerRequest) (*dto.BannerResponse, error) {
	mutation := s.client.Banner.UpdateOneID(id).
		SetName(req.Name).
		SetDescription(req.Description).
		SetClickURL(req.ClickURL).
		SetType(banner.Type(req.Type)).
		SetSize(req.Size).
		SetStatus(banner.Status(req.Status)).
		SetAllowedCountries(req.AllowedCountries)

	b, err := mutation.Save(ctx)
	if err != nil {
		return nil, err
	}

	return s.GetBannerByID(ctx, b.ID)
}

// CreateCreative creates a new creative
func (s *BannerService) CreateCreative(ctx context.Context, req *dto.CreateBannerCreativeRequest) (*dto.BannerCreativeResponse, error) {
	c, err := s.client.Creative.Create().
		SetName(req.Name).
		SetImageURL(req.ImageURL).
		SetSize(req.Size).
		SetEnabled(req.Enabled).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &dto.BannerCreativeResponse{
		ID:        c.ID,
		Name:      c.Name,
		ImageURL:  c.ImageURL,
		Size:      c.Size,
		Enabled:   c.Enabled,
		CreatedAt: c.CreatedAt,
	}, nil
}

// AssignCreativeToBanner assigns a creative to a banner
func (s *BannerService) AssignCreativeToBanner(ctx context.Context, req *dto.AssignCreativeRequest) error {
	return s.client.BannerCreative.Create().
		SetBannerID(req.BannerID).
		SetCreativeID(req.CreativeID).
		SetDisplayOrder(req.DisplayOrder).
		SetIsPrimary(req.IsPrimary).
		Exec(ctx)
}

//func (s *BannerService) GetAllPublisherBanners(ctx context.Context) ([]*dto.BannerResponse, error) {
//	banners, err := s.client.Banner.Query().
//		WithCreatives().
//		Where(banner.StatusEQ(constant.StatusActive)).
//		All(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	var response []*dto.BannerResponse
//	for _, banner := range banners {
//		enrichedBanner, err := s.enrichBannerResponse(ctx, banner)
//		if err != nil {
//			return nil, err
//		}
//		response = append(response, enrichedBanner)
//	}
//
//	return response, nil
//}

// GetBannerCreatives gets all creatives for a banner
func (s *BannerService) GetBannerCreatives(ctx context.Context, bannerID int64) ([]*dto.BannerCreativeResponse, error) {
	creatives, err := s.client.BannerCreative.Query().
		Where(bannercreative.HasBannerWith(banner.ID(bannerID))).
		WithCreative().
		Order(ent.Asc(bannercreative.FieldDisplayOrder)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.BannerCreativeResponse, len(creatives))
	for i, bc := range creatives {
		creative := bc.Edges.Creative
		responses[i] = &dto.BannerCreativeResponse{
			ID:           creative.ID,
			Name:         creative.Name,
			ImageURL:     creative.ImageURL,
			Size:         creative.Size,
			Enabled:      creative.Enabled,
			CreatedAt:    creative.CreatedAt,
			DisplayOrder: bc.DisplayOrder,
			IsPrimary:    bc.IsPrimary,
		}
	}

	return responses, nil
}

// Helper functions for generating tracking related content
func (s *BannerService) enrichBannerResponse(ctx context.Context, banner *ent.Banner) (*dto.BannerResponse, error) {
	creative, err := s.client.Creative.Query().
		Where(
			creative.HasBannerCreativesWith(
				bannercreative.BannerIDEQ(banner.ID),
			),
			creative.EnabledEQ(true),
		).
		First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	response := &dto.BannerResponse{
		ID:               banner.ID,
		Name:             banner.Name,
		Description:      banner.Description,
		Type:             banner.Type.String(),
		Size:             banner.Size,
		Status:           banner.Status.String(),
		AllowedCountries: banner.AllowedCountries,
		CreatedAt:        banner.CreatedAt,
		TrackingURL:      s.generateTrackingURL(banner),
	}

	if creative != nil {
		response.HTMLCode = s.generateHTMLCode(banner, creative)
	}

	return response, nil
}
func (s *BannerService) generateTrackingURL(banner *ent.Banner) string {
	return fmt.Sprintf("%s/click/%d", constant.TrackingDomain, banner.ID)
}

func (s *BannerService) generateHTMLCode(banner *ent.Banner, creative *ent.Creative) string {
	dimensions := strings.Split(banner.Size, "x")
	width, height := dimensions[0], dimensions[1]

	return fmt.Sprintf(`
    <div class="aff-banner" style="position:relative;display:inline-block">
        <a href="javascript:void(0);" style="text-decoration:none">
            <img src="%s" 
                width="%s" 
                height="%s" 
                alt="%s" 
                style="border:none;display:block" 
            />
        </a>
        <script>%s</script>
    </div>`,
		creative.ImageURL,
		width,
		height,
		banner.Name,
		s.generateTrackingScript(banner),
	)
}
func (s *BannerService) generateTrackingScript(banner *ent.Banner) string {
	return fmt.Sprintf(`
    (function() {
        // Configuration
        const config = {
            trackingDomain: '%s',
            bannerID: %d,
            clickURL: '%s',
            debug: false
        };

        // Utility functions
        function generateUID() {
            return Date.now().toString(36) + Math.random().toString(36).substr(2);
        }

        function getDeviceInfo() {
            return {
                userAgent: navigator.userAgent,
                language: navigator.language,
                screenSize: window.screen.width + 'x' + window.screen.height,
                viewportSize: window.innerWidth + 'x' + window.innerHeight,
                timezone: Intl.DateTimeFormat().resolvedOptions().timeZone
            };
        }

        // Tracking function
        async function track(eventType, extraData = {}) {
            try {
                const baseData = {
                    timestamp: new Date().toISOString(),
                    eventID: generateUID(),
                    bannerID: config.bannerID,
                    url: window.location.href,
                    referer: document.referrer,
                    ...getDeviceInfo(),
                    ...extraData
                };

                const response = await fetch(config.trackingDomain + '/api/tracking/' + eventType + '/' + config.bannerID, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(baseData),
                    keepalive: true // Ensures the request completes even if page changes
                });

                if (!response.ok && config.debug) {
                    console.error('Tracking failed:', response.statusText);
                }
            } catch (error) {
                if (config.debug) {
                    console.error('Tracking error:', error);
                }
                
                // Fallback to image pixel for impression tracking
                if (eventType === 'impression') {
                    new Image().src = config.trackingDomain + '/api/tracking/pixel/' + config.bannerID + '?t=' + Date.now();
                }
            }
        }

        // Handle impression
        function trackImpression() {
            if (window.IntersectionObserver) {
                // Use Intersection Observer for viewability
                const observer = new IntersectionObserver((entries) => {
                    entries.forEach(entry => {
                        if (entry.isIntersecting) {
                            track('impression');
                            observer.disconnect();
                        }
                    });
                }, {
                    threshold: 0.5 // 50% visibility required
                });
                
				const container = document.currentScript?.parentElement || document.querySelector('.aff-banner');
				if (!container) {
					console.error('Container element not found');
					return;
				}
                observer.observe(container);
            } else {
                // Fallback for older browsers
                track('impression');
            }
        }

        // Handle click
        function handleClick(e) {
            e.preventDefault();
            
            track('click').then(() => {
                window.open(config.clickURL, '_blank');
            }).catch(() => {
                // Fallback: direct navigation if tracking fails
                window.open(config.clickURL, '_blank');
            });
        }

        // Initialize
        window.addEventListener('load', () => {
			const container = document.currentScript?.parentElement || document.querySelector('.aff-banner');
			if (!container) {
				console.error('Container element not found');
				return;
			}
            const link = container.querySelector('a');
            
            // Set up click handler
            if (link) {
                link.addEventListener('click', handleClick);
            }
            
            // Track impression
            trackImpression();
        });
    })();`, constant.TrackingDomain, banner.ID, banner.ClickURL)
}
