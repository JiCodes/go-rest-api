package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/JiCodes/go-rest-api/database"

)

func main() {
		database.ConnectDb()
    app := fiber.New()

		setUpRoutes(app)

    app.Listen(":3000")
}