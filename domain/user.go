package domain

import "time"

type User struct {
	Fullname  string    `json:"fullname" validate:"required,max=256"`
	Email     string    `json:"email" validate:"required,email"`
	Id        UUID      `json:"id"`
	Password  string    `json:"password,omitempty" validate:"required,max=256"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	Create(user User) (*User, error)
	GetByEmail(email string) (*User, error)
	GetById(id string) (*User, error)
}
