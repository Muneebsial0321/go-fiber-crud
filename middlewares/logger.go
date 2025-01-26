package middlewares

import (
	"log"
	"time"
	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware(c *fiber.Ctx) error {

	start := time.Now()
	err := c.Next() // Process request
	log.Printf("Request took %v\n", time.Since(start))
	return err
}
