package routes

import (
	"github.com/CesarDelgadoM/tutorials-api/api/handlers"
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	app *fiber.App
}

func NewRoutes(app *fiber.App) *Routes {
	return &Routes{
		app: app,
	}
}

func (r *Routes) TutorialRoutes(handlers *handlers.TutorialHandlers) {

	r.app.Get("/api/tutorials", handlers.GetAllTutorials)
	r.app.Get("/api/tutorial/:id", handlers.GetTutorialById)
	r.app.Post("/api/tutorial", handlers.CreateTutorial)
	r.app.Put("/api/tutorial/:id", handlers.UpdateTutorial)
	r.app.Delete("/api/tutorial/:id", handlers.DeleteTutorialById)
	r.app.Delete("/api/tutorials", handlers.DeleteAllTutorials)
	r.app.Get("/api/tutorial/title/:title", handlers.GetTutorialByTitle)
}
