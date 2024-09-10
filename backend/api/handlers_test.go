package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/hadyrashwan/golang-for-node-devs/dboperations"
)

// TestGetTodos is an example of how you can test your handler with sqlmock
func TestGetTodos(t *testing.T) {
	// Create a Fiber app for testing
	app := fiber.New()

	// Create a mock database connection and mock object
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Create a new SQLDB wrapper with the mock DB
	sqlDB := &dboperations.SQLDB{DB: db}

	// Create a new handler with the mock database
	handlers := NewTodoHandlers(sqlDB)
	app.Get("/api/todos", handlers.GetApi)

	// Define the expected query and the mocked rows
	rows := sqlmock.NewRows([]string{"id", "body", "completed"}).
		AddRow(1, "Test todo", false)

	// Expect the query to be called
	mock.ExpectQuery("SELECT \\* FROM todos").WillReturnRows(rows)

	// Create a new HTTP request for testing
	req := httptest.NewRequest("GET", "/api/todos", nil)
	req.Header.Set("Content-Type", "application/json")

	// Test the request
	resp, err := app.Test(req)

	// Ensure no error occurred during the test
	assert.NoError(t, err)

	// Check if the status code is 200 (OK)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Decode the JSON response
	var result map[string][]Todo
	err = json.NewDecoder(resp.Body).Decode(&result)
	assert.NoError(t, err)

	// Check the response data
	assert.Len(t, result["todos"], 1)
	assert.Equal(t, "Test todo", result["todos"][0].Body)

	// Ensure that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostTodo(t *testing.T) {
	// Create a Fiber app for testing
	app := fiber.New()

	// Create a mock database connection and mock object
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Create a new SQLDB wrapper with the mock DB
	sqlDB := &dboperations.SQLDB{DB: db}

	// Create a new handler with the mock database
	handlers := NewTodoHandlers(sqlDB)
	app.Post("/api/todos", handlers.PostApi)

	// Define the expected exec query for inserting a new todo
	mock.ExpectExec("INSERT INTO todos").
		WithArgs("New todo", false).
		WillReturnResult(sqlmock.NewResult(1, 1)) // ID of new todo is 1, and 1 row affected

	// Create the JSON body for the POST request
	todoJSON := `{"body":"New todo","completed":false}`
	req := httptest.NewRequest("POST", "/api/todos", bytes.NewBufferString(todoJSON))
	req.Header.Set("Content-Type", "application/json")

	// Test the request
	resp, err := app.Test(req)
	assert.NoError(t, err)

	// Check if the status code is 201 (Created)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// Decode the JSON response
	var result Todo
	err = json.NewDecoder(resp.Body).Decode(&result)
	assert.NoError(t, err)

	// Check the response data
	assert.Equal(t, "New todo", result.Body)
	assert.Equal(t, 1, result.ID) // The ID should be 1 as we mocked it

	// Ensure that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPatchTodo(t *testing.T) {
	// Create a Fiber app for testing
	app := fiber.New()

	// Create a mock database connection and mock object
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Create a new SQLDB wrapper with the mock DB
	sqlDB := &dboperations.SQLDB{DB: db}

	// Create a new handler with the mock database
	handlers := NewTodoHandlers(sqlDB)
	app.Patch("/api/todos/:id", handlers.PatchApi)

	// Mock the UPDATE query
	mock.ExpectExec("UPDATE todos SET completed = ? WHERE id = ?").
		WithArgs(true, "1").
		WillReturnResult(sqlmock.NewResult(0, 1)) // 1 row affected

	// Mock the SELECT query to return the updated todo
	rows := sqlmock.NewRows([]string{"id", "body", "completed"}).
		AddRow(1, "Updated todo", true)
	mock.ExpectQuery("SELECT \\* FROM todos WHERE id = ?").
		WithArgs("1").
		WillReturnRows(rows)

	// Create the PATCH request
	req := httptest.NewRequest("PATCH", "/api/todos/1", nil)
	req.Header.Set("Content-Type", "application/json")

	// Test the request
	resp, err := app.Test(req)
	assert.NoError(t, err)

	// Check if the status code is 200 (OK)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Decode the JSON response
	var result Todo
	err = json.NewDecoder(resp.Body).Decode(&result)
	assert.NoError(t, err)

	// Check the response data
	assert.Equal(t, "Updated todo", result.Body)
	assert.True(t, result.Completed)

	// Ensure that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteTodo(t *testing.T) {
	// Create a Fiber app for testing
	app := fiber.New()

	// Create a mock database connection and mock object
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Create a new SQLDB wrapper with the mock DB
	sqlDB := &dboperations.SQLDB{DB: db}

	// Create a new handler with the mock database
	handlers := NewTodoHandlers(sqlDB)
	app.Delete("/api/todos/:id", handlers.DeleteApi)

	// Mock the DELETE query
	mock.ExpectExec("DELETE FROM todos WHERE id = ?").
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(0, 1)) // 1 row affected

	// Create the DELETE request
	req := httptest.NewRequest("DELETE", "/api/todos/1", nil)
	req.Header.Set("Content-Type", "application/json")

	// Test the request
	resp, err := app.Test(req)
	assert.NoError(t, err)

	// Check if the status code is 200 (OK)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Decode the JSON response
	var result map[string]string
	err = json.NewDecoder(resp.Body).Decode(&result)
	assert.NoError(t, err)

	// Check the response data
	assert.Equal(t, "Todo deleted", result["message"])

	// Ensure that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

