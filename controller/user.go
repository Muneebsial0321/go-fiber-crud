package controller

import (
	"fmt"
	"go-fiber-curd/models"

	"github.com/gofiber/fiber/v2"
)

func User_GetAll(c *fiber.Ctx) error {
	users := models.User{ID: 1, Name: "muneeb", Email: "email"}
	// if err != nil {
	// 	return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	// }
	return c.JSON(users)
}

func User_Post(c *fiber.Ctx) error {
	fmt.Println("posting user")
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	return c.JSON(user)
}

