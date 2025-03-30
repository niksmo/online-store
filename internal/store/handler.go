package store

import (
	"niksmo/online-store/pkg/logger"
	"niksmo/online-store/pkg/scheme"

	"github.com/gofiber/fiber/v2"
)

type StoreHandler struct {
}

func NewHandler() StoreHandler {
	return StoreHandler{}
}

func (h *StoreHandler) CreateOrder(c *fiber.Ctx) error {
	var order scheme.Order
	err := c.BodyParser(&order)
	if err != nil {
		logger.Instance.Error().Caller().Err(err).Send()
		return fiber.ErrBadRequest
	}

	// send to channel for async processing

	return c.SendStatus(fiber.StatusCreated)
}
