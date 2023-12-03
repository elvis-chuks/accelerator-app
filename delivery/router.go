package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"inventory/delivery/product"
	"inventory/delivery/sale"
	"inventory/delivery/user"
	"inventory/domain"
	"inventory/pkg/middleware"
)

type Config struct {
	ProductRepo domain.ProductRepository
	SaleRepo    domain.SaleRepository
	UserRepo    domain.UserRepository
}

func SetupRouter(config Config) *fiber.App {
	app := fiber.New()
	app.Use(cors.New())

	v1 := app.Group("/api/v1")

	productRouter := v1.Group("/product", middleware.Protected())
	saleRouter := v1.Group("/sale", middleware.Protected())
	authRouter := v1.Group("auth")

	product.New(productRouter, config.ProductRepo)
	sale.New(saleRouter, config.SaleRepo)
	user.New(authRouter, config.UserRepo)
	return app
}
