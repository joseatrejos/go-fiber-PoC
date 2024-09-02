package main

import (
	config "go-fiber-PoC/config"
	routers "go-fiber-PoC/routers"
)

func main() {

	// Load environment variables
	config.LoadEnv()

	// Initialize the Fiber app
	app := config.InitializeFiberApp()

	// Initialize the database
	config.InitDB()

	// Routes startup
	routers.Setup(app)

	// Start server
	config.StartServer(app)
}
