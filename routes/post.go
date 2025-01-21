package routes

import "github.com/gofiber/fiber/v2"

func PostRoutes(app *fiber.App) {
	posts := app.Group("/posts")
	posts.Get("/",)
}
