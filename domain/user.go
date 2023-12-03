package domain

import "time"

type User struct {
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Id        UUID      `json:"id"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	Create(user User) (*User, error)
	GetByEmail(email string) (*User, error)
	GetById(id string) (*User, error)
}
