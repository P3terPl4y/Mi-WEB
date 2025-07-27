package middleware

import "github.com/gofiber/fiber/v3"

func Middleware(c fiber.Ctx) error {

	return c.Next()
}
