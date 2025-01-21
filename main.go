package main

import (
	"go-fiber-curd/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	


	// Setup routes
	routes.SetupRoutes(app)

	// Start the server
	app.Listen(":" + "3000")
}