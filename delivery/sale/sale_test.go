package sale

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

	productId, err := uuid.FromString("21c0a2b2-c851-4a56-b3be-aed6eb3304ea")

	if err != nil {
		t.Error(err)
		return
	}

	sale := domain.Sale{
		ProductName: "Nike",
		ProductId:   domain.UUID{UUID: productId},
		Quantity:    10,
		Total:       530.31,
	}

	jsonStr, err := json.Marshal(sale)

	if err != nil {
		t.Error(err)
		return
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/sale", bytes.NewBuffer(jsonStr))

	repo := repository.SetupDb(l)

	app := fiber.New()

	router := app.Group("/api/v1/sale")

	New(router, repo.SaleRepo)

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

	req := httptest.NewRequest(http.MethodGet, "/api/v1/sale/330700b9-793b-4f04-85c0-1fb417e718d2", nil)

	repo := repository.SetupDb(l)

	app := fiber.New()

	router := app.Group("/api/v1/sale")

	New(router, repo.SaleRepo)

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

	req := httptest.NewRequest(http.MethodGet, "/api/v1/sale?page=1&limit=10", nil)

	repo := repository.SetupDb(l)

	app := fiber.New()

	router := app.Group("/api/v1/sale")

	New(router, repo.SaleRepo)

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

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/sale/330700b9-793b-4f04-85c0-1fb417e718d2", nil)

	repo := repository.SetupDb(l)

	app := fiber.New()

	router := app.Group("/api/v1/sale")

	New(router, repo.SaleRepo)

	resp, err := app.Test(req, 10000)

	if err != nil {
		t.Error(err)
		return
	}

	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode, "should pass")
}
