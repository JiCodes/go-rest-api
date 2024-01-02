package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/JiCodes/go-rest-api/handlers"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	app.Get("/facts", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)
}



