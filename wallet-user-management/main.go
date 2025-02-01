package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GerthDALA/elasta-wallet/wallet-user-management/config"
	"github.com/GerthDALA/elasta-wallet/wallet-user-management/controllers"
	"github.com/GerthDALA/elasta-wallet/wallet-user-management/repositories"
	"github.com/GerthDALA/elasta-wallet/wallet-user-management/routes"
	"github.com/GerthDALA/elasta-wallet/wallet-user-management/services"
)



func main() {
	// Load environment variables
	config.LoadEnv()

	// Load configuration structure
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := config.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repository, service, and controller for user operations.
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Register API routes.
	routes.RegisterRoutes(userController)

	// Start the HTTP server.
	fmt.Printf("Server is running on port %s\n", cfg.AppPort)
	if err := http.ListenAndServe(":"+cfg.AppPort, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
