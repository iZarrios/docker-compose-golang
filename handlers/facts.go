package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/iZarrios/docker-compose-golang/database"
	"github.com/iZarrios/docker-compose-golang/models"
)

func ListFacts(c *fiber.Ctx) error {
	// get all facts
	var facts []models.Fact
	err := database.DB.Db.Find(&facts).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(facts)
	// return c.Status(http.StatusOK).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	var fact models.Fact
	// unmarshal the body to fact using body parser
	err := c.BodyParser(&fact)
    fmt.Println("got:",fact)
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

func GetFact(c* fiber.Ctx) error {
    id := c.Params("id")
    var fact models.Fact
    err := database.DB.Db.First(&fact, id).Error
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(http.StatusOK).JSON(fact)
}
