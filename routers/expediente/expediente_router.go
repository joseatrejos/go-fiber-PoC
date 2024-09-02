package expediente_router

import (
	"go-fiber-PoC/config"
	"go-fiber-PoC/models"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/expedientes", getExpedientes)
	app.Post("/expedientes", createExpediente)
	app.Put("/expedientes/:id", updateExpediente)
	app.Delete("/expedientes/:id", deleteExpediente)
}

func getExpedientes(c *fiber.Ctx) error {
	var expedientes []models.Expediente
	if err := config.DB.Find(&expedientes).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(expedientes)
}

func createExpediente(c *fiber.Ctx) error {
	var expediente models.Expediente
	if err := c.BodyParser(&expediente); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := config.DB.Create(&expediente).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(expediente)
}

func updateExpediente(c *fiber.Ctx) error {
	id := c.Params("id")
	var expediente models.Expediente
	if err := config.DB.First(&expediente, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	if err := c.BodyParser(&expediente); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	config.DB.Save(&expediente)
	return c.Status(fiber.StatusOK).JSON(expediente)
}

func deleteExpediente(c *fiber.Ctx) error {
	id := c.Params("id")
	var expediente models.Expediente
	if err := config.DB.First(&expediente, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	config.DB.Delete(&expediente)
	return c.Status(fiber.StatusNoContent).SendString("Expediente deleted successfully")
}
