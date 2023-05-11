package server

import (
	"app/pkg/routes/home"
	"app/pkg/routes/login"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	app := fiber.New()
	{
		app.Get("/", home.HomeHandler)
		app.Post("/login", login.LoginHandler)
	}

	return app
}
