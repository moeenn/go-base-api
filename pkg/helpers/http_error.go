package helpers

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Details any    `json:"details"`
}

func HTTPError(c *fiber.Ctx, status int, message string) error {
	resp := ErrorResponse{
		Status:  status,
		Error:   message,
		Details: nil,
	}

	return c.Status(status).JSON(resp)
}

func ValidationError(c *fiber.Ctx, details any) error {
	status := fiber.StatusUnprocessableEntity

	resp := ErrorResponse{
		Status:  status,
		Error:   "Invalid data provided",
		Details: details,
	}

	return c.Status(status).JSON(resp)
}

func BadRequestError(c *fiber.Ctx, message string) error {
	status := fiber.StatusBadRequest

	resp := ErrorResponse{
		Status:  status,
		Error:   message,
		Details: nil,
	}

	return c.Status(status).JSON(resp)
}
