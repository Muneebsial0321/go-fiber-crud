package routes

import (
	"go-fiber-curd/controller"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	user := app.Group("/user")
	user.Get("/", controller.GetUsers)
}