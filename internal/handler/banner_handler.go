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

	return Success(c, dto.NewBannersResponse(banners))
}
func (h *BannerHandler) GetPublisherBanners(c *fiber.Ctx) error {
	banners, err := h.bannerService.GetAllPublisherBanners(c.Context())
	if err != nil {
		return Error(c, "Failed to retrieve banners", err.Error())
	}

	return Success(c, fiber.Map{"items": banners})
}

func (h *BannerHandler) GetAllBannerCreatives(c *fiber.Ctx) error {
	bannerCreatives, err := h.bannerService.GetAllBannerCreatives(c.Context())
	if err != nil {
		return Error(c, "Failed to retrieve banner creatives", err.Error())
	}

	return Success(c, dto.NewBannerCreativesResponse(bannerCreatives))
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
