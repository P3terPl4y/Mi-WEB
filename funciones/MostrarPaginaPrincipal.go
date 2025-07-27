package funciones

import "github.com/gofiber/fiber/v3"

func MostrarPaginaPrincipal(c fiber.Ctx) error {
	return c.SendFile("static/view/index.html")
}
