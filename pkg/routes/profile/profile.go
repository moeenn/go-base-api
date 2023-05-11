package profile

import (
	"app/pkg/services/auth"

	"github.com/gofiber/fiber/v2"
)

func ProfileHandler(c *fiber.Ctx) error {
	user := auth.User(c)

	return c.JSON(fiber.Map{
		"message": "You have reached protected route",
		"user":    user,
	})
}
