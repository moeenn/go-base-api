package server

import (
	"app/pkg/server/router"

	"app/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func New() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.EnvConfig.FrontendHost,
	}))

	router.RegisterRoutes(app)
	return app
}
