package routes

import (
	"github.com/gofiber/fiber"
	"github.com/mofodox/omh-rest-api/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controllers.Home)

	apiV1PropertiesV1 := app.Group("/api/v1/properties")
	apiV1PropertiesV1.Get("/", controllers.ShowAllProperties)
	apiV1PropertiesV1.Get("/:id", controllers.ShowProperty)
	apiV1PropertiesV1.Post("/", controllers.AddProperty)
	apiV1PropertiesV1.Put("/:id", controllers.UpdateProperty)
	apiV1PropertiesV1.Delete("/:id", controllers.DeleteProperty)

	apiV1Countries := app.Group("/api/v1/countries")
	apiV1Countries.Get("/", controllers.GetAllCountries)
	apiV1Countries.Get("/:id", controllers.GetCountry)
	apiV1Countries.Post("/", controllers.CreateCountry)
}
