package models

import "time"

type Todo struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Completed   bool   `json:"completed"`
	Description string `json:"description"`
	Date        string `json:"date" `
}

type CreateTodoDTO struct {
	Title       string    `json:"title"`
	Completed   bool      `json:"completed"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type CreateTodoResDTO struct {
	InsertedId int64 `json:"inserted_id"`
}

type UpdateTodoDTO struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Completed   bool      `json:"completed"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
