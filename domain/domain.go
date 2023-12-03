package domain

import "github.com/gofiber/fiber/v2"

type Repository struct {
	ProductRepo ProductRepository
	SaleRepo    SaleRepository
	UserRepo    UserRepository
}

func HandleError(c *fiber.Ctx, err error) error {
	return c.Status(400).JSON(
		fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
}
