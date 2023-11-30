package domain

import "time"

type Sale struct {
	Id          UUID      `json:"id"`
	ProductName string    `json:"product_name"`
	ProductId   UUID      `json:"product_id"`
	Amount      int64     `json:"amount"`
	Total       float64   `json:"total"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SaleRepository interface {
	Create(sale Sale) (*Sale, error)
	Get(id string) (*Sale, error)
	Update(id string, sale Sale) (*Sale, error)
	Delete(id string) error
	GetAll(page, limit int64) ([]Sale, error)
}
