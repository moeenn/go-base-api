package router

import (
	"app/pkg/routes/home"
	"app/pkg/routes/login"
	"app/pkg/routes/profile"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")
	{
		api.Get("/", home.HomeHandler)
		api.Post("/login", login.LoginHandler)
		api.Get("/profile", profile.ProfileHandler)
	}
}
