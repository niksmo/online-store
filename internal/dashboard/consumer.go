package dashboard

import (
	"context"
	"encoding/json"
	"niksmo/online-store/pkg/logger"
	"niksmo/online-store/pkg/logkafka"
	"niksmo/online-store/pkg/scheme"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	logKafkaLevel    = 6
	pollingTimeoutMs = 400
	fetchMinBytes    = 256
	fetchWaitMaxMs   = 4000
)

type BatchMessageConsumer struct {
	kafkaC *kafka.Consumer
	topic  string
}

func NewSingleMessageConsumer(
	ctx context.Context, bootstrapServers, topic, group string,
) BatchMessageConsumer {
	kConfig := &kafka.ConfigMap{
		"bootstrap.servers":     bootstrapServers,
		"group.id":              group,
		"heartbeat.interval.ms": 2000,
		"session.timeout.ms":    6000,
		"enable.auto.commit":    false,
		"fetch.min.bytes":       fetchMinBytes,
		"fetch.wait.max.ms":     fetchWaitMaxMs,
		"auto.offset.reset":     "earliest",
	}

	consumer, err := kafka.NewConsumer(
		logkafka.Config(ctx, kConfig, logKafkaLevel, &logger.Instance),
	)

	if err != nil {
		logger.Instance.Fatal().
			Err(err).
			Caller().
			Msg("invalid consumer config")
	}

	return BatchMessageConsumer{kafkaC: consumer, topic: topic}
}

func (c BatchMessageConsumer) Close() error {
	return c.kafkaC.Close()
}

func (c BatchMessageConsumer) Run(ctx context.Context) {
	err := c.kafkaC.Subscribe(c.topic, nil)
	if err != nil {
		logger.Instance.Fatal().
			Err(err).
			Caller().
			Msg("failed to subscribe on topic")
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
			event := c.kafkaC.Poll(pollingTimeoutMs)
			c.handleEvent(event)
		}
	}
}

func (c BatchMessageConsumer) handleEvent(event kafka.Event) {
	if event == nil {
		logger.Instance.Info().Msg("consumer polling return nil")
		return
	}

	switch ev := event.(type) {
	case *kafka.Message:
		c.handleMessage(ev)
	case kafka.Error:
		c.handleError(ev)
	case kafka.OffsetsCommitted:
		logger.Instance.Info().Msg("offset committed")
	default:
		logger.Instance.Info().Msg("kafka event ignored")
	}
}

func (c BatchMessageConsumer) handleMessage(kafkaMsg *kafka.Message) {
	var order scheme.Order
	err := json.Unmarshal(kafkaMsg.Value, &order)
	if err != nil {
		logger.Instance.Error().
			Err(err).
			Caller().
			Str("topicPartition", kafkaMsg.TopicPartition.String()).
			Msg("order scheme deserialization failed")
		return
	}

	logger.Instance.Info().
		RawJSON("order", kafkaMsg.Value).
		Msg("receive order in message")

	c.kafkaC.Commit()
}

func (c BatchMessageConsumer) handleError(kafkaErr kafka.Error) {
	logger.Instance.Info().
		Int("kafkaErrCode", int(kafkaErr.Code())).
		Str("error", kafkaErr.Error()).
		Caller().
		Msg("receive error")
}
