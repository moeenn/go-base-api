package middleware

import (
	"app/pkg/jwt"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateToken(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := extractAuthHeader(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		jwtPayload, err := jwt.ValidateToken(jwtSecret, token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Set("userId", jwtPayload.UserId)
		c.Set("userRoles", jwtPayload.UserRoles)

		// TODO: request is not short circuited if we return from here
		// without calling Next(). Response is sent twice!

		c.Next()
	}
}

func extractAuthHeader(c *gin.Context) (string, error) {
	header := c.GetHeader("Authorization")
	if header == "" {
		return "", errors.New("'Authorization' header missing")
	}

	isBearer := strings.Contains(header, "Bearer ")
	if !isBearer {
		return "", errors.New("unsupported authorization scheme")
	}

	token := strings.ReplaceAll(header, "Bearer ", "")
	if token == "" {
		return "", errors.New("missing token value in 'Authorization' header")
	}

	return token, nil
}
