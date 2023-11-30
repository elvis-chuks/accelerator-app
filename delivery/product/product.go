package product

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"inventory/domain"
	"inventory/pkg/logger"
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
