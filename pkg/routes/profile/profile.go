package profile

import (
	"github.com/gofiber/fiber/v2"
)

func ProfileHandler(c *fiber.Ctx) error {
	headers := c.Request().Header.Header()

	return c.JSON(fiber.Map{
		"message": "You have reached protected route",
		"headers": headers,
	})
}
