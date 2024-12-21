package handler

import (
	"affluo/internal/dto"
	"affluo/internal/service"
	"github.com/gofiber/fiber/v2"
)

type BannerHandler struct {
	bannerService *service.BannerService
}

func NewBannerHandler(bannerService *service.BannerService) *BannerHandler {
	return &BannerHandler{bannerService: bannerService}
}
func (h *BannerHandler) GetAllBanners(c *fiber.Ctx) error {
	banners, err := h.bannerService.GetAllBanners(c.Context())
	if err != nil {
		return Error(c, "Failed to retrieve banners", err.Error())
	}

	return Success(c, banners)
}
func (h *BannerHandler) GetPublisherBanners(c *fiber.Ctx) error {
	banners, err := h.bannerService.GetAllPublisherBanners(c.Context())
	if err != nil {
		return Error(c, "Failed to retrieve banners", err.Error())
	}

	return Success(c, fiber.Map{"items": banners})
}

func (h *BannerHandler) GetAllCreatives(c *fiber.Ctx) error {
	bannerCreatives, err := h.bannerService.GetAllCreatives(c.Context())
	if err != nil {
		return Error(c, "Failed to retrieve banner creatives", err.Error())
	}

	return Success(c, fiber.Map{"items": bannerCreatives})
}
func (h *BannerHandler) CreateBanner(c *fiber.Ctx) error {
	// Parse request body
	req := new(dto.CreateBannerRequest)
	if err := c.BodyParser(req); err != nil {
		return Error(c, "Invalid request body", err.Error())
	}

	// Create banner
	banner, err := h.bannerService.CreateBanner(c.Context(), req)
	if err != nil {
		return Error(c, "Failed to create banner", err.Error())
	}

	return Success(c, banner)
}

func (h *BannerHandler) GetBanner(c *fiber.Ctx) error {
	// Parse banner ID
	bannerID, err := c.ParamsInt("id")
	if err != nil {
		return Error(c, "Invalid banner ID", err.Error())
	}

	// Fetch banner
	banner, err := h.bannerService.GetBannerByID(c.Context(), int64(bannerID))
	if err != nil {
		return Error(c, "Failed to retrieve banner", err.Error())
	}

	return Success(c, banner)
}

func (h *BannerHandler) CreateBannerCreative(c *fiber.Ctx) error {
	// Parse request body
	req := new(dto.CreateBannerCreativeRequest)
	if err := c.BodyParser(req); err != nil {
		return Error(c, "Invalid request body", err.Error())
	}

	// Create banner creative
	creative, err := h.bannerService.CreateBannerCreative(c.Context(), req)
	if err != nil {
		return Error(c, "Failed to create banner creative", err.Error())
	}

	return Success(c, creative)
}

func (h *BannerHandler) GetBannerCreatives(c *fiber.Ctx) error {
	// Parse banner ID
	bannerID, err := c.ParamsInt("id")
	if err != nil {
		return Error(c, "Invalid banner ID", err.Error())
	}

	// Fetch banner creatives
	creatives, err := h.bannerService.GetBannerCreativesByBannerID(c.Context(), int64(bannerID))
	if err != nil {
		return Error(c, "Failed to retrieve banner creatives", err.Error())
	}

	return Success(c, creatives)
}

// Add new handler methods
func (h *BannerHandler) AssignCreative(c *fiber.Ctx) error {
	req := new(dto.AssignCreativeRequest)
	if err := c.BodyParser(req); err != nil {
		return Error(c, "Invalid request body", err.Error())
	}

	err := h.bannerService.AssignCreativeToBanner(c.Context(), req)
	if err != nil {
		return Error(c, "Failed to assign creative to banner", err.Error())
	}

	return Success(c, fiber.Map{"message": "Creative assigned successfully"})
}

func (h *BannerHandler) UpdateCreativeOrder(c *fiber.Ctx) error {
	bannerID, err := c.ParamsInt("id")
	if err != nil {
		return Error(c, "Invalid banner ID", err.Error())
	}

	var orders []dto.CreativeOrder
	if err := c.BodyParser(&orders); err != nil {
		return Error(c, "Invalid request body", err.Error())
	}

	err = h.bannerService.UpdateBannerCreativeOrder(c.Context(), int64(bannerID), orders)
	if err != nil {
		return Error(c, "Failed to update creative order", err.Error())
	}

	return Success(c, fiber.Map{"message": "Creative order updated successfully"})
}

func (h *BannerHandler) RemoveCreative(c *fiber.Ctx) error {
	bannerID, err := c.ParamsInt("banner_id")
	if err != nil {
		return Error(c, "Invalid banner ID", err.Error())
	}

	creativeID, err := c.ParamsInt("creative_id")
	if err != nil {
		return Error(c, "Invalid creative ID", err.Error())
	}

	err = h.bannerService.RemoveCreativeFromBanner(c.Context(), int64(bannerID), int64(creativeID))
	if err != nil {
		return Error(c, "Failed to remove creative from banner", err.Error())
	}

	return Success(c, fiber.Map{"message": "Creative removed successfully"})
}

func (h *BannerHandler) SetPrimaryCreative(c *fiber.Ctx) error {
	bannerID, err := c.ParamsInt("banner_id")
	if err != nil {
		return Error(c, "Invalid banner ID", err.Error())
	}

	creativeID, err := c.ParamsInt("creative_id")
	if err != nil {
		return Error(c, "Invalid creative ID", err.Error())
	}

	err = h.bannerService.SetPrimaryCreative(c.Context(), int64(bannerID), int64(creativeID))
	if err != nil {
		return Error(c, "Failed to set primary creative", err.Error())
	}

	return Success(c, fiber.Map{"message": "Primary creative set successfully"})
}
