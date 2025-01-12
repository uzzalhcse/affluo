// handler/tracking_handler.go
package handler

import (
	"affluo/internal/dto"
	"affluo/internal/helper"
	"affluo/internal/service"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

type TrackingHandler struct {
	trackingService  *service.TrackingService
	affiliateService *service.AffiliateService
}

func NewTrackingHandler(trackingService *service.TrackingService, affiliateService *service.AffiliateService) *TrackingHandler {
	return &TrackingHandler{trackingService: trackingService, affiliateService: affiliateService}
}

func (h *TrackingHandler) RecordImpression(c *fiber.Ctx) error {
	bannerID, err := c.ParamsInt("id")
	if err != nil {
		return Error(c, "Invalid banner ID", err.Error())
	}

	publisherId := c.QueryInt("pub")
	req := new(dto.ImpressionRequest)
	if err := c.BodyParser(req); err != nil {
		return Error(c, "Invalid request body", err.Error())
	}

	err = h.trackingService.RecordImpression(c.Context(), int64(bannerID), int64(publisherId), req)
	if err != nil {
		return Error(c, "Failed to record impression", err.Error())
	}

	return Success(c, fiber.Map{"status": "recorded"})
}

// RecordClick handles the click tracking with encrypted tracking code
func (h *TrackingHandler) RecordClick(c *fiber.Ctx) error {
	bannerID := c.QueryInt("banner_id")
	publisherID := c.QueryInt("pub")

	clickUrl, err := h.trackingService.RecordClick(c.Context(), int64(bannerID), int64(publisherID))
	if err != nil {
		return Error(c, "Failed to record click", err.Error())
	}

	// Create tracking code and encrypt it
	trackingCode := fmt.Sprintf("publisher-%d_banner-%d", publisherID, bannerID)
	encryptedCode, err := helper.Encrypt(trackingCode)
	if err != nil {
		return Error(c, "Failed to generate tracking code", err.Error())
	}
	redirectUrl := fmt.Sprintf("%s?track_id=%s", clickUrl, encryptedCode)
	return c.Redirect(redirectUrl, 301)
}

func (h *TrackingHandler) RecordLead(c *fiber.Ctx) error {
	bannerID, err := c.ParamsInt("id")
	if err != nil {
		return Error(c, "Invalid banner ID", err.Error())
	}

	req := new(dto.LeadRequest)

	if err := c.BodyParser(req); err != nil {
		return Error(c, "Invalid request body", err.Error())
	}

	err = h.trackingService.RecordLead(c.Context(), int64(bannerID), req)
	if err != nil {
		return Error(c, "Failed to record lead", err.Error())
	}

	return Success(c, fiber.Map{"status": "recorded"})
}

func (h *TrackingHandler) SelectBanner(c *fiber.Ctx) error {
	campaignID, err := c.ParamsInt("campaign_id")
	if err != nil {
		return Error(c, "Invalid campaign ID", err.Error())
	}

	banner, err := h.trackingService.SelectBanner(c.Context(), int64(campaignID))
	if err != nil {
		return Error(c, "Failed to select banner", err.Error())
	}

	return Success(c, banner)
}

func (h *TrackingHandler) GetStats(c *fiber.Ctx) error {
	bannerID := c.QueryInt("id")

	// Parse date range from query parameters
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	stats, err := h.trackingService.GetStats(c.Context(), int64(bannerID), startDate, endDate)
	if err != nil {
		return Error(c, "Failed to get stats", err.Error())
	}

	return Success(c, fiber.Map{"stats": stats})
}
func (h *TrackingHandler) GetGigReports(c *fiber.Ctx) error {
	// Parse date range from query parameters
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	publisherId, err := helper.GetUserIDFromContext(c)
	if err != nil {
		return Error(c, "Failed to get reports", err.Error())
	}

	stats, err := h.trackingService.GigReports(c.Context(), publisherId, startDate, endDate)
	if err != nil {
		return Error(c, "Failed to get stats", err.Error())
	}

	return Success(c, fiber.Map{"stats": stats})
}

func (h *TrackingHandler) VisitTracking(c *fiber.Ctx) error {
	pubId := c.QueryInt("pub")
	landingPage := c.Query("lp")
	term := c.Query("type")
	terget := c.Query("target")
	utm_query := c.Query("utm_query")

	trackingCode := fmt.Sprintf("publisher-%d_%s-%s", pubId, terget, term)
	encryptedCode, err := helper.Encrypt(trackingCode)

	err = h.trackingService.SyncVisit(c.Context(), int64(pubId), landingPage, term, utm_query, encryptedCode)
	if err != nil {
		return Error(c, "Something went wrong", err.Error())
	}
	redirectUrl := fmt.Sprintf("%s?track_id=%s", landingPage, encryptedCode)

	return c.Redirect(redirectUrl, 301)
}
func (h *TrackingHandler) GigLeadCallBack(c *fiber.Ctx) error {
	trackId := c.Query("track_id")
	eventType := c.Query("event_type")
	affiliateUserId := c.Query("affiliate_user_id")
	publisherID, targetType, _, err := helper.DecodeTrackingCode(trackId)
	if err != nil {
		return Error(c, "Failed to decrypt tracking code", err.Error())
	}

	if eventType == "lead" {
		err = h.trackingService.GigLead(c.Context(), publisherID, trackId, targetType, affiliateUserId)
		if err != nil {
			return Error(c, "Something went wrong", err.Error())
		}
	} else if eventType == "create_account" {
		req := new(dto.CreateAffiliateRequest)
		req.UserID = publisherID
		req.Source = targetType
		req.TrackingCode = trackId
		req.AffiliateUserID = affiliateUserId
		_, err = h.affiliateService.CreateAffiliate(c.Context(), req)
		if err != nil {
			return Error(c, "Failed to create affiliate", err.Error())
		}
	} else {
		return fmt.Errorf("invalid event type")
	}
	return Success(c, fiber.Map{"status": "recorded"})
}

// handler/tracking_handler.go

func (h *TrackingHandler) PixelTracking(c *fiber.Ctx) error {
	bannerID, err := c.ParamsInt("id")
	if err != nil {
		return Error(c, "Invalid banner ID", err.Error())
	}

	publisherId := c.QueryInt("pub")
	// Record impression
	err = h.trackingService.RecordImpression(c.Context(), int64(bannerID), int64(publisherId), &dto.ImpressionRequest{
		IPAddress: c.IP(),
		UserAgent: c.Get("User-Agent"),
		Referer:   c.Get("Referer"),
	})
	if err != nil {
		// Log error but don't return it
		log.Printf("Failed to record pixel impression: %v", err)
	}

	// Return a transparent 1x1 GIF
	c.Set("Content-Type", "image/gif")
	c.Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Set("Pragma", "no-cache")
	c.Set("Expires", "0")

	// Transparent 1x1 GIF
	return c.Send([]byte{
		0x47, 0x49, 0x46, 0x38, 0x39, 0x61, 0x01, 0x00, 0x01, 0x00, 0x80, 0x00,
		0x00, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x21, 0xF9, 0x04, 0x01, 0x00,
		0x00, 0x00, 0x00, 0x2C, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x00, 0x02, 0x02, 0x44, 0x01, 0x00, 0x3B,
	})
}
