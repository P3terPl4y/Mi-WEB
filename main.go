package main

import (
	"Proyecto/database"
	"Proyecto/middleware"
	"Proyecto/rutas"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {

	database.ConectarDB()
	defer database.DB.Close()

	app := fiber.New()

	app.Use(static.New("static"))

	app.Use(cors.New())
	app.Use(logger.New())

	app.Use(middleware.Middleware)

	rutas.Rutas(app)

	log.Fatal(app.Listen(":8000"))
}
