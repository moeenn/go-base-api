package config

import (
	"errors"
	"os"
)

type AuthConfig struct {
	JwtSecret               string
	LoginTokenExpiryMinutes uint
}

func loadAuthConfig() (*AuthConfig, error) {
	config := &AuthConfig{}
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return config, errors.New("failed to read JWT_SECRET from environment")
	}

	config.JwtSecret = secret
	config.LoginTokenExpiryMinutes = 60 * 24

	return config, nil
}
