package services

import (
	"context"
	"log"
	"time"

	"github.com/GerthDALA/elasta-wallet/wallet-user-management/models"
	"github.com/GerthDALA/elasta-wallet/wallet-user-management/repositories"
)

// UserService defines business operations logic for user operations.
type UserService interface {
	CreateUser(ctx context.Context, user models.User) (int64, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

// NewUserService returns an implementation of UserService.
func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// CreateUser creates a new user with tracing.
func (us *userService) CreateUser(ctx context.Context, user models.User) (int64, error) {
	// Begin tracing the operation.
	startTime := time.Now()
	log.Println("[Tracing] Start CreateUser service")

	id, err := us.userRepo.CreateUser(ctx, user)
	if err != nil {
		log.Printf("[Tracing] CreateUser encountered an error: %v", err)
		return 0, err
	}

	// End tracing the operation.
	elapsed := time.Since(startTime)
	log.Printf("[Tracing] End CreateUser service, elapsed time: %s", elapsed)

	return id, nil
}