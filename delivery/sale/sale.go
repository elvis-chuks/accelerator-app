package sale

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"inventory/domain"
	"inventory/pkg/logger"
	"strconv"
)

type handler struct {
	repo   domain.SaleRepository
	logger *zap.Logger
}

func New(router fiber.Router, repo domain.SaleRepository) {
	handler := handler{
		repo: repo,
	}

	handler.logger, _ = logger.Init()

	router.Get("/:id", handler.Get)
	router.Delete("/:id", handler.Delete)
	router.Get("/", handler.GetAll)
	router.Put("/:id", handler.Update)
	router.Post("/", handler.Create)
}

func (h handler) Get(c *fiber.Ctx) error {
	id := c.Params("id")

	response, err := h.repo.Get(id)

	if err != nil {
		return domain.HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data":  response,
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

	var sale domain.Sale

	if err := json.Unmarshal(c.Body(), &sale); err != nil {
		return domain.HandleError(c, err)
	}

	response, err := h.repo.Update(id, sale)

	if err != nil {
		return domain.HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data":  response,
	})
}

func (h handler) Create(c *fiber.Ctx) error {

	var sale domain.Sale

	if err := json.Unmarshal(c.Body(), &sale); err != nil {
		return domain.HandleError(c, err)
	}

	response, err := h.repo.Create(sale)

	if err != nil {
		return domain.HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data":  response,
	})
}
