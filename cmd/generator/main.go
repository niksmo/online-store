package main

import (
	"context"
	"niksmo/online-store/internal/generator"
	"niksmo/online-store/pkg/logger"
	"os"
	"os/signal"
)

func main() {
	stopCtx, stopFn := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stopFn()

	FlagsInit()
	logger.Init()

	orderGenerator := generator.NewOrderGenerator(
		generator.NewProductStore(),
	)

	orderStream := orderGenerator.Run(stopCtx)
	generator.OrderSendersPool(stopCtx, WorkersFlagValue, orderStream, AddrFlagValue)

	logger.Instance.Info().
		Str("AddrFlagValue", AddrFlagValue).
		Int("WorkersFlagValue", WorkersFlagValue).
		Msg("Start order generator application")

	<-stopCtx.Done()
	logger.Instance.Info().Msg("Gracefully shutdown")
}
