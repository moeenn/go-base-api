package home

import (
	"github.com/gofiber/fiber/v2"
)

func HomeHandler(c *fiber.Ctx) error {
	res := fiber.Map{
		"message": "welcome to home",
	}

	return c.JSON(res)
}
