package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xvbnm48/rest-todo/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/health", handlers.HandleHealthCheck)

	// setup the todos group
	todos := app.Group("/todos")
	todos.Get("/", handlers.HandleAllTodos)
	todos.Post("/", handlers.HandleCreateTodo)
	todos.Post("/update", handlers.HandleUpdateTodo)
	//todos.Put("/:id", handlers.HandleUpdateTodo)
	//todos.Get("/:id", handlers.HandleGetOneTodo)
	//todos.Delete("/:id", handlers.HandleDeleteTodo)
}
