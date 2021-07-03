package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/joho/godotenv"
	"github.com/mofodox/omh-rest-api/database"
	"github.com/mofodox/omh-rest-api/routes"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("unable to load values from env file")
	} else {
		log.Println("successfully loaded the values from env file")
	}
}

func main() {
	app := fiber.New()

	app.Use(middleware.Logger())

	database.Connect()
	
	routes.SetupRoutes(app)

	app.Listen(":3000")
}
