package main

import (
	"github.com/carlosmoran97/go-auth/database"
	"github.com/carlosmoran97/go-auth/routes"
	"github.com/gofiber/fiber"
)

func main() {
	database.Connect()
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
