package product

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"inventory/domain"
	"inventory/pkg/logger"
	"strconv"
)

type handler struct {
	repo   domain.ProductRepository
	logger *zap.Logger
}

func New(router fiber.Router, repo domain.ProductRepository) {
	handler := handler{
		repo: repo,
	}

	handler.logger, _ = logger.Init()

	router.Get("/:id", handler.Get)
	router.Delete("/:id", handler.Delete)
	router.Get("/", handler.GetAll)
	router.Put("/:id", handler.Update)
	router.Post("/", handler.Create)
	router.Get("/recommend/restock", handler.RecommendRestock)
}

func (h handler) Get(c *fiber.Ctx) error {
	id := c.Params("id")

	product, err := h.repo.Get(id)

	if err != nil {
		return domain.HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data": fiber.Map{
			"product": product,
		},
	})
}

func (h handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.repo.Delete(id)

	if err != nil {
		return domain.HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"error": false,
	})
}

func (h handler) GetAll(c *fiber.Ctx) error {

	pageStr := c.Query("page", "1")
	limitStr := c.Query("limit", "10")

	page, err := strconv.ParseInt(pageStr, 10, 64)

	if err != nil {
		return domain.HandleError(c, err)
	}

	limit, err := strconv.ParseInt(limitStr, 10, 64)

	if err != nil {
		return domain.HandleError(c, err)
	}

	response, err := h.repo.GetAll(page, limit)

	if err != nil {
		return domain.HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data":  response,
	})
}

func (h handler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var product domain.Product

	if err := json.Unmarshal(c.Body(), &product); err != nil {
		return domain.HandleError(c, err)
	}

	response, err := h.repo.Update(id, product)

	if err != nil {
		return domain.HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data":  response,
	})
}

func (h handler) Create(c *fiber.Ctx) error {

	var product domain.Product

	if err := json.Unmarshal(c.Body(), &product); err != nil {
		return domain.HandleError(c, err)
	}

	err := validator.New().Struct(product)

	if err != nil {
		return domain.HandleValidationError(c, err)
	}

	response, err := h.repo.Create(product)

	if err != nil {
		return domain.HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data":  response,
	})
}

func (h handler) RecommendRestock(c *fiber.Ctx) error {

	pageStr := c.Query("page", "1")
	limitStr := c.Query("limit", "10")

	page, err := strconv.ParseInt(pageStr, 10, 64)

	if err != nil {
		return domain.HandleError(c, err)
	}

	limit, err := strconv.ParseInt(limitStr, 10, 64)

	if err != nil {
		return domain.HandleError(c, err)
	}

	products, err := h.repo.GetRestockRecommendation(page, limit)

	return c.JSON(fiber.Map{
		"error": false,
		"data":  products,
	})
}
