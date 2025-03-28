package main

import (
	"context"
	"niksmo/online-store/internal/generator"
	"os"
	"os/signal"
)

const nWorkers = 8

func main() {
	stopCtx, stopFn := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stopFn()

	orderGenerator := generator.NewOrderGenerator(
		"URL",
		generator.NewProductStore(),
	)

	orderGenerator.Run(stopCtx, nWorkers)

	<-stopCtx.Done()
}
