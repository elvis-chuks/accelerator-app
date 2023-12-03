package domain

import "time"

type Sale struct {
	Id          UUID      `json:"id"`
	ProductName string    `json:"product_name" validate:"required,max=256"`
	ProductId   UUID      `json:"product_id" validate:"required,uuid"`
	Quantity    int64     `json:"quantity" validate:"required,max=256"`
	Total       float64   `json:"total"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PaginatedSales struct {
	Sales []Sale `json:"sales"`
	Page  int64  `json:"page"`
	Limit int64  `json:"limit"`
	Total int64  `json:"total"`
	Next  int64  `json:"next"`
}

type SaleRepository interface {
	Create(sale Sale) (*Sale, error)
	Get(id string) (*Sale, error)
	Update(id string, sale Sale) (*Sale, error)
	Delete(id string) error
	GetAll(page, limit int64) (*PaginatedSales, error)
}
