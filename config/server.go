package config

import (
	"os"
	"runtime"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// InitializeFiberApp initializes and returns a new Fiber instance
func InitializeFiberApp() *fiber.App {

	// We default to 1 core
	numCores := 1

	// Get the desired number of CPU cores from the environment variable
	envCores := os.Getenv("CPU_CORES")

	if envCores == "MAX" {
		numCores = runtime.NumCPU()
	} else if envCores != "" {
		// Parse the number of cores from the environment variable
		parsedCores, err := strconv.Atoi(envCores)
		if err == nil && parsedCores > 0 {
			numCores = parsedCores
		}
	}

	runtime.GOMAXPROCS(numCores)

	// Always use Prefork
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	return app
}

// Start server
func StartServer(app *fiber.App) {
	port := "8000"
	parsedPort, err := strconv.Atoi(os.Getenv("PORT"))
	if err == nil && parsedPort > 0 {
		port = strconv.Itoa(parsedPort)
	}
	app.Listen(":" + port)
}
