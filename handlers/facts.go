package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/iZarrios/docker-compose-go-psql/database"
	"github.com/iZarrios/docker-compose-go-psql/models"
)

func ListFacts(c *fiber.Ctx) error {
	// get all facts
	var facts []models.Fact
	database.DB.Db.Find(&facts)

	return c.Status(http.StatusOK).JSON("hi")
	// return c.Status(http.StatusOK).JSON(facts)
}
func CreateFact(c *fiber.Ctx) error {
	var fact models.Fact
	// unmarshal the body to fact using body parser
	err := c.BodyParser(&fact)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// insert fact in database
	database.DB.Db.Create(&fact)
	return c.Status(http.StatusCreated).JSON(fact)
}

func DeleteFact(c *fiber.Ctx) error {
	// get id from url
	id := c.Params("id")
	// delete fact from database
	database.DB.Db.Delete(&models.Fact{}, id)
	return c.SendStatus(http.StatusNoContent)
}