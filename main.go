package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-curd/routes"
)

func main() {

	app := fiber.New()
	


	// Setup routes
	routes.UserRoutes(app)
	
	// Start the server
	app.Listen(":" + "3000")
}