package routes

import (
	"go-fiber-curd/controller"
	"go-fiber-curd/middlewares"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {

	user := app.Group("/user")
	user.Get("/", middlewares.LoggerMiddleware, controller.User_GetAll)
	user.Post("/", controller.User_Post)
}
