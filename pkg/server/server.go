package server

import (
	"app/pkg/server/router"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	app := fiber.New()
	router.RegisterRoutes(app)

	return app
}
