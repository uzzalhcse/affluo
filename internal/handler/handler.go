package handler

import "github.com/gofiber/fiber/v2"

// Success sends a successful JSON response
func Success(c *fiber.Ctx, data interface{}) error {
	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

// Error sends a JSON error response
func Error(c *fiber.Ctx, message string, errors ...interface{}) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"success": false,
		"message": message,
		"errors":  errors,
	})
}

func Unauthorized(c *fiber.Ctx, message string, errors ...error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"success": false,
		"message": message,
		"errors":  errors,
	})

}
