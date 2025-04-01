package main

import (
	"context"
	"niksmo/online-store/internal/store"
	"niksmo/online-store/pkg/httpserver"
	"niksmo/online-store/pkg/logger"
	"os"
	"os/signal"
)

func main() {
	stopCtx, stopFn := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stopFn()

	FlagsInit()
	logger.Init()

	server := httpserver.New(AddrFlagValue)

	store.SetupAPIRouter(
		stopCtx, server.FiberApp, KafkaServersFlagValue, KafkaTopicFlagValue,
	)

	go server.Listen(func(serverErr error) {
		if serverErr != nil {
			logger.Instance.Error().Caller().Err(serverErr).Msg("server listen")
		}
	})

	<-stopCtx.Done()
	logger.Instance.Info().Msg("gracefully shutdown")

	if err := server.Close(); err != nil {
		logger.Instance.Error().Caller().Err(err).Msg("server shutdown")
	}
}
