package user

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"inventory/domain"
	"inventory/pkg/crypt"
	"inventory/pkg/logger"
)

type handler struct {
	repo   domain.UserRepository
	logger *zap.Logger
}

func New(router fiber.Router, repo domain.UserRepository) {
	handler := handler{
		repo: repo,
	}

	handler.logger, _ = logger.Init()

	router.Post("/signup", handler.Signup)
	router.Post("/signin", handler.Signin)
}

func (h handler) Signup(c *fiber.Ctx) error {
	var user domain.User

	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return domain.HandleError(c, err)
	}

	existingUser, err := h.repo.GetByEmail(user.Email)

	if err != nil {
		return domain.HandleError(c, err)
	}

	if existingUser != nil {
		return domain.HandleError(c, errors.New("user with that email already exists"))
	}

	user.Password, err = crypt.HashPassword(user.Password)

	if err != nil {
		return domain.HandleError(c, err)
	}

	newUser, err := h.repo.Create(user)

	if err != nil {
		return domain.HandleError(c, err)
	}

	token, err := crypt.GenerateToken(user)

	if err != nil {
		return domain.HandleError(c, err)
	}

	newUser.Password = ""

	return c.JSON(fiber.Map{
		"error": false,
		"data": fiber.Map{
			"user":  newUser,
			"token": token,
		},
	})
}

func (h handler) Signin(c *fiber.Ctx) error {
	var user domain.User

	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return domain.HandleError(c, err)
	}

	existingUser, err := h.repo.GetByEmail(user.Email)

	if err != nil {
		return domain.HandleError(c, err)
	}

	if existingUser == nil {
		return domain.HandleError(c, errors.New("invalid login credentials"))
	}

	if !crypt.CheckPasswordHash(user.Password, existingUser.Password) {
		return domain.HandleError(c, errors.New("invalid login credentials"))
	}

	token, err := crypt.GenerateToken(user)

	if err != nil {
		return domain.HandleError(c, err)
	}

	existingUser.Password = ""

	return c.JSON(fiber.Map{
		"error": false,
		"data": fiber.Map{
			"user":  existingUser,
			"token": token,
		},
	})
}
