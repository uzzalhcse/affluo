// internal/middleware/auth.go
package middleware

import (
	"affluo/internal/service"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	authService *service.AuthService
}

func NewAuthMiddleware(authService *service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{authService: authService}
}

func (m *AuthMiddleware) Authenticate(requiredRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing authorization token",
			})
		}

		// Extract token (expecting "Bearer <token>")
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token format",
			})
		}

		token := parts[1]

		// Validate token
		_, claims, err := m.authService.ValidateJWT(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		// Check role if required roles are specified
		if len(requiredRoles) > 0 {
			userRole, ok := claims["role"].(string)
			if !ok {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"error": "Invalid user role",
				})
			}

			roleAllowed := false
			for _, role := range requiredRoles {
				if userRole == role {
					roleAllowed = true
					break
				}
			}

			if !roleAllowed {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"error": fmt.Sprintf("Access denied. Required role: %v", requiredRoles),
				})
			}
		}

		// Set user context for downstream use
		c.Locals("user_id", claims["user_id"])
		c.Locals("email", claims["email"])
		c.Locals("role", claims["role"])

		return c.Next()
	}
}
