package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Process request
		err := c.Next()

		// Log request details
		log.Printf(
			"%s | %d | %s | %s",
			c.Method(),
			c.Response().StatusCode(),
			c.Path(),
			time.Since(start),
		)

		return err
	}
}
