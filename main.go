package main

import (
	"web/internal/config"
	"web/internal/controller"
	"web/internal/helpers/jwthelpers"
	"web/internal/helpers/validator"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := config.New()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	authMiddleware := echojwt.WithConfig(jwthelpers.NewJWTConfig(config.Auth.JWTSecret))

	e.Validator = validator.New()
	{
		e.POST("/login", controller.LoginHandler(config.Auth.JWTSecret))
		e.GET("/protected", authMiddleware(controller.ProtectedHandler))
	}

	if err := e.Start(config.Server.Address); err != nil {
		e.Logger.Fatal(err)
	}
}
