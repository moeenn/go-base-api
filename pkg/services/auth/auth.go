package auth

import (
	"app/pkg/config"

	"app/pkg/helpers/jwt"
)

func GenerateLoginToken(payload jwt.JWTPayload) (*jwt.TokenResult, error) {
	result, err := jwt.GenerateToken(
		config.EnvConfig.Secret,
		config.SecurityConfig.LoginTokenExpiryMinutes,
		payload,
	)

	if err != nil {
		return &jwt.TokenResult{}, err
	}

	return result, nil
}

// func User(c *fiber.Ctx) (*jwt.JWTPayload, error) {
// 	token := c.Locals("user").(*jwt.Token)
// 	payload, err := jwt.ValidateToken(config.EnvConfig.Secret, token.Raw)
// 	return payload, err
// }
