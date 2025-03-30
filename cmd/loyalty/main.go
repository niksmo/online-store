package main

import (
	"context"
	"niksmo/online-store/pkg/logger"
	"os"
	"os/signal"
)

func main() {
	stopCtx, stopFn := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stopFn()

	logger.Init()

	<-stopCtx.Done()
	logger.Instance.Info().Msg("gracefully shutdown")
}
