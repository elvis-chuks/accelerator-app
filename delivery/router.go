package delivery

import (
	"github.com/gofiber/fiber/v2"
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

	return app
}
