package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/JiCodes/go-rest-api/database"

)

func main() {
		database.ConnectDb()
		
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, Go API with fiber :)")
    })

    app.Listen(":3000")
}