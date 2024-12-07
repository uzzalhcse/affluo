package handler

import (
	"affluo/internal/dto"
	"affluo/internal/service"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type TrackingHandler struct {
	trackingService *service.TrackingService
}

func NewTrackingHandler(trackingService *service.TrackingService) *TrackingHandler {
	return &TrackingHandler{
		trackingService: trackingService,
	}
}

// HandleTrack processes tracking events
func (h *TrackingHandler) HandleTrack(c *fiber.Ctx) error {
	var trackRequest dto.TrackingRequest

	// Bind request data
	if err := c.BodyParser(&trackRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid tracking request",
		})
	}

	// Validate request
	if err := validateTrackingRequest(&trackRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Extract client details
	trackRequest.IPAddress = c.IP()
	trackRequest.UserAgent = string(c.Context().UserAgent())

	// Track the event
	track, err := h.trackingService.Track(c.Context(), &trackRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to track event",
		})
	}

	// Return tracking response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"track_id": track.ID,
		"status":   track.Status,
	})
}

// GenerateTrackingScript creates a secure client-side tracking script
func (h *TrackingHandler) GenerateTrackingScript(c *fiber.Ctx) error {
	campaignID := c.Params("campaign_id")

	// Generate secure tracking script
	script := h.trackingService.GenerateTrackingScript(campaignID)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"tracking_script": script,
	})
}

// validateTrackingRequest checks request integrity
func validateTrackingRequest(req *dto.TrackingRequest) error {
	if req.CampaignID == "" {
		return errors.New("campaign ID is required")
	}

	if req.TrackType == "" {
		return errors.New("tracking type is required")
	}

	return nil
}
