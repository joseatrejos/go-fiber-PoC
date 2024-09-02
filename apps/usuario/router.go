package usuario

import (
	"fmt"
	daos "go-fiber-PoC/backend/daos"
	models "go-fiber-PoC/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// SetupRoutes initializes the routes for the entity
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Create DAOs using factory function
	entityDAO := daos.NewBaseDAO(db, models.User{})

	// Routes' handlers
	app.Get("/users/:id", getByID(entityDAO))
	app.Get("/users", getAll(entityDAO))
	app.Post("/users", create(entityDAO))
	app.Put("/users/:id", update(entityDAO))
	app.Delete("/users/:id", delete(entityDAO))
}

func getByID(modelDao *daos.BaseDAO[models.User]) func(c *fiber.Ctx) error {
	// We need to return an anonymous function that receives a fiber.Ctx and returns an error
	return func(c *fiber.Ctx) error {
		// Get the ID from the URL and validate it
		idParam := c.Params("id")
		id, err := strconv.Atoi(idParam)
		if err != nil || id < 1 {
			return c.Status(422).JSON(fiber.Map{"message": fmt.Sprintf("Invalid ID: %s", idParam)})
		}

		// Use the DAO to fetch the record by ID
		getResult, err := modelDao.Get(uint(id))
		if err != nil {
			return c.Status(500).SendString("Error fetching users")
		}
		return c.JSON(getResult)
	}
}

func getAll(modelDao *daos.BaseDAO[models.User]) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		getAllResult, err := modelDao.GetAll()
		if err != nil {
			return c.Status(500).SendString("Error fetching users")
		}
		return c.JSON(getAllResult)
	}
}

func create(modelDao *daos.BaseDAO[models.User]) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Initialize an empty User struct
		var entity models.User

		// Parse the JSON body into the Entity's struct
		if err := c.BodyParser(&entity); err != nil {
			return c.Status(400).SendString("Invalid request data")
		}

		// Use the DAO to create the new record
		createResult := modelDao.Create(&entity)
		if createResult != nil {
			return c.Status(500).JSON(fiber.Map{"message": fmt.Sprintf("Error creating user! %s", createResult.Error())})
		}

		return c.Status(201).JSON(entity)
	}
}

func update(modelDao *daos.BaseDAO[models.User]) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Get the ID from the URL and validate it
		idParam := c.Params("id")
		id, err := strconv.Atoi(idParam)
		if err != nil || id < 1 {
			return c.Status(422).JSON(fiber.Map{"message": fmt.Sprintf("Invalid ID: %s", idParam)})
		}

		// Find the existing resource by ID
		existingEntity, err := modelDao.Get(uint(id))
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"message": "Resource not found"})
		}

		// Parse the JSON body into the entity struct
		if err := c.BodyParser(&existingEntity); err != nil {
			return c.Status(422).JSON(fiber.Map{"message": "Invalid request data"})
		}

		// Use the DAO to update the record
		if err := modelDao.Update(&existingEntity); err != nil {
			return c.Status(500).JSON(fiber.Map{"message": fmt.Sprintf("Error updating resource! %s", err.Error())})
		}

		return c.JSON(existingEntity)
	}
}

func delete(modelDao *daos.BaseDAO[models.User]) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Get the ID from the URL
		idParam := c.Params("id")

		// Validate the ID
		id, err := strconv.Atoi(idParam)
		if err != nil || id < 1 {
			return c.Status(400).JSON(fiber.Map{"message": fmt.Sprintf("Invalid ID: %s", idParam)})
		}

		// Use the DAO to delete the record
		deleteResult := modelDao.Delete(uint(id))
		if deleteResult != nil {
			return c.Status(500).JSON(fiber.Map{"message": fmt.Sprintf("Error deleting user! %s", deleteResult.Error())})
		}
		return c.JSON(fiber.Map{"message": "Resource deleted successfully"})
	}
}
