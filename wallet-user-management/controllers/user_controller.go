package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/GerthDALA/elasta-wallet/wallet-user-management/models"
	"github.com/GerthDALA/elasta-wallet/wallet-user-management/services"

)

// UserController handles HTTP requests for user-related operations.
type UserController struct {
	UserService services.UserService
}

// NewUserController creates a new UserController.
func NewUserController(userService services.UserService) *UserController {
	return &UserController{UserService: userService}
}

// CreateUserHandler handles POST requests to create a new user.
func (uc *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read and parse the request body.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, "Bad Request: Invalid JSON", http.StatusBadRequest)
		return
	}

	// In a production system, you should hash the user's password here.

	// Use the request's context for tracing and cancellation.
	ctx := r.Context()
	userID, err := uc.UserService.CreateUser(ctx, user)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Respond with the created user's ID.
	resp := map[string]interface{}{
		"id": userID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
