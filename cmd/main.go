package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"github.com/JiCodes/go-rest-api/database"

	"github.com/JiCodes/go-rest-api/handlers"
)

func main() {
		database.ConnectDb()

		// Create a new engine
		engine := html.New("./views", ".html")

		app := fiber.New(fiber.Config{
			Views: engine,
			ViewsLayout: "layouts/main",
		})

		setUpRoutes(app)

		app.Static("/", "./public")

		app.Use(handlers.NotFound)

    app.Listen(":3000")
}