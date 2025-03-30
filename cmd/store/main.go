package main

import (
	"context"
	"niksmo/online-store/internal/store"
	"niksmo/online-store/pkg/logger"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

// flags
// RUN_ADDRESS
// kafka broker address
// kafka topic

func main() {

	stopCtx, stopFn := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stopFn()

	logger.Init()

	app := fiber.New()
	store.SetupAPIRouter(app)

	<-stopCtx.Done()
	logger.Instance.Info().Msg("Gracefully shutdown")
}
