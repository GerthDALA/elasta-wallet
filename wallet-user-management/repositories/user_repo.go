package repositories

import (
	"context"
	"database/sql"
	"github.com/GerthDALA/elasta-wallet/wallet-user-management/models"
)

// UserRepository defines data-access methids for users
type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) (int64, error)
}

type userRepo struct {
	DB *sql.DB
}

// NewUserRepository returns an implementation of UserRepository.
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{DB: db}
}

// CreateUser inserts a new user into the database.
func (ur *userRepo) CreateUser(ctx context.Context, user models.User) (int64, error) {

	query := `
	    INSERT INTO users (email, phone, hashed_password, first_name, last_name, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING id;
	`

	var id int64
	err := ur.DB.QueryRowContext(ctx, query,
		user.Email,
		user.Phone,
		user.HashedPassword,
		user.FirstName,
		user.LastName,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil

}
