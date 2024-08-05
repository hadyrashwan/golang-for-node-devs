package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
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

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}

	app := fiber.New()

	// create todos example
	todos := []Todo{}

	app.Get("/api/todos/", func(c *fiber.Ctx) error {
		todos, err = db_query_helper[Todo](db, "SELECT * FROM todos")
		if err != nil {
			return err
		}
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
		query := "INSERT INTO todos (body, completed) VALUES (?, ?)"

		result, err := db_exec_helper[Todo](db, query, todo.Body, todo.Completed)
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
		result, err := db_exec_helper[Todo](db, query,true, id)
		if err != nil {
			return err
		}
		if result.RowsAffected == 0{
			return c.Status(400).JSON(fiber.Map{
				"error": "Todo not found",
			})
		}
		todos, err = db_query_helper[Todo](db, "SELECT * FROM todos WHERE id = ?", id)
		if err != nil {
			return err
		}
		return c.Status(200).JSON(todos[0])
		
	})

	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		query := `
		DELETE FROM todos WHERE id = ?`
		result, err := db_exec_helper[Todo](db, query, id)
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

func db_query_helper[T any](db *sql.DB, query string, args ...interface{}) ([]T, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	// rows.Close()

	var objects []T

	for rows.Next() {
		var object T

		s := reflect.ValueOf(&object).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err := rows.Scan(columns...)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}

		objects = append(objects, object)
		// fmt.Println(user.ID, user.Name)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
	}
	return objects, nil
}

func db_exec_helper[T any](db *sql.DB, query string, args ...interface{}) (Exec, error) {
	exec, err := db.Exec(query, args...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	// rows.Close()

	id, err := exec.LastInsertId()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	rows, err := exec.RowsAffected()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	result := Exec{
		LastInsertId: id,
		RowsAffected: rows,
	}
	exec.LastInsertId()
	exec.RowsAffected()

	return result, nil
}
