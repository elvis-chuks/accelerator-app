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
	l, _ := logger.Init()
	initenv()

	repo := SetupDb(l)

	product, err := repo.ProductRepo.Get("35d24b1f-f0be-422f-bb66-9040d0356337")

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(product)
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

func TestProductRepository_GetRestockRecommendation(t *testing.T) {
	l, _ := logger.Init()
	initenv()

	repo := SetupDb(l)

	products, err := repo.ProductRepo.GetRestockRecommendation(1, 10)

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(products)

}
