package controller;

import (
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

// func CreateUser(c *fiber.Ctx) error {
// 	var user models.User
// 	if err := c.BodyParser(&user); err != nil {
// 		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	if err := services.CreateUser(&user); err != nil {
// 		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return c.JSON(user)
// }
