package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/mofodox/omh-rest-api/database"
	"github.com/mofodox/omh-rest-api/routes"
)

func main() {
	app := fiber.New()

	app.Use(middleware.Logger())

	database.Connect()
	
	routes.SetupRoutes(app)

	app.Listen(":3000")
}
