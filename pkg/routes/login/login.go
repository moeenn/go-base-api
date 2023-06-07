package login

import (
	"app/pkg/helpers/jwt"
	"app/pkg/validator"

	"app/pkg/services/auth"

	"app/pkg/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   jwt.TokenResult
}

func LoginHandler(c *fiber.Ctx) error {
	body := new(LoginRequest)
	if err := c.BodyParser(body); err != nil {
		return helpers.HTTPError(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.ValidateStruct(body); err != nil {
		return helpers.ValidationError(c, err)
	}

	if body.Email != "admin@site.com" && body.Password != "abc123123" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token, err := auth.GenerateLoginToken(jwt.JWTPayload{
		UserId:   uuid.New().String(),
		UserRole: "ADMIN",
	})

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	res := LoginResponse{
		Message: "login successful",
		Token:   *token,
	}

	return c.JSON(res)
}
