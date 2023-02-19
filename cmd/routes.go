package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iZarrios/docker-compose-go-psql/handlers"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/facts", handlers.ListFacts)
	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	api.Post("/facts", handlers.CreateFact)
	api.Delete("/facts/:id", handlers.DeleteFact)
}

// curl command ot list all ListFacts
// curl -X GET http://localhost:3000/api/v1/facts
