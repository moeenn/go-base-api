package login

import (
	"app/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	Message string `json:"message"`
}

func LoginHandler(c *fiber.Ctx) error {
	body := new(LoginRequest)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := validator.ValidateStruct(body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err,
		})
	}

	res := LoginResponse{
		Message: "login successful",
	}

	return c.JSON(res)
}
