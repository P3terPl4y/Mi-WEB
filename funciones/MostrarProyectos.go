package funciones

import "github.com/gofiber/fiber/v3"

func MostrarProyectos(c fiber.Ctx) error {
	return c.SendFile("static/view/proyectos.html")
}
