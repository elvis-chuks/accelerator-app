package domain

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

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

func HandleValidationError(c *fiber.Ctx, err error) error {

	if _, ok := err.(*validator.InvalidValidationError); ok {

		return HandleError(c, err)
	}

	var errMessage string
	for _, err := range err.(validator.ValidationErrors) {
		errMessage = fmt.Sprintf("enter a valid %v in %v field", err.Kind().String(), err.Field())
		break
	}

	return HandleError(c, errors.New(errMessage))
}
