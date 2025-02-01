package models

import "time"

type User struct {
	ID             int64     `json:"id"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	HashedPassword string    `json:"hashed_password"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
