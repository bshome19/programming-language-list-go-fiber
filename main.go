package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/bshome19/programming-language-list-go-fiber/routers"
)




func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Coders! Welcome to Go programming language.")
	})
	routers.SetupRoutes(app)
	
	app.Listen(":8080")
}
