package home

import (
	"github.com/gofiber/fiber/v2"
)

type HomeResponse struct {
	Message string `json:"message"`
}

func HomeHandler(c *fiber.Ctx) error {
	res := HomeResponse{
		Message: "Hello from home",
	}

	return c.JSON(res)
}
