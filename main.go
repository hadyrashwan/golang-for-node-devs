package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
)

type Todo struct {
	ID        int    `json:"id"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}


func main() {
	

	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")


	app := fiber.New()

	// create todos example
	todos := []Todo{}


	app.Get("/api/todos/", func(c *fiber.Ctx) error {
		// return c.SendString("Hello, World!")
		return c.Status(200).JSON(fiber.Map{
			"todos": todos,
		})
	})

	// create a post endpoint for todos
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
		todo.ID = len(todos) + 1
		todos = append(todos, *todo)
		return c.Status(201).JSON(todo)
	})

	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		// todo := Todo{}
		// loop over todos and find hte matching by id
		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(400).JSON(fiber.Map{
			"error": "Todo not found",
		})
	})

	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				// delete matched item from array
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{
					"message": "Todo deleted",
				})
			}
			}
		return c.Status(400).JSON(fiber.Map{
			"error": "Todo not found",
		})
	})

	log.Fatal(app.Listen(":" + PORT))
}
