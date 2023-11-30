package repository

import (
	"database/sql"
	"go.uber.org/zap"
	"inventory/domain"
)

type productRepository struct {
	Db     *sql.DB
	Logger *zap.Logger
}

func (p productRepository) Create(product domain.Product) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p productRepository) Get(id string) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p productRepository) Update(id string, product domain.Product) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p productRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func (p productRepository) GetAll(page, limit int64) ([]domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func NewProductRepository(db *sql.DB, logger *zap.Logger) domain.ProductRepository {
	return &productRepository{
		Db:     db,
		Logger: logger,
	}
}
