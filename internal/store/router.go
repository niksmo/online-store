package store

import (
	"context"
	"niksmo/online-store/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
)

func SetupAPIRouter(stopCtx context.Context, app *fiber.App) {

	router := app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger.Instance,
	}))

	storeService := NewService()
	storeHandler := NewHandler(storeService)
	storeProducer := NewProducer(stopCtx, "127.0.0.1:19094,127.0.0.1:29094", "orders_create")
	go storeService.CreatedOrdersStream(stopCtx, storeProducer)

	router.Post("/", storeHandler.PostOrder)
}
