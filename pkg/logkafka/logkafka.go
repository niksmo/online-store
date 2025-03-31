package logkafka

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/rs/zerolog"
)

const defaultLevel = 6

func Config(
	ctx context.Context,
	kConfig *kafka.ConfigMap,
	level int,
	logger *zerolog.Logger,
) *kafka.ConfigMap {
	if level < 0 || level > 7 {
		level = defaultLevel
	}
	logEventStream := make(chan kafka.LogEvent)
	go handleKafkaLogEvent(ctx, logEventStream, logger)
	(*kConfig)["log_level"] = level
	(*kConfig)["go.logs.channel.enable"] = true
	(*kConfig)["go.logs.channel"] = logEventStream
	return kConfig
}

func handleKafkaLogEvent(
	ctx context.Context,
	logEventStream chan kafka.LogEvent,
	logger *zerolog.Logger,
) {
	defer close(logEventStream)
	for {
		select {
		case <-ctx.Done():
			return
		case logEvent := <-logEventStream:
			logger.Info().
				Str("name", logEvent.Name).
				Str("tag", logEvent.Tag).
				Int("kafkaLogLevel", logEvent.Level).
				Msg(logEvent.Message)
		}
	}
}
