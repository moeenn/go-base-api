package controller

import (
	"net/http"
	"web/internal/helpers/jwthelpers"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func LoginHandler(jwtSecret string) func(echo.Context) error {
	handler := func(c echo.Context) error {
		body := &LoginRequest{}
		if err := c.Bind(body); err != nil {
			return err
		}

		if err := c.Validate(body); err != nil {
			return err
		}

		// TODO: check from DB
		if body.Email == "admin@site.com" && body.Password == "abc123123123" {
			id := uuid.New().String()
			token, err := jwthelpers.NewJWT(id, "ADMIN", jwtSecret)
			if err != nil {
				return echo.ErrUnauthorized
			}
			return c.JSON(http.StatusOK, token)
		}

		return echo.ErrUnauthorized
	}

	return handler
}
