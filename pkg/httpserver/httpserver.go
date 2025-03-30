package httpserver

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

const shutdownTimeout = 5 * time.Second

func Bootstrap(app *fiber.App, addr string) <-chan error {
	errCh := make(chan error)

	go func(app *fiber.App, addr string) {
		defer close(errCh)

		err := app.Listen(addr)
		if err != nil {
			errCh <- err
			return
		}
	}(app, addr)

	return errCh
}

func Close(app *fiber.App) error {
	return app.ShutdownWithTimeout(shutdownTimeout)
}
