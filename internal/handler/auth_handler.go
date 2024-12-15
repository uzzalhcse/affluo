// internal/handler/auth_handler.go
package handler

import (
	"affluo/internal/dto"
	"affluo/internal/service"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return Error(c, err.Error())
	}

	user, err := h.authService.Register(c.Context(), req)
	if err != nil {
		return Error(c, err.Error())
	}

	return Success(c, fiber.Map{
		"user_id": user.ID,
		"email":   user.Email,
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return Error(c, "Invalid request body", err.Error())
	}

	authResponse, err := h.authService.Login(c.Context(), req)
	if err != nil {
		return Unauthorized(c, err.Error())
	}

	return Success(c, authResponse)
}

func (h *AuthHandler) Profile(c *fiber.Ctx) error {
	// Convert user_id to int64
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

	user, err := h.authService.GetUserFromContext(c.Context(), userID)
	if err != nil {
		return Unauthorized(c, "User not authenticated", err)
	}
	return Success(c, user)
}
func (h *AuthHandler) InitiatePasswordReset(c *fiber.Ctx) error {
	var req dto.PasswordResetRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := h.authService.InitiatePasswordReset(c.Context(), req.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Password reset token sent",
	})
}

func (h *AuthHandler) ResetPassword(c *fiber.Ctx) error {
	var req dto.PasswordResetConfirmRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := h.authService.ResetPassword(c.Context(), req.Token, req.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Password reset successfully",
	})
}
