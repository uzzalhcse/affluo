package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Recover() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":   "Internal Server Error",
					"message": "An unexpected error occurred",
					"stack":   r,
				})
			}
		}()

		return c.Next()
	}
}
