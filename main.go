package main

import (
	"os"

	"github.com/CesarDelgadoM/tutorials-api/api/handlers"
	"github.com/CesarDelgadoM/tutorials-api/api/routes"
	"github.com/CesarDelgadoM/tutorials-api/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())

	routes := routes.NewRoutes(app)

	db := database.ConnectMongoDB()
	handler := handlers.NewTutorialHandler(db)

	routes.TutorialRoutes(handler)

	app.Listen(os.Getenv("PORT"))
}
