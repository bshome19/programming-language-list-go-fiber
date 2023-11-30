package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/bshome19/programming-language-list-go-fiber/handlers"

)

func SetupRoutes(app *fiber.App) {

	app.Get("/languages", handlers.GetAllLanguagesData)
	app.Post("/languages", handlers.CreateNewLanguageData)
	app.Delete("/languages/:id", handlers.DeleteLanguageById)
	app.Put("/languages/:id", handlers.UpdateLanguageDataById)

}