package main

import (

	"github.com/gofiber/fiber/v2"
	"github.com/hadyrashwan/golang-for-node-devs/dboperations"
)

func NewTodoHandlers(db *dboperations.SQLDB) *TodoHandlers {
	return &TodoHandlers{DB: db}
}

type TodoHandlers struct {
	DB *dboperations.SQLDB
}


type Todo struct {
	ID        int    `json:"id"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}

// Handlers 
func (h *TodoHandlers)  GetApi(c *fiber.Ctx) error {
		
	todos, err := dboperations.Query_helper[Todo](h.DB, "SELECT * FROM todos")
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{
		"todos": todos,
	})
}

func (h *TodoHandlers) PostApi(c *fiber.Ctx) error {
	todo := &Todo{}
	if err := c.BodyParser(todo); err != nil {
		return err
	}
	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Todo body cannot be empty",
		})
	}
	query := "INSERT INTO todos (body, completed) VALUES (?, ?)"
	result, err := dboperations.Exec_helper[Todo](h.DB, query, todo.Body, todo.Completed)
	if err != nil {
		return err
	}
	todo.ID = int(result.LastInsertId)
	return c.Status(201).JSON(todo)
}

func (h *TodoHandlers)  PatchApi(c *fiber.Ctx) error {
	id := c.Params("id")

	query := `
	UPDATE todos
	SET completed = ?
	WHERE id = ?`
	result, err := dboperations.Exec_helper[Todo](h.DB, query, true, id)
	if err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}
	todos, err := dboperations.Query_helper[Todo](h.DB, "SELECT * FROM todos WHERE id = ?", id)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(todos[0])

}

func (h *TodoHandlers)  DeleteApi(c *fiber.Ctx) error {
	id := c.Params("id")

	query := `
	DELETE FROM todos WHERE id = ?`
	result, err := dboperations.Exec_helper[Todo](h.DB, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Todo deleted",
	})
}
