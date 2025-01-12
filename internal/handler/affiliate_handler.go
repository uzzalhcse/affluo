// internal/handler/affiliate_handler.go
package handler

import (
	"affluo/internal/dto"
	"affluo/internal/helper"
	"affluo/internal/service"
	"github.com/gofiber/fiber/v2"
)

type AffiliateHandler struct {
	affiliateService *service.AffiliateService
}

func NewAffiliateHandler(affiliateService *service.AffiliateService) *AffiliateHandler {
	return &AffiliateHandler{affiliateService: affiliateService}
}

func (h *AffiliateHandler) CreateAffiliate(c *fiber.Ctx) error {
	req := new(dto.CreateAffiliateRequest)
	if err := c.BodyParser(req); err != nil {
		return Error(c, "Invalid request body", err.Error())
	}

	affiliate, err := h.affiliateService.CreateAffiliate(c.Context(), req)
	if err != nil {
		return Error(c, "Failed to create affiliate", err.Error())
	}

	return Success(c, fiber.Map{
		"message":   "Affiliate created successfully",
		"affiliate": affiliate,
	})
}
func (h *AffiliateHandler) GetAffiliate(c *fiber.Ctx) error {
	id, err := helper.GetUserIDFromContext(c)
	affiliates, err := h.affiliateService.GetUserAffiliates(c.Context(), id)
	if err != nil {
		return Error(c, "Failed to fetch user affiliates", err.Error())
	}

	return Success(c, fiber.Map{
		"affiliates": affiliates,
	})
}
