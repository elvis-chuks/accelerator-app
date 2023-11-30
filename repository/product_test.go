package repository

import (
	"fmt"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"inventory/domain"
	"inventory/pkg/env"
	"inventory/pkg/logger"
	"testing"
)

func initenv() {
	env.LoadConfig("../")
}

func TestProductRepository_Get(t *testing.T) {
	//l, _ := logger.Init()
	//
	////mockDB, mock, err := sqlmock.New()
	////
	////if err != nil {
	////	t.Errorf("Error creating mock: %v", err)
	////	return
	////}
	////
	////defer func(mockDB *sql.DB) {
	////	err := mockDB.Close()
	////	if err != nil {
	////		return
	////	}
	////}(mockDB)
	////
	////mock.ExpectExec(fmt.Sprintf("SELECT * FROM products WHERE id ='%s'", "fd971b3b-8147-42f9-ac8b-07e1469d2820")).
	////	WillReturnResult(
	////		sqlmock.NewResult(1, 1))
	//
	////[]string{"name", "id", "price", "stock", "min_stock", "supplier_id", "created_at", "updated_at"}).
	////AddRow("Product_1", "fd971b3b-8147-42f9-ac8b-07e1469d2820", 53.31, 94, 86, "151afdd8-888b-4d5d-9b8c-3e6ae67cedc8", "2023-05-17 21:22:23.222091", "2023-05-17 21:22:23.222091"))
	//
	////result, err := NewProductRepository(mockDB, l).Get("fd971b3b-8147-42f9-ac8b-07e1469d2820")
	////
	////if err != nil {
	////	t.Error(err)
	////	return
	////}
	////
	////fmt.Println(result)
	//
	//mock, err := pgxmock.NewPool()
	//
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//
	//defer mock.Close()
	//
	//mock.ExpectBegin()
	//mock.ExpectQuery("SELECT * FROM products").WillReturnRows(pgxmock.NewRows([]string{"name", "id", "price", "stock", "min_stock", "supplier_id", "created_at", "updated_at"}).AddRow("Product_1", "fd971b3b-8147-42f9-ac8b-07e1469d2820", 53.31, 94, 86, "151afdd8-888b-4d5d-9b8c-3e6ae67cedc8", "2023-05-17 21:22:23.222091", "2023-05-17 21:22:23.222091"))
	////WillReturnResult(pgxmock.NewResult("SELECT", 1))
	//mock.ExpectCommit()
	//
	//result, err := NewProductRepository(mock, l).Get("fd971b3b-8147-42f9-ac8b-07e1469d2820")

}

func TestProductRepository_Create(t *testing.T) {
	l, _ := logger.Init()
	initenv()

	repo := SetupDb(l)

	supplierId, err := uuid.FromString("151afdd8-888b-4d5d-9b8c-3e6ae67cedc8")

	if err != nil {
		t.Error(err)
		return
	}

	product, err := repo.ProductRepo.Create(domain.Product{
		Name:       "Nike Bag",
		Price:      100,
		Stock:      10,
		MinStock:   5,
		SupplierId: domain.UUID{UUID: supplierId},
	})

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(product)
}

func TestProductRepository_Update(t *testing.T) {
	l, _ := logger.Init()
	initenv()

	repo := SetupDb(l)

	product, err := repo.ProductRepo.Update("35d24b1f-f0be-422f-bb66-9040d0356337", domain.Product{
		Name:     "Nike Bag 2",
		Price:    100,
		Stock:    90,
		MinStock: 50,
	})

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(product)
}

func TestProductRepository_Delete(t *testing.T) {
	l, _ := logger.Init()
	initenv()

	repo := SetupDb(l)

	err := repo.ProductRepo.Delete("13572cbd-bb98-4b3b-83bd-36397c2f27b6")

	if err != nil {
		t.Error(err)
		return
	}
}

func TestProductRepository_GetAll(t *testing.T) {
	l, _ := logger.Init()
	initenv()

	repo := SetupDb(l)

	products, err := repo.ProductRepo.GetAll(2, 20)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(products.Total, products.Next)
}
