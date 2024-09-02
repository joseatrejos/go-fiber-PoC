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

	// Router setup
	routers.Setup(app, config.DB)

	// Start server
	config.StartServer(app)
}
