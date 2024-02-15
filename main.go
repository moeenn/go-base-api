package main

import (
	"fmt"
	"os"
	"web/internal/config"
	"web/internal/module/auth"
	"web/internal/module/dashboard"
	"web/internal/server"
)

func main() {
	config, err := config.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "environment error: %s\n", err.Error())
		os.Exit(1)
	}

	e := server.New()
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
