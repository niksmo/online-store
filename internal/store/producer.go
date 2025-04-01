package store

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"niksmo/online-store/pkg/logger"
	"niksmo/online-store/pkg/logkafka"
	"niksmo/online-store/pkg/scheme"
	"niksmo/online-store/pkg/serializer"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	deliveryMsgWaitTimeout = 3 * time.Second
	deliveryRetries        = 3
	logKafkaLevel          = 6
)

var ErrNotKafkaMsg = errors.New("event is not kafka message")

type OrdersCreateProducer struct {
	kafkaP *kafka.Producer
	topic  string
}

func NewProducer(
	ctx context.Context, bootstrapServers string, topic string,
) OrdersCreateProducer {
	kConfig := &kafka.ConfigMap{
		"bootstrap.servers":  bootstrapServers,
		"retries":            deliveryRetries,
		"acks":               "all",
		"enable.idempotence": true,
	}

	producer, err := kafka.NewProducer(
		logkafka.Config(ctx, kConfig, logKafkaLevel, &logger.Instance),
	)
	if err != nil {
		logger.Instance.Fatal().
			Err(err).
			Caller().
			Msg("invalid producer config")
	}
	return OrdersCreateProducer{kafkaP: producer, topic: topic}
}

func (p OrdersCreateProducer) Close() {
	p.kafkaP.Close()
}

func (p OrdersCreateProducer) Produce(
	ctx context.Context, order scheme.Order,
) (string, error) {
	message, err := p.createMessage(order)
	if err != nil {
		return "", err
	}

	deliveryCh := make(chan kafka.Event)
	defer close(deliveryCh)

	err = p.kafkaP.Produce(message, deliveryCh)
	if err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(ctx, deliveryMsgWaitTimeout)
	defer cancel()

	return p.waitDeliveryResult(ctx, deliveryCh)
}

func (p OrdersCreateProducer) createMessage(
	order scheme.Order,
) (*kafka.Message, error) {
	value, err := json.Marshal(order)
	if err != nil {
		logger.Instance.Error().Caller().Err(err).Send()
		return nil, err
	}
	logger.Instance.Info().RawJSON("msgPayload", value).Msg("prepare payload to kafka")

	partition := kafka.TopicPartition{
		Topic: &p.topic, Partition: kafka.PartitionAny,
	}

	msg := &kafka.Message{
		TopicPartition: partition,
		Key:            serializer.Int(order.UserID),
		Value:          value,
	}

	return msg, nil
}

func (p OrdersCreateProducer) waitDeliveryResult(
	ctx context.Context, deliveryCh chan kafka.Event,
) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case event := <-deliveryCh:
		result, err := p.handleDeliveryEvent(event)
		if err != nil {
			return "", err
		}
		return result, nil
	}
}

func (p OrdersCreateProducer) handleDeliveryEvent(
	event kafka.Event,
) (string, error) {
	msg, ok := event.(*kafka.Message)
	if !ok {
		return "", ErrNotKafkaMsg
	}

	if msg.TopicPartition.Error != nil {
		return "", msg.TopicPartition.Error
	}

	result := fmt.Sprintf(
		"delivered message to topic %s [%d] at offset %v\n",
		*msg.TopicPartition.Topic,
		msg.TopicPartition.Partition,
		msg.TopicPartition.Offset,
	)

	return result, nil
}
