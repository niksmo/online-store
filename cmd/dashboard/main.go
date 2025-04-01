package main

import (
	"context"
	"niksmo/online-store/internal/dashboard"
	"niksmo/online-store/pkg/logger"
	"os"
	"os/signal"
)

func main() {
	stopCtx, stopFn := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stopFn()

	FlagsInit()
	logger.Init()

	consumer := dashboard.NewSingleMessageConsumer(
		stopCtx,
		KafkaServersFlagValue,
		KafkaTopicFlagValue,
		KafkaConsumerGroupFlagValue,
	)

	service := dashboard.NewService(consumer)

	go service.Run(stopCtx)

	<-stopCtx.Done()
	logger.Instance.Info().Msg("gracefully shutdown")

	if err := service.Close(); err != nil {
		logger.Instance.Error().Err(err).Caller().Msg("service close fail")
	}
}
