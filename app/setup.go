package app

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/xvbnm48/rest-todo/config"
	"github.com/xvbnm48/rest-todo/database"
	"github.com/xvbnm48/rest-todo/router"
)

func SetupAndRunApp() error {
	// load env
	err := config.LoadENV()
	if err != nil {
		return err
	}

	// start database
	err = database.StartDB()
	if err != nil {
		return err
	}

	// defer closing database
	defer database.CloseDB()

	// create app
	app := fiber.New()

	// attach middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	// setup routes
	router.SetupRoutes(app)

	// attach swagger
	config.AddSwaggerRoutes(app)

	// get the port and start
	port := os.Getenv("PORT")
	app.Listen(":" + port)

	return nil
}
