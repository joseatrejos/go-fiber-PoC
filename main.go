package main

import (
	routers "go-fiber-PoC/backend/routers"
	config "go-fiber-PoC/config"
)

func main() {

	// Load environment variables
	config.LoadEnv()

	// Initialize the Fiber app
	app := config.InitializeFiberApp()

	// Initialize the database
	config.InitDB()

	// Router setup
	routers.Setup(app, config.DB)

	// Start server
	config.StartServer(app)
}
