package repository

import (
	"fmt"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"inventory/domain"
	"inventory/pkg/logger"
	"testing"
	"time"
)

func TestSaleRepository_Get(t *testing.T) {
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

func TestSaleRepository_Create(t *testing.T) {
	l, _ := logger.Init()
	initenv()

	repo := SetupDb(l)

	productId, err := uuid.FromString("21c0a2b2-c851-4a56-b3be-aed6eb3304ea")

	if err != nil {
		t.Error(err)
		return
	}

	sale, err := repo.SaleRepo.Create(domain.Sale{
		ProductName: "Nike",
		ProductId:   domain.UUID{UUID: productId},
		Quantity:    10,
		Total:       530.31,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(sale)
}

func TestSaleRepository_Update(t *testing.T) {
	l, _ := logger.Init()
	initenv()

	repo := SetupDb(l)

	sale, err := repo.SaleRepo.Update("86d9bacc-9471-470b-9873-39b30f65ef58", domain.Sale{
		Quantity: 10,
		Total:    530.31,
	})

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(sale)
}

func TestSaleRepository_Delete(t *testing.T) {
	l, _ := logger.Init()
	initenv()

	repo := SetupDb(l)

	err := repo.SaleRepo.Delete("86d9bacc-9471-470b-9873-39b30f65ef58")

	if err != nil {
		t.Error(err)
		return
	}
}

func TestSaleRepository_GetAll(t *testing.T) {
	l, _ := logger.Init()
	initenv()

	repo := SetupDb(l)

	sales, err := repo.SaleRepo.GetAll(2, 20)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(sales.Total, sales.Next)
}
