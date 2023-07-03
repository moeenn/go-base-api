package login

import (
	"app/config"
	"app/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
