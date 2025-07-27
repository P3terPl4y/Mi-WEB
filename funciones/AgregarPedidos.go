package funciones

import (
	"Proyecto/database"
	"Proyecto/modelos"
	"log"

	"github.com/gofiber/fiber/v3"
)

func AgregarPedido(c fiber.Ctx) error {

	pedido := new(modelos.Datos)

	err := c.Bind().Body(pedido)

	if err != nil {
		return c.SendStatus(500)
	}

	result, err := database.DB.Exec("INSERT INTO pedidos (cliente, pedido, telefono) VALUES (?, ?, ?)", pedido.Cliente, pedido.Pedido, pedido.Telefono)

	if err != nil {
		log.Println(result)
		return c.SendStatus(500)
	} else {
		return c.SendStatus(200)
	}
}
