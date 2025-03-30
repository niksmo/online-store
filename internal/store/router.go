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
	go storeService.MessageStream(stopCtx)

	router.Post("/", storeHandler.PostOrder)
}
