package handlers

import (
	"net/http"
	"github.com/bshome19/programming-language-list-go-fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/bshome19/models"
)


func GetAllLanguagesData(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(Languages)
}

func CreateNewLanguageData(c *fiber.Ctx) error {
	var language ProgrammingLanguage

	if err := c.BodyParser(&language); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return err
	}
	Languages = append(Languages, language)
	return c.Status(http.StatusCreated).JSON(language)
}

func DeleteLanguageById(c *fiber.Ctx) error {
	id := c.Params("id")

	for i, lang := range Languages {
		if lang.Id == id {
			Languages = append(Languages[:i], Languages[i+1:]...)
			break
		}
	}
	return c.Status(http.StatusNoContent).JSON(fiber.Map{"data": Languages})

}

func UpdateLanguageDataById(c *fiber.Ctx) error {
	id := c.Params("id")
	var updatedlanguage ProgrammingLanguage
	updatedlanguage.Id = id 

	if err := c.BodyParser(&updatedlanguage); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return err
	}
	for i, lang := range Languages {
		if lang.Id == updatedlanguage.Id {
			Languages = append(Languages[:i], Languages[i+1:]...)
			Languages = append(Languages, updatedlanguage)
		}
	}
	return c.Status(http.StatusCreated).JSON(updatedlanguage)
}