package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/xvbnm48/rest-todo/database"
	"github.com/xvbnm48/rest-todo/models"
)

// @Summary Get all todos.
// @Description fetch every todo available.
// @Tags todos
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Todo
// @Router /todos [get]
func HandleAllTodos(c *fiber.Ctx) error {
	// fetch all todos
	db, err := database.GetDB()
	if err != nil {
		log.Println("error open db:", err)
		return c.Status(500).JSON(fiber.Map{
			"internal server error": err.Error(),
		})
	}

	//defer db.Close()

	rows, err := db.Query("SELECT * FROM todo")
	if err != nil {
		log.Println("error exec query:", err)
		return c.Status(500).JSON(fiber.Map{
			"internal server error": err.Error(),
		})
	}
	defer rows.Close()
	var todos []models.Todo

	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(
			&todo.ID, &todo.Title, &todo.Completed, &todo.Description, &todo.Date); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"internal server error": err.Error(),
			})
		}

		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":                http.StatusText(fiber.StatusInternalServerError),
			"code":                  fiber.StatusInternalServerError,
			"internal server error": err.Error(),
		})
	}

	log.Println("success get todos:", todos)
	// return todos
	return c.Status(200).JSON(fiber.Map{
		"status": http.StatusText(fiber.StatusOK),
		"code":   fiber.StatusOK,
		"data":   todos,
	})
}

// @Summary Create a todo.
// @Description create a single todo.
// @Tags todos
// @Accept json
// @Param todo body CreateTodoDTO true "Todo to create"
// @Produce json
// @Success 200 {object} CreateTodoResDTO
// @Router /todos [post]
func HandleCreateTodo(c *fiber.Ctx) error {
	// get the todo from the request body
	nTodo := new(models.CreateTodoDTO)
	//var nTodo *models.CreateTodoDTO
	if err := c.BodyParser(nTodo); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":            http.StatusText(fiber.StatusInternalServerError),
			"code":              fiber.StatusInternalServerError,
			"error body parser": err.Error(),
		})

	}
	db, err := database.GetDB()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":        http.StatusText(fiber.StatusInternalServerError),
			"code":          fiber.StatusInternalServerError,
			"error open DB": err.Error(),
		})
	}

	dateNow := time.Now()

	//defer db.Close()

	res, err := db.Exec("INSERT INTO todo (title, completed, description, date) VALUES (?,?,?,?)", nTodo.Title, nTodo.Completed, nTodo.Description, dateNow)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":          http.StatusText(fiber.StatusInternalServerError),
			"code":            fiber.StatusInternalServerError,
			"error open exec": err.Error(),
		})
	}

	insertedId, err := res.LastInsertId()

	log.Printf("success insert data in id %d", insertedId)
	resp := models.CreateTodoResDTO{
		InsertedId: insertedId,
	}
	// return the inserted todo
	return c.Status(200).JSON(fiber.Map{
		"status":   fiber.StatusOK,
		"code":     fiber.StatusOK,
		"message":  "success",
		"response": resp,
	})
}

// TODO MAKE A UDPATE

func HandleUpdateTodo(c *fiber.Ctx) error {
	nTodo := new(models.UpdateTodoDTO)
	fmt.Println("isi dari ntdoo:", nTodo)

	if err := c.BodyParser(nTodo); err != nil {
		return c.Status(501).JSON(fiber.Map{
			"status":            http.StatusText(fiber.StatusInternalServerError),
			"code":              fiber.StatusInternalServerError,
			"error body parser": err.Error(),
		})
	}
	db, err := database.GetDB()
	if err != nil {
		return err
	}

	_, err = db.Exec(` update todo set title = ? , completed = ? , description = ? , date = ?  where id = ? `, nTodo.Title, nTodo.Completed, nTodo.Description, time.Now(), nTodo.ID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":     http.StatusText(fiber.StatusInternalServerError),
			"code":       fiber.StatusInternalServerError,
			"error exec": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":   fiber.StatusOK,
		"code":     fiber.StatusOK,
		"message":  "success",
		"response": "success",
	})
}
