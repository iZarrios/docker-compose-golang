package main

import "github.com/gofiber/fiber/v2"

func main() {
    // disable startup message in fiber

    config := fiber.Config{
        DisableStartupMessage: true,
    }
    app := fiber.New(config)

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Helllo, World111!")
    })

    app.Listen(":3000")
}
