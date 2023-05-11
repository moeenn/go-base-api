package auth

import (
	"app/pkg/config"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type LoginTokenPayload struct {
	Id   string
	Role string
}

func GenerateLoginToken(payload LoginTokenPayload) (string, error) {
	claims := jwt.MapClaims{
		"id":   payload.Id,
		"role": payload.Role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.EnvConfig.Secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

// TODO: test
func User(c *fiber.Ctx) LoginTokenPayload {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	result := LoginTokenPayload{
		Id:   claims["id"].(string),
		Role: claims["role"].(string),
	}

	return result
}
