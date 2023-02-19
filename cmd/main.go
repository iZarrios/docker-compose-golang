package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// disable startup message in fiber

	config := fiber.Config{
		// DisableStartupMessage: true,
		DisableStartupMessage: false,
	}
	app := fiber.New(config)

    setupRoutes(app)


    // listen on port 3000
	app.Listen(":3000")
}

