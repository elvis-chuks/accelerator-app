package domain

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type UUID struct {
	uuid.UUID
}

type Product struct {
	Name       string `json:"name"`
	Id         UUID
	Price      float64   `json:"price"`
	Stock      int64     `json:"stock"`
	MinStock   int64     `json:"min_stock"`
	SupplierId UUID      `json:"supplier_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ProductRepository interface {
	Create(product Product) (*Product, error)
	Get(id string) (*Product, error)
	Update(id string, product Product) (*Product, error)
	Delete(id string) error
	GetAll(page, limit int64) ([]Product, error)
}
