package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"inventory/domain"
	"log"
)

func SetupDb(l *zap.Logger) *domain.Repository {

	connStr := viper.GetString("DB_URL")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &domain.Repository{
		ProductRepo: NewProductRepository(db, l),
		SaleRepo:    NewSaleRepository(db, l),
	}
}
