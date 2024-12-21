package middleware

import (
	"affluo/internal/service"
	"github.com/gofiber/fiber/v2"
	"time"
)

func AntiFraud(antiFraud *service.AntiFraudService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Execute the next handler first to ensure route parameters are parsed
		err := c.Next()
		if err != nil {
			return err
		}

		// Now we can safely access route parameters
		id := c.Params("id")
		if id == "" {
			// Skip validation if no ID parameter (might be a different route)
			return nil
		}

		bannerID, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid banner ID: " + err.Error(),
			})
		}

		// Determine event type based on the route
		var eventType string
		switch {
		case c.Path() == "/api/tracking/impression/"+id:
			eventType = "impression"
		case c.Path() == "/api/tracking/click/"+id:
			eventType = "click"
		case c.Path() == "/api/tracking/lead/"+id:
			eventType = "lead"
		default:
			eventType = "unknown"
		}

		activity := service.ActivityLog{
			IP:          c.IP(),
			UserAgent:   c.Get("User-Agent"),
			Timestamp:   time.Now(),
			EventType:   eventType,
			BannerID:    int64(bannerID),
			Fingerprint: c.Get("X-Device-Fingerprint"),
			Metadata: map[string]interface{}{
				"referer": c.Get("Referer"),
				"path":    c.Path(),
				"method":  c.Method(),
			},
		}

		if err := antiFraud.ValidateAndLogActivity(activity); err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Activity blocked due to suspicious behavior",
			})
		}

		return nil
	}
}
