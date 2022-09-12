package main

import (
	"os"

	"github.com/CesarDelgadoM/tutorials-api/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())
	routes.TutorialRoutes(app)
	app.Listen(os.Getenv("PORT"))
}
