package routers

import (
	expediente "go-fiber-PoC/apps/expediente"
	usuario "go-fiber-PoC/apps/usuario"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Setup initializes all the routes for the application's entities/resources
func Setup(app *fiber.App, db *gorm.DB) {
	usuario.SetupRoutes(app, db)
	expediente.SetupRoutes(app, db)
}
