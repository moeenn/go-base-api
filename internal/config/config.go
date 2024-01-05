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

func New() *GlobalConfig {
	return &GlobalConfig{
		Server: ServerConfig{
			Address: "0.0.0.0:5000",
		},
		Auth: AuthConfig{
			JWTSecret: env.Env("JWT_SECRET"),
		},
	}
}
