package middlewares

import "github.com/gofiber/fiber/v2"

func JwtAuth(c *fiber.Ctx) {
	c.GetReqpHeaders('Authen')
}
