package routes

import (
	"go-fiber-curd/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// User routes
	user := app.Group("/user")
	user.Get("/", controller.GetUsers)
}