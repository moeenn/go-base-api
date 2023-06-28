package auth

import (
	"app/config"

	"app/pkg/jwt"
)

func GenerateLoginToken(authConfig config.AuthConfig, payload jwt.JWTPayload) (*jwt.TokenResult, error) {
	result, err := jwt.GenerateToken(
		authConfig.JwtSecret,
		authConfig.LoginTokenExpiryMinutes,
		payload,
	)

	if err != nil {
		return &jwt.TokenResult{}, err
	}

	return result, nil
}

// TODO: implement
// func User(c *fiber.Ctx) (*jwt.JWTPayload, error) {
// 	token := c.Locals("user").(*jwt.Token)
// 	payload, err := jwt.ValidateToken(config.EnvConfig.Secret, token.Raw)
// 	return payload, err
// }
