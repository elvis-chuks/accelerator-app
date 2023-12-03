package product

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"inventory/domain"
	"inventory/pkg/env"
	"inventory/pkg/logger"
	"inventory/repository"
	"net/http"
	"net/http/httptest"
	"testing"
)

func initenv() {
	env.LoadConfig("../../")
}

func TestHandler_Create(t *testing.T) {

	initenv()

	l, _ := logger.Init()

	supplierId, err := uuid.FromString("151afdd8-888b-4d5d-9b8c-3e6ae67cedc8")

	if err != nil {
		t.Error(err)
		return
	}

	body := domain.Product{
		Name:       "Nike Bag",
		Price:      100,
		Stock:      10,
		MinStock:   5,
		SupplierId: domain.UUID{UUID: supplierId},
	}

	jsonStr, err := json.Marshal(body)

	if err != nil {
		t.Error(err)
		return
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/product", bytes.NewBuffer(jsonStr))

	repo := repository.SetupDb(l)

	app := fiber.New()

	router := app.Group("/api/v1/product")

	New(router, repo.ProductRepo)

	resp, err := app.Test(req, 10000)

	if err != nil {
		t.Error(err)
		return
	}

	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode, "should pass")

}

func TestHandler_Get(t *testing.T) {
	initenv()

	l, _ := logger.Init()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/product/41bc611e-ace6-441c-a374-570ce63d005a", nil)

	repo := repository.SetupDb(l)

	app := fiber.New()

	router := app.Group("/api/v1/product")

	New(router, repo.ProductRepo)

	resp, err := app.Test(req, 10000)

	if err != nil {
		t.Error(err)
		return
	}

	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode, "should pass")
}

func TestHandler_GetAll(t *testing.T) {
	initenv()

	l, _ := logger.Init()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/product?page=1&limit=10", nil)

	repo := repository.SetupDb(l)

	app := fiber.New()

	router := app.Group("/api/v1/product")

	New(router, repo.ProductRepo)

	resp, err := app.Test(req, 10000)

	if err != nil {
		t.Error(err)
		return
	}

	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode, "should pass")
}

func TestHandler_Update(t *testing.T) {

	initenv()

	l, _ := logger.Init()

	body := domain.Product{
		Name:     "Nike Bag",
		Price:    100,
		Stock:    10,
		MinStock: 5,
	}

	jsonStr, err := json.Marshal(body)

	if err != nil {
		t.Error(err)
		return
	}

	req := httptest.NewRequest(http.MethodPut, "/api/v1/product/41bc611e-ace6-441c-a374-570ce63d005a", bytes.NewBuffer(jsonStr))

	repo := repository.SetupDb(l)

	app := fiber.New()

	router := app.Group("/api/v1/product")

	New(router, repo.ProductRepo)

	resp, err := app.Test(req, 10000)

	if err != nil {
		t.Error(err)
		return
	}

	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode, "should pass")

}

func TestHandler_GetRecommendation(t *testing.T) {
	initenv()

	l, _ := logger.Init()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/product/recommend/restock", nil)

	repo := repository.SetupDb(l)

	app := fiber.New()

	router := app.Group("/api/v1/product")

	New(router, repo.ProductRepo)

	resp, err := app.Test(req, 10000)

	if err != nil {
		t.Error(err)
		return
	}

	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode, "should pass")
}

func TestHandler_Delete(t *testing.T) {
	initenv()

	l, _ := logger.Init()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/product/bf98b509-65f7-4fbe-8277-6ddfbd00ec00", nil)

	repo := repository.SetupDb(l)

	app := fiber.New()

	router := app.Group("/api/v1/product")

	New(router, repo.ProductRepo)

	resp, err := app.Test(req, 10000)

	if err != nil {
		t.Error(err)
		return
	}

	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode, "should pass")
}
