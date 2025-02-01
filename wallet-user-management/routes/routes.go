package routes

import (
	"net/http"
    "github.com/GerthDALA/elasta-wallet/wallet-user-management/controllers"
    
)
// RegisterRoutes initializes all API routes.
func RegisterRoutes(userController *controllers.UserController) {
	// Healthcheck endpoint.
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "OK"}`))
	})

	// User endpoints.
	http.HandleFunc("/users", userController.CreateUserHandler)
}