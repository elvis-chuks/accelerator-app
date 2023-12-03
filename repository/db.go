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

	err = CreateTables(db)

	if err != nil {
		log.Fatal(err)
	}

	productRepo := NewProductRepository(db, l)

	return &domain.Repository{
		ProductRepo: productRepo,
		SaleRepo:    NewSaleRepository(db, l, productRepo),
		UserRepo:    NewUserRepository(db, l),
	}
}

func CreateTables(db *sql.DB) error {

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS products (name VARCHAR(255),id UUID,price DOUBLE PRECISION,stock BIGINT,min_stock BIGINT,supplier_id UUID,created_at TIMESTAMP,updated_at TIMESTAMP,PRIMARY KEY (id));")

	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_supplier_id ON products(supplier_id);`)

	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS sales (
			id UUID,
			product_name varchar(255),
			product_id UUID,
			quantity BIGINT,
			total DOUBLE PRECISION,
			created_at TIMESTAMP,
			updated_at TIMESTAMP,
			PRIMARY KEY (id)
		);`)

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_product_id on sales(product_id);`)

	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
		 fullname varchar(255),
		 email varchar(255) UNIQUE,
		 id UUID,
		 password varchar(255),
		 created_at TIMESTAMP,
		 updated_at TIMESTAMP,
		 PRIMARY KEY (id)
	);`)

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_user_email on users(email);`)

	if err != nil {
		return err
	}

	return nil
}
