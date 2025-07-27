package funciones

import (
	"Proyecto/database"
	"Proyecto/modelos"
	"log"

	"github.com/gofiber/fiber/v3"
)

func MostrarPedidos(c fiber.Ctx) error {
	filas, err := database.DB.Query("SELECT * FROM pedidos;")

	if err != nil {
		log.Fatal(err)
		return c.SendStatus(500)
	}
	defer filas.Close()

	var lista_de_pedidos []modelos.Datos

	for filas.Next() {
		var pedido modelos.Datos

		err := filas.Scan(&pedido.Cliente, &pedido.Pedido, &pedido.Telefono)

		if err != nil {
			log.Fatal(err)
			return c.SendStatus(500)
		}
		lista_de_pedidos = append(lista_de_pedidos, pedido)
	}

	return c.JSON(lista_de_pedidos)
}
