package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hadyrashwan/golang-for-node-devs/dboperations"
	"github.com/joho/godotenv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
)
var fiberLambda *fiberadapter.FiberLambda

var fiber_server *fiber.App

type Exec struct {
	LastInsertId int64 `json:"lastInsertId"`
	RowsAffected int64 `json:"rowsAffected"`
}



func init() {
	
	// for local dev will handle later
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("The .env is not loaded")
	}
	
	DB_URL := os.Getenv("DB_URL")
	DB_TOKEN := os.Getenv("DB_TOKEN")
	BASE_URL := os.Getenv("BACKEND_BASE_URL")
	PORT := os.Getenv("BACKEND_PORT")


	log.Println("DB_URL: ", DB_URL)
	log.Println("PORT: ", PORT)


	url := fmt.Sprintf("%s?authToken=%s", DB_URL, DB_TOKEN)

	db, err := dboperations.Connect(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}

	app := fiber.New()

	crosConfig := cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin,Content-Type,Accept",
	}
	app.Use(cors.New(crosConfig))


	handlers := NewTodoHandlers(db)


	app.Get(BASE_URL +"/api/todos/", handlers.GetApi)
	app.Post(BASE_URL +"/api/todos", handlers.PostApi)
	app.Patch(BASE_URL +"/api/todos/:id", handlers.PatchApi)
	app.Delete(BASE_URL +"/api/todos/:id", handlers.DeleteApi)

	fiberLambda = fiberadapter.New(app)

	fiber_server = app
}




// Handler will deal with Fiber working with Lambda
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("The .env is not loaded")
	}

	IS_LOCAL := os.Getenv("IS_LOCAL")
	PORT := os.Getenv("BACKEND_PORT")

	if( IS_LOCAL == "true" ){
		log.Fatal(fiber_server.Listen(":" + PORT))
	}else{
		lambda.Start(Handler)
	}
}