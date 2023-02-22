package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/iZarrios/docker-compose-go-psql/database"
)

func main() {

    database.Connect()
	// disable startup message in fiber

	config := fiber.Config{
		DisableStartupMessage: true,
		// DisableStartupMessage: false,
	}
	app := fiber.New(config)
    app.Use(logger.New())

    setupRoutes(app)

    // listen on port 3000
	app.Listen(":3000")
}
