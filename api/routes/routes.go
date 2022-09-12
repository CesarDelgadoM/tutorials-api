package routes

import (
	"github.com/CesarDelgadoM/tutorials-api/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func TutorialRoutes(app *fiber.App) {

	app.Get("/api/tutorials", handlers.GetAllTutorials)
	app.Get("/api/tutorial/:id", handlers.GetTutorialById)
	app.Post("/api/tutorial", handlers.CreateTutorial)
	app.Put("/api/tutorial/:id", handlers.UpdateTutorial)
	app.Delete("/api/tutorial/:id", handlers.DeleteTutorialById)
	app.Delete("/api/tutorials", handlers.DeleteAllTutorials)
	app.Get("/api/tutorial/title/:title", handlers.GetTutorialByTitle)
}
