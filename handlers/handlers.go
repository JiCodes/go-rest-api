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
	result := database.DB.Db.First(&fact, id)
	if result.Error != nil {
		return NotFound(c)
	}

	return c.Render("show", fiber.Map{
		"Title": "Single Fact",
		"Fact": fact,
	})
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendFile("./public/404.html")
}

func EditFact(c *fiber.Ctx) error {
	id := c.Params("id")
	fact := models.Fact{}
	result := database.DB.Db.First(&fact, id)
	if result.Error != nil {
		return NotFound(c)
	}

	return c.Render("edit", fiber.Map{
		"Title": "Edit Fact",
		"Subtitle": "Edit this fact!",
		"Fact": fact,
	})
}

func UpdateFact(c *fiber.Ctx) error {
	id := c.Params("id")
	fact := new(models.Fact) 
	
	// Parse the request body and populate the 'fact' struct with the data
	if err := c.BodyParser(fact); err != nil { 
		return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
	}
	
	result := database.DB.Db.Model(&fact).Where("id = ?", id).Updates(fact)
	if result.Error != nil {
		return EditFact(c)
	}

	return ShowFact(c)
}

func DeleteFact(c *fiber.Ctx) error {
	id := c.Params("id")
	fact := models.Fact{}
	result := database.DB.Db.Delete(&fact, id)
	if result.Error != nil {
		return NotFound(c)
	}

	return ListFacts(c)
}