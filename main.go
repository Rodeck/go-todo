package main

import (
	"log"

	"github.com/Rodeck/go-todo/database"
	"github.com/Rodeck/go-todo/todo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(cors.New())
	database.ConnectDB()
	defer database.DB.Close()

	api := app.Group("/api")

	// Test handler
	todo.Register(api, database.DB)

	log.Fatal(app.Listen(":5000"))
}
