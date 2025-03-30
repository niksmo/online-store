package main

import (
	"context"
	"niksmo/online-store/internal/store"
	"niksmo/online-store/pkg/httpserver"
	"niksmo/online-store/pkg/logger"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

// flags
// kafka broker address
// kafka topic

func main() {
	stopCtx, stopFn := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stopFn()

	FlagsInit()
	logger.Init()

	app := fiber.New()
	store.SetupAPIRouter(stopCtx, app)

	serverErrCh := httpserver.Bootstrap(app, AddrFlagValue)
	go handleListenErr(serverErrCh)

	<-stopCtx.Done()
	logger.Instance.Info().Msg("gracefully shutdown")

	if err := httpserver.Close(app); err != nil {
		logger.Instance.Error().Caller().Err(err).Msg("server shutdown")
	}
}

func handleListenErr(errStream <-chan error) {
	if err := <-errStream; err != nil {
		logger.Instance.Error().Caller().Err(err).Msg("server listen")
	}
}
