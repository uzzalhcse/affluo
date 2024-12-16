package handler

import (
	"affluo/internal/dto"
	"affluo/internal/service"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CampaignHandler struct {
	campaignService *service.CampaignService
}

func NewCampaignHandler(campaignService *service.CampaignService) *CampaignHandler {
	return &CampaignHandler{
		campaignService: campaignService,
	}
}

// CreateCampaign handles the creation of a new campaign
func (h *CampaignHandler) CreateCampaign(c *fiber.Ctx) error {
	var req dto.CreateCampaignRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}
	userIDValue := c.Locals("user_id")
	var userID int64

	switch v := userIDValue.(type) {
	case float64:
		userID = int64(v)
	case int64:
		userID = v
	default:
		return Unauthorized(c, "Invalid user ID type", fmt.Errorf("unexpected user ID type: %T", userIDValue))
	}
	// Create campaign
	campaign, err := h.campaignService.CreateCampaign(c.Context(), &req, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(campaign)
}

// UpdateCampaign handles updating an existing campaign
func (h *CampaignHandler) UpdateCampaign(c *fiber.Ctx) error {
	// Parse campaign ID from URL
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid campaign ID",
		})
	}

	var req dto.UpdateCampaignRequest
	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	// Update campaign
	campaign, err := h.campaignService.UpdateCampaign(c.Context(), id, &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(campaign)
}

// GetCampaign retrieves a specific campaign
func (h *CampaignHandler) GetCampaign(c *fiber.Ctx) error {
	// Parse campaign ID from URL
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid campaign ID",
		})
	}

	// Retrieve campaign
	campaign, err := h.campaignService.GetCampaign(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Campaign not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(campaign)
}

// ListCampaigns retrieves multiple campaigns
func (h *CampaignHandler) ListCampaigns(c *fiber.Ctx) error {
	// Parse filter parameters
	filter := &dto.CampaignFilter{
		Status:        c.Query("status"),
		Type:          c.Query("type"),
		StartDateFrom: parseQueryDate(c.Query("start_date_from")),
		StartDateTo:   parseQueryDate(c.Query("start_date_to")),
	}

	// List campaigns
	campaigns, err := h.campaignService.ListCampaigns(c.Context(), filter)
	if err != nil {
		return Error(c, "Failed to retrieve campaigns", err)
	}

	return Success(c, fiber.Map{
		"items": campaigns,
	})
}

// DeleteCampaign soft deletes a campaign
func (h *CampaignHandler) DeleteCampaign(c *fiber.Ctx) error {
	// Parse campaign ID from URL
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid campaign ID",
		})
	}

	// Delete campaign
	if err := h.campaignService.DeleteCampaign(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete campaign",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Campaign deleted successfully",
	})
}

// Helper function to parse date from query parameter
func parseQueryDate(dateStr string) time.Time {
	if dateStr == "" {
		return time.Time{}
	}

	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return time.Time{}
	}

	return date
}
func (h *CampaignHandler) GenerateTrackingLink(c *fiber.Ctx) error {
	campaignID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid campaign ID",
		})
	}

	// Generate unique tracking link with UTM parameters
	trackingLink, err := h.campaignService.GenerateUniqueTrackingLink(c.Context(), campaignID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate tracking link",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"tracking_link": trackingLink,
	})
}

// New method for retrieving campaign performance
func (h *CampaignHandler) GetCampaignPerformance(c *fiber.Ctx) error {
	campaignID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid campaign ID",
		})
	}

	// Retrieve performance metrics
	performance, err := h.campaignService.GetCampaignPerformance(c.Context(), campaignID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve campaign performance",
		})
	}

	return c.Status(fiber.StatusOK).JSON(performance)
}
