package main

import (
	"context"
	"niksmo/online-store/internal/generator"
	"niksmo/online-store/pkg/logger"
	"os"
	"os/signal"
)

const nWorkers = 8

func main() {
	stopCtx, stopFn := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stopFn()

	logger.Init()

	orderGenerator := generator.NewOrderGenerator(
		generator.NewProductStore(),
	)

	orderStream := orderGenerator.Run(stopCtx)
	generator.OrderSendersPool(stopCtx, nWorkers, orderStream, "http://127.0.0.1:8000/")

	<-stopCtx.Done()
}
