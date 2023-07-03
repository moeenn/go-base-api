package main

import (
	"app/config"
	"app/pkg/middleware"
	"fmt"
	"os"

	"app/routes/login"
	"app/routes/profile"

	"github.com/gin-gonic/gin"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err.Error())
		os.Exit(1)
	}

	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/login", login.LoginHandler(conf.AuthConfig))
		api.GET("/profile", middleware.ValidateToken(conf.AuthConfig.JwtSecret), profile.ProfileHandler)
	}

	router.Run(conf.ServerConfig.Port)
}
