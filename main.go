package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/bshome19/programming-language-list-go-fiber/handlers"
)




func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Coders! Welcome to Go programming language.")
	})

	app.Get("/languages", handlers.GetAllLanguagesData)

	app.Post("/languages", handlers.CreateNewLanguageData)

	app.Delete("/languages/:id", handlers.DeleteLanguageById)

	app.Put("/languages/:id", handlers.UpdateLanguageDataById)

	app.Listen(":8080")
}
