package main

import (
	"github.com/carlosmoran97/go-auth/database"
	"github.com/carlosmoran97/go-auth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	app := fiber.New()

	routes.Setup(app)

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Listen(":8000")
}
