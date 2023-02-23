package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/iZarrios/docker-compose-golang/database"
)

func main() {

    database.Connect()
	// disable startup message in fiber

	config := fiber.Config{
		DisableStartupMessage: true,
		// DisableStartupMessage: false,
	}
    //use cors
	app := fiber.New(config)
    app.Use(logger.New())
    app.Use(cors.New())

    setupRoutes(app)

    // listen on port 3000
	app.Listen(":3000")
}
