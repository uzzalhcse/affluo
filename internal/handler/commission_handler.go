// handler/commission_handler.go
package handler

import (
	"affluo/internal/dto"
	"affluo/internal/service"
	"github.com/gofiber/fiber/v2"
)

type CommissionHandler struct {
	commissionService *service.CommissionService
}

func NewCommissionHandler(commissionService *service.CommissionService) *CommissionHandler {
	return &CommissionHandler{commissionService: commissionService}
}
func (h *CommissionHandler) CreateCommissionPlan(c *fiber.Ctx) error {
	req := new(dto.CreateCommissionPlanRequest)
	if err := c.BodyParser(req); err != nil {
		return Error(c, "Invalid request body", err.Error())
	}

	plan, err := h.commissionService.CreateCommissionPlan(c.Context(), req)
	if err != nil {
		return Error(c, "Failed to create commission plan", err.Error())
	}

	return Success(c, fiber.Map{
		"message": "Commission plan created successfully",
		"plan":    plan,
	})
}

func (h *CommissionHandler) ListCommissionPlans(c *fiber.Ctx) error {
	plans, err := h.commissionService.GetAllCommissionPlans(c.Context())
	if err != nil {
		return Error(c, "Failed to fetch commission plans", err.Error())
	}

	return Success(c, fiber.Map{
		"plans": plans,
	})
}

func (h *CommissionHandler) GetCommissionPlan(c *fiber.Ctx) error {
	planID, err := c.ParamsInt("id")
	if err != nil {
		return Error(c, "Invalid plan ID", err.Error())
	}

	plan, err := h.commissionService.GetCommissionPlanByID(c.Context(), int64(planID))
	if err != nil {
		return Error(c, "Failed to fetch commission plan", err.Error())
	}

	return Success(c, fiber.Map{
		"plan": plan,
	})
}

func (h *CommissionHandler) UpdateCommissionPlan(c *fiber.Ctx) error {
	planID, err := c.ParamsInt("id")
	if err != nil {
		return Error(c, "Invalid plan ID", err.Error())
	}

	req := new(dto.CreateCommissionPlanRequest)
	if err := c.BodyParser(req); err != nil {
		return Error(c, "Invalid request body", err.Error())
	}

	plan, err := h.commissionService.UpdateCommissionPlan(c.Context(), int64(planID), req)
	if err != nil {
		return Error(c, "Failed to update commission plan", err.Error())
	}

	return Success(c, fiber.Map{
		"message": "Commission plan updated successfully",
		"plan":    plan,
	})
}

func (h *CommissionHandler) AssignToPublisher(c *fiber.Ctx) error {
	planID, err := c.ParamsInt("id")
	if err != nil {
		return Error(c, "Invalid plan ID", err.Error())
	}

	publisherID, err := c.ParamsInt("publisherId")
	if err != nil {
		return Error(c, "Invalid publisher ID", err.Error())
	}

	err = h.commissionService.AssignCommissionPlanToPublisher(c.Context(), int64(planID), int64(publisherID))
	if err != nil {
		return Error(c, "Failed to assign commission plan", err.Error())
	}

	return Success(c, fiber.Map{
		"message": "Commission plan assigned successfully",
	})
}

func (h *CommissionHandler) GetPublisherCommission(c *fiber.Ctx) error {
	publisherID, err := c.ParamsInt("publisherId")
	if err != nil {
		return Error(c, "Invalid publisher ID", err.Error())
	}

	plan, err := h.commissionService.GetPublisherCommission(c.Context(), int64(publisherID))
	if err != nil {
		return Error(c, "Failed to fetch publisher's commission plan", err.Error())
	}

	return Success(c, fiber.Map{
		"plan": plan,
	})
}
