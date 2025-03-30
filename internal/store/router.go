package store

import (
	"niksmo/online-store/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
)

func SetupAPIRouter(app *fiber.App) {

	router := app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger.Instance,
	}))

	storeHandler := NewHandler(NewService())

	router.Post("/", storeHandler.PostOrder)
}
