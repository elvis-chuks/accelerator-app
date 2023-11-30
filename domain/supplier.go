package domain

import "time"

type Supplier struct {
	Name      string    `json:"name"`
	Id        UUID      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
