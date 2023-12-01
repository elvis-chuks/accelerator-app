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

type saleRepository struct {
	Db          *sql.DB
	Logger      *zap.Logger
	productRepo domain.ProductRepository
}

func (s saleRepository) Create(sale domain.Sale) (*domain.Sale, error) {
	sale.Id = domain.UUID{UUID: uuid.NewV4()}
	sale.CreatedAt = time.Now()
	sale.UpdatedAt = time.Now()

	product, err := s.productRepo.Get(sale.ProductId.String())

	if err != nil {
		return nil, err
	}

	tx, err := s.Db.Begin()

	if err != nil {
		return nil, err
	}

	sale.Total = product.Price * float64(sale.Amount)

	_, err = tx.Exec("INSERT INTO sales VALUES ($1, $2, $3, $4, $5, $6, $7)", sale.Id, sale.ProductName, sale.ProductId, sale.Amount, sale.Total, sale.CreatedAt, sale.UpdatedAt)

	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	err = s.productRepo.DecrementStock(sale.ProductId.String(), tx)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &sale, nil
}

func (s saleRepository) Get(id string) (*domain.Sale, error) {
	row := s.Db.QueryRow(fmt.Sprintf("SELECT * FROM sales WHERE id ='%s'", id))

	if row.Err() != nil {
		return nil, row.Err()
	}

	var sale domain.Sale

	err := row.Scan(&sale.Id, &sale.ProductName, &sale.ProductId, &sale.Amount, &sale.Total, &sale.CreatedAt, &sale.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("sale with that id not found")
		}
		return nil, err
	}

	return &sale, nil
}

func (s saleRepository) Update(id string, sale domain.Sale) (*domain.Sale, error) {
	_, err := s.Db.Exec("UPDATE sales SET amount=$1,total=$2, updated_at=$3 WHERE id=$4", sale.Amount, sale.Total, time.Now(), id)

	if err != nil {
		return nil, err
	}

	sale_, err := s.Get(id)

	if err != nil {
		return nil, err
	}

	return sale_, nil
}

func (s saleRepository) Delete(id string) error {
	_, err := s.Db.Exec("DELETE from sales WHERE id=$1", id)

	if err != nil {
		return err
	}

	return nil
}

func (s saleRepository) GetAll(page, limit int64) (*domain.PaginatedSales, error) {
	var sales []domain.Sale
	var sale domain.Sale

	offset := (page - 1) * limit

	rows, err := s.Db.Query("SELECT * FROM sales ORDER BY created_at OFFSET $1 LIMIT $2", offset, limit)

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
		err = rows.Scan(&sale.Id, &sale.ProductName, &sale.ProductId, &sale.Amount, &sale.Total, &sale.CreatedAt, &sale.UpdatedAt)
		if err != nil {
			return nil, err
		}
		sales = append(sales, sale)
	}

	var count int

	err = s.Db.QueryRow("SELECT  COUNT(*) FROM sales").Scan(&count)

	if err != nil {
		return nil, err
	}

	readRows := page * limit

	var next int64

	if int64(count) > readRows {
		next = page + 1
	}

	return &domain.PaginatedSales{
		Sales: sales,
		Page:  page,
		Limit: limit,
		Total: int64(count),
		Next:  next,
	}, nil
}

func NewSaleRepository(db *sql.DB, logger *zap.Logger, repository domain.ProductRepository) domain.SaleRepository {
	return &saleRepository{
		Db:          db,
		Logger:      logger,
		productRepo: repository,
	}
}
