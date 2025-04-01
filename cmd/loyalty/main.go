package main

import (
	"context"
	"niksmo/online-store/internal/loyalty"
	"niksmo/online-store/pkg/logger"
	"os"
	"os/signal"
)

func main() {
	stopCtx, stopFn := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stopFn()

	FlagsInit()
	logger.Init()

	consumer := loyalty.NewSingleMessageConsumer(
		stopCtx,
		KafkaServersFlagValue,
		KafkaTopicFlagValue,
		KafkaConsumerGroupFlagValue,
	)

	service := loyalty.NewService(consumer)

	go service.Run(stopCtx)

	<-stopCtx.Done()
	logger.Instance.Info().Msg("gracefully shutdown")

	if err := service.Close(); err != nil {
		logger.Instance.Error().Err(err).Caller().Msg("service close fail")
	}
}
