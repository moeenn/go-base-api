package env

import (
	"os"
	"strings"
)

func Env(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("environment variable " + key + " not set")
	}

	return strings.TrimSpace(value)
}
