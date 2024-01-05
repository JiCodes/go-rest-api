package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/JiCodes/go-rest-api/database"
	"github.com/JiCodes/go-rest-api/models"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, Go API with fiber :)")
}

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Render("index", fiber.Map{
		"Title": "Facts",
		"Subtitle": "Facts about Go",
		"Facts": facts,
		})
}

func NewFactView(c *fiber.Ctx) error {
	return c.Render("new", fiber.Map{
		"Title": "New Fact",
		"Subtitle": "Add a new fact to the list!",
	})
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&fact)

	return ConfirmationView(c)
}

func ConfirmationView(c *fiber.Ctx) error {
	return c.Render("confirmation", fiber.Map{
		"Title": "Fact added successfully!",
		"Subtitle": "Add more facts to the list!",
	})
}