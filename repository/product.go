package repository

import (
	"database/sql"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"inventory/domain"
	"time"
)

type productRepository struct {
	Db     *sql.DB
	Logger *zap.Logger
}

func (p productRepository) Create(product domain.Product) (*domain.Product, error) {

	product.Id = domain.UUID{UUID: uuid.NewV4()}
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	_, err := p.Db.Exec("INSERT INTO products VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", product.Name, product.Id, product.Price, product.Stock, product.MinStock, product.SupplierId, product.CreatedAt, product.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p productRepository) Get(id string) (*domain.Product, error) {
	row := p.Db.QueryRow(fmt.Sprintf("SELECT * FROM products WHERE id ='%s'", id))

	if row.Err() != nil {
		return nil, row.Err()
	}

	var product domain.Product

	err := row.Scan(&product.Name, &product.Id, &product.Price, &product.Stock, &product.MinStock, &product.SupplierId, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product with that id not found")
		}
		return nil, err
	}

	return &product, nil
}

func (p productRepository) Update(id string, product domain.Product) (*domain.Product, error) {
	_, err := p.Db.Exec("UPDATE products SET name=$1, price=$2,stock=$3, min_stock=$4, updated_at=$5 WHERE id=$6", product.Name, product.Price, product.Stock, product.MinStock, time.Now(), id)

	if err != nil {
		return nil, err
	}

	product_, err := p.Get(id)

	if err != nil {
		return nil, err
	}

	return product_, nil
}

func (p productRepository) DecrementStock(id string, tx *sql.Tx) error {
	_, err := tx.Exec("UPDATE products SET stock=stock-1 WHERE id=$1", id)

	if err != nil {
		return err
	}

	return nil
}

func (p productRepository) Delete(id string) error {
	_, err := p.Db.Exec("DELETE from products WHERE id=$1", id)

	if err != nil {
		return err
	}

	return nil
}

func (p productRepository) GetAll(page, limit int64) (*domain.PaginatedProducts, error) {
	var products []domain.Product
	var product domain.Product

	offset := (page - 1) * limit

	rows, err := p.Db.Query("SELECT * FROM products OFFSET $1 LIMIT $2", offset, limit)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&product.Name, &product.Id, &product.Price, &product.Stock, &product.MinStock, &product.SupplierId, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	var count int

	err = p.Db.QueryRow("SELECT  COUNT(*) FROM products").Scan(&count)

	if err != nil {
		return nil, err
	}

	readRows := page * limit

	var next int64

	if int64(count) > readRows {
		next = page + 1
	}

	return &domain.PaginatedProducts{
		Products: products,
		Page:     page,
		Limit:    limit,
		Total:    int64(count),
		Next:     next,
	}, nil
}

func (p productRepository) GetRestockRecommendation() (*domain.PaginatedProducts, error) {
	//TODO implement me
	panic("implement me")
}

func NewProductRepository(db *sql.DB, logger *zap.Logger) domain.ProductRepository {
	return &productRepository{
		Db:     db,
		Logger: logger,
	}
}
