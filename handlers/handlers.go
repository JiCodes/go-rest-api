package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/JiCodes/go-rest-api/database"
	"github.com/JiCodes/go-rest-api/models"
)

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
		return NewFactView(c)
	}

	result := database.DB.Db.Create(&fact)
	if result.Error != nil {
		return NewFactView(c)
	}

	return ListFacts(c)
}

func ShowFact(c *fiber.Ctx) error {
	id := c.Params("id")
	fact := models.Fact{}
	database.DB.Db.First(&fact, id)

	return c.Render("show", fiber.Map{
		"Title": "Single Fact",
		"Fact": fact,
	})
}