package main

import (
	"fmt"
	"os"
	"web/internal/config"
	"web/internal/helpers/validator"
	"web/internal/module/auth"
	"web/internal/module/dashboard"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config, err := config.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "environment error: %s\n", err.Error())
		os.Exit(1)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = validator.New()
	authMiddleware := auth.AuthMiddleware(config.Auth.JWTSecret)

	api := e.Group("/api")
	{
		api.POST("/login", auth.LoginHandler(config.Auth.JWTSecret))
		api.GET("/protected", authMiddleware(dashboard.ProtectedHandler))
	}

	if err := e.Start(config.Server.Address); err != nil {
		e.Logger.Fatal(err)
	}
}
