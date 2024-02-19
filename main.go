package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
	"web/internal/config"
	"web/internal/module/auth"
	"web/internal/module/dashboard"
	"web/internal/server"
)

type Controllers struct {
	AuthController      *auth.AuthController
	DashboardController *dashboard.DashboardController
}

type InitControllersArgs struct {
	Config         *config.GlobalConfig
	AuthMiddleware *echo.MiddlewareFunc
}

func InitControllers(args InitControllersArgs) *Controllers {
	return &Controllers{
		AuthController: &auth.AuthController{
			Config: args.Config,
		},
		DashboardController: &dashboard.DashboardController{
			AuthMiddleware: args.AuthMiddleware,
		},
	}
}

func main() {
	config, err := config.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "environment error: %s\n", err.Error())
		os.Exit(1)
	}

	// initialize global dependencies
	e := server.New()
	authMiddleware := auth.AuthMiddleware(config.Auth.JWTSecret)

	controllers := InitControllers(InitControllersArgs{
		Config:         config,
		AuthMiddleware: &authMiddleware,
	})

	// register all controllers here
	controllers.AuthController.RegisterRoutes(e)
	controllers.DashboardController.RegisterRoutes(e)

	// start the server
	if err := e.Start(config.Server.Address); err != nil {
		e.Logger.Fatal(err)
	}
}
