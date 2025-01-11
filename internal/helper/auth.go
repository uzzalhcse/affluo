package helper

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func GetUserIDFromContext(c *fiber.Ctx) (int64, error) {
	// Convert user_id to int64
	userIDValue := c.Locals("user_id")
	var userID int64

	switch v := userIDValue.(type) {
	case float64:
		userID = int64(v)
	case int64:
		userID = v
	default:
		return 0, errors.New("invalid user_id type")
	}

	return userID, nil

}
