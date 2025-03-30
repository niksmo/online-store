package store

import (
	"niksmo/online-store/pkg/di"
	"niksmo/online-store/pkg/logger"
	"niksmo/online-store/pkg/scheme"

	"github.com/gofiber/fiber/v2"
)

type StoreHandler struct {
	service di.OrderCreater
}

func NewHandler(service di.OrderCreater) StoreHandler {
	return StoreHandler{service: service}
}

func (h StoreHandler) PostOrder(c *fiber.Ctx) error {
	log := logger.Instance.With().Caller().Logger()

	var order scheme.Order
	err := c.BodyParser(&order)
	if err != nil {
		log.Error().Err(err).Send()
		return fiber.ErrBadRequest
	}

	err = h.service.CreateOrder(c.Context(), order)
	if err != nil {
		log.Error().Err(err).Send()
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusCreated)
}
