package usuario_router

import (
	"go-fiber-PoC/config"
	"go-fiber-PoC/models"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/users", getUsers)
	app.Post("/users", createUser)
	app.Put("/users/:id", updateUser)
	app.Delete("/users/:id", deleteUser)
}

func getUsers(c *fiber.Ctx) error {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(users)
}

func createUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

func updateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	config.DB.Save(&user)
	return c.Status(fiber.StatusOK).JSON(user)
}

func deleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	config.DB.Delete(&user)
	return c.Status(fiber.StatusNoContent).SendString("User deleted successfully")
}
