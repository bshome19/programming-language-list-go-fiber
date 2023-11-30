package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/bshome19/models"
	"github.com/bshome19/handlers"
)



var Languages = []ProgrammingLanguage{
	{Id: "1", Language: "C", Creator: "Dennis Ritchie"},
	{Id: "2", Language: "Java", Creator: "James Gosling"},
	{Id: "3", Language: "C++", Creator: " Bjarne Stroustrup"},
	{Id: "4", Language: "Python", Creator: "Guido van Rossum"},
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Coders! Welcome to Go programming language.")
	})

	app.Get("/languages", )

	app.Post("/languages", )

	app.Delete("/languages/:id", )

	app.Put("/languages/:id", )

	app.Listen(":8080")
}
