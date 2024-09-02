package routers

import (
	expediente_router "go-fiber-PoC/routers/expediente"
	usuario_router "go-fiber-PoC/routers/usuario"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Setup initializes all the routes for the application's entities/resources
func Setup(app *fiber.App, db *gorm.DB) {
	usuario_router.SetupRoutes(app, db)
	expediente_router.SetupRoutes(app, db)
}
