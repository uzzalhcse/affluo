// internal/middleware/cors.go
package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func ConfigureCORS() fiber.Handler {
	return cors.New(cors.Config{
		// Specify exact origins instead of wildcard
		AllowOrigins:     "http://localhost:3000,http://192.168.0.109:3000",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
		MaxAge:           86400,
	})
}
