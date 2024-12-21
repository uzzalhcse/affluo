package service

import (
	"affluo/constant"
	"affluo/ent"
	"affluo/ent/banner"
	"affluo/ent/bannercreative"
	"affluo/internal/dto"
	"context"
	"fmt"
	"strings"
)

type BannerService struct {
	client *ent.Client
}

func NewBannerService(client *ent.Client) *BannerService {
	return &BannerService{client: client}
}
func (s *BannerService) GetAllBanners(ctx context.Context) ([]*ent.Banner, error) {
	return s.client.Banner.Query().WithCreatives().All(ctx)
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

func (s *BannerService) generateTrackingURL(banner *ent.Banner) string {
	return fmt.Sprintf("%s/click/%d",
		constant.TrackingDomain,
		banner.ID)
}

// service/banner_service.go

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

func (s *BannerService) generateHTMLCode(banner *ent.Banner, creative *ent.BannerCreative) string {
	dimensions := strings.Split(banner.Size, "x")
	width, height := dimensions[0], dimensions[1]

	// Generate the HTML with embedded tracking script
	html := fmt.Sprintf(`
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

	return html
}
func (s *BannerService) enrichBannerResponse(ctx context.Context, banner *ent.Banner) (*dto.BannerResponse, error) {
	// Get the first creative for the banner where enabled = true
	creative, err := banner.QueryCreatives().
		Where(bannercreative.EnabledEQ(true)).
		First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	response := &dto.BannerResponse{
		ID:          banner.ID,
		Name:        banner.Name,
		Description: banner.Description,
		Type:        banner.Type.String(),
		Size:        banner.Size,
		Status:      banner.Status.String(),
		CreatedAt:   banner.CreatedAt,
		TrackingURL: s.generateTrackingURL(banner),
	}

	if creative != nil {
		response.HTMLCode = s.generateHTMLCode(banner, creative)
	}

	return response, nil
}

func (s *BannerService) GetAllPublisherBanners(ctx context.Context) ([]*dto.BannerResponse, error) {
	banners, err := s.client.Banner.Query().
		WithCreatives().
		Where(banner.StatusEQ(constant.StatusActive)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var response []*dto.BannerResponse
	for _, banner := range banners {
		enrichedBanner, err := s.enrichBannerResponse(ctx, banner)
		if err != nil {
			return nil, err
		}
		response = append(response, enrichedBanner)
	}

	return response, nil
}
