package routes

import (
	"github.com/gofiber/fiber"
	"github.com/mofodox/omh-rest-api/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controllers.Home)

	apiV1 := app.Group("/api/v1")
	apiV1.Get("/properties", controllers.ShowAllProperties)
	apiV1.Get("/properties/:id", controllers.ShowProperty)
	apiV1.Post("/properties", controllers.AddProperty)
	apiV1.Patch("/properties/:id", controllers.UpdateProperty)
	apiV1.Delete("/properties/:id", controllers.DeleteProperty)
}
