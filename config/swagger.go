package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "github.com/xvbnm48/rest-todo/docs"
)

func AddSwaggerRoutes(app *fiber.App) {
	// setup swagger
	app.Get("/swagger/*", swagger.HandlerDefault)
}
