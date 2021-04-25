package routes

import (
	"github.com/carlosmoran97/go-auth/controllers"
	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
}
