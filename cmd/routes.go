package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/iZarrios/docker-compose-golang/handlers"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/facts", handlers.ListFacts)
	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	api.Post("/facts", handlers.CreateFact)
	api.Delete("/facts/:id", handlers.DeleteFact)

	api.Get("/facts/:id", handlers.GetFact)

	// static files interface
	api.Static("/", "./public", fiber.Static{
		Index: "index.html",
	})

	// expirementing with websockets
	// make a handler that uses gorilla mux to handle websockets
	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			fmt.Println("allowed")
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		// c.Locals is added to the *websocket.Conn
		log.Println(c.Locals("allowed"))  // true
		log.Println(c.Params("id"))       // 123
		log.Println(c.Query("v"))         // 1.0
		log.Println(c.Cookies("session")) // ""

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)

			if err = c.WriteMessage(mt, []byte(fmt.Sprintf("got (%v)",
				string(msg)))); err != nil {
				log.Println("write:", err)
				break
			}
		}

	}))

}
