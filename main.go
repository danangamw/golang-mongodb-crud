package main

import (
	"fmt"
	"log"

	"github.com/danangamw/go_mongo_crud/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .enf file")
	}

	app := fiber.New()

	app.Post("/users", handlers.CreateUser)
	app.Get("/users", handlers.GetAllUsers)
	app.Get("/users/:id", handlers.GetUserByID)
	app.Put("/users/:id", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.DeleteUser)

	fmt.Println("Server running on Port 8000")
	log.Fatal(app.Listen(":8000"))
}
