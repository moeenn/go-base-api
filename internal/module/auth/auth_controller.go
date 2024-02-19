package auth

import (
	"net/http"
	"web/internal/config"
	"web/internal/helpers/jwthelpers"
	"web/internal/helpers/responses"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type AuthController struct {
	Config *config.GlobalConfig
}

func (c *AuthController) RegisterRoutes(e *echo.Echo) {
	e.POST("/api/auth/login", c.loginHandler)
}

func (controller *AuthController) loginHandler(c echo.Context) error {
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
		token, err := jwthelpers.NewJWT(id, "ADMIN", controller.Config.Auth.JWTSecret)
		if err != nil {
			return echo.ErrUnauthorized
		}

		return c.JSON(http.StatusOK, responses.NewOkResponse(token))
	}

	return c.JSON(http.StatusUnauthorized, responses.NewErrorResponse(
		http.StatusUnauthorized,
		"Invalid email or password",
	))
}
