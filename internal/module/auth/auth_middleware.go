package auth

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"web/internal/helpers/jwthelpers"
)

func AuthMiddleware(jwtSecret string) echo.MiddlewareFunc {
	return echojwt.WithConfig(jwthelpers.NewJWTConfig(jwtSecret))
}
