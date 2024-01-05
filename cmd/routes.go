package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/JiCodes/go-rest-api/handlers"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Get("/fact", handlers.NewFactView)

	app.Post("/fact", handlers.CreateFact)

	// add new routes to 'show' single fact, given ':id'	
	app.Get("/fact/:id", handlers.ShowFact)

	app.Get("/fact/:id/edit", handlers.EditFact)

	app.Patch("/fact/:id/", handlers.UpdateFact)
	
}



