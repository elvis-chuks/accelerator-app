package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"inventory/delivery/product"
	"inventory/delivery/sale"
	"inventory/domain"
)

type Config struct {
	ProductRepo domain.ProductRepository
	SaleRepo    domain.SaleRepository
}

func SetupRouter(config Config) *fiber.App {
	app := fiber.New()
	app.Use(cors.New())

	v1 := app.Group("/api/v1")

	productRouter := v1.Group("/product")
	saleRouter := v1.Group("/sale")

	product.New(productRouter, config.ProductRepo)
	sale.New(saleRouter, config.SaleRepo)
	return app
}
