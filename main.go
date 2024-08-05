package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/hadyrashwan/golang-for-node-devs/dboperations"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int    `json:"id"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}

type Exec struct {
	LastInsertId int64 `json:"lastInsertId"`
	RowsAffected int64 `json:"rowsAffected"`
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	DB_URL := os.Getenv("DB_URL")
	DB_TOKEN := os.Getenv("DB_TOKEN")

	url := fmt.Sprintf("%s?authToken=%s", DB_URL, DB_TOKEN)

	db, err := dboperations.Connect(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}

	app := fiber.New()

	todos := []Todo{}

	app.Get("/api/todos/", func(c *fiber.Ctx) error {
		todos, err = dboperations.Query_helper[Todo](db, "SELECT * FROM todos")
		if err != nil {
			return err
		}
		return c.Status(200).JSON(fiber.Map{
			"todos": todos,
		})
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
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
		result, err := dboperations.Exec_helper[Todo](db, query, todo.Body, todo.Completed)
		if err != nil {
			return err
		}
		todo.ID = int(result.LastInsertId)
		return c.Status(201).JSON(todo)
	})

	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		query := `
		UPDATE todos
		SET completed = ?
		WHERE id = ?`
		result, err := dboperations.Exec_helper[Todo](db, query, true, id)
		if err != nil {
			return err
		}
		if result.RowsAffected == 0 {
			return c.Status(400).JSON(fiber.Map{
				"error": "Todo not found",
			})
		}
		todos, err = dboperations.Query_helper[Todo](db, "SELECT * FROM todos WHERE id = ?", id)
		if err != nil {
			return err
		}
		return c.Status(200).JSON(todos[0])

	})

	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		query := `
		DELETE FROM todos WHERE id = ?`
		result, err := dboperations.Exec_helper[Todo](db, query, id)
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

	})

	log.Fatal(app.Listen(":" + PORT))
}
