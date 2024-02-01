package config

import (
	"web/internal/helpers/env"
)

type ServerConfig struct {
	Address string
}

type AuthConfig struct {
	JWTSecret string
}

type GlobalConfig struct {
	Server ServerConfig
	Auth   AuthConfig
}

func New() (*GlobalConfig, error) {
	config := &GlobalConfig{
		Server: ServerConfig{
			Address: "0.0.0.0:5000",
		},
		Auth: AuthConfig{},
	}

	jwtSecret, err := env.Env("JWT_SECRET")
	if err != nil {
		return config, err
	}
	config.Auth.JWTSecret = jwtSecret

	return config, nil
}
