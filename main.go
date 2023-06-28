package main

import (
	"app/config"
	"app/pkg/jwt"
	"app/pkg/middleware"
	"fmt"
	"net/http"
	"os"

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
		api.GET("/login", LoginHandler(conf.AuthConfig))
		api.GET("/profile", middleware.ValidateToken(conf.AuthConfig.JwtSecret), ProfileHandler)
	}

	router.Run(conf.ServerConfig.Port)
}

type ProfileResponse struct {
	Message string `json:"message"`
	UserId  string `json:"userId"`
}

func ProfileHandler(c *gin.Context) {
	userId := c.MustGet("userId").(string)

	res := ProfileResponse{
		Message: "Hello world",
		UserId:  userId,
	}

	c.JSON(http.StatusOK, res)
}

type LoginResponse struct {
	Token  string `json:"token"`
	Expiry int64  `json:"expiry"`
}

func LoginHandler(authConfig *config.AuthConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		payload := jwt.JWTPayload{
			UserId:    "300-400-500",
			UserRoles: []string{"ADMIN"},
		}

		token, err := jwt.GenerateToken(
			authConfig.JwtSecret,
			authConfig.LoginTokenExpiryMinutes,
			payload,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		res := LoginResponse{
			Token:  token.Token,
			Expiry: token.Expiry,
		}

		c.JSON(http.StatusOK, res)
	}
}
