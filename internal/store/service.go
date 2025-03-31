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

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const orderStreamSize = 1024

type StoreService struct {
	orderStream chan scheme.Order
}

func NewService() StoreService {
	orderStream := make(chan scheme.Order, orderStreamSize)
	return StoreService{orderStream: orderStream}
}

func (s StoreService) CreateOrder(ctx context.Context, order scheme.Order) {
	select {
	case <-ctx.Done():
		return
	case s.orderStream <- order:
	default:
		logger.Instance.Warn().Msg("order stream is full")
	}
}

func (s StoreService) MessageStream(ctx context.Context) {
	producer, err := createProducer(ctx)
	if err != nil {
		logger.Instance.Error().Err(err).Msg("invalid producer config")
		return
	}
	defer producer.Close()

	for {
		select {
		case <-ctx.Done():
			return
		case order := <-s.orderStream:
			result, err := produce(ctx, producer, order)
			if err != nil {
				logger.Instance.Error().
					Err(err).
					Int("orderID", order.OrderID).
					Msg("order not produced")
				continue
			}
			logger.Instance.Info().
				Int("orderID", order.OrderID).
				Str("result", result).
				Msg("order produced")
		}
	}
}

func createProducer(ctx context.Context) (*kafka.Producer, error) {
	kConfig := &kafka.ConfigMap{
		"bootstrap.servers": "127.0.0.1:19094,127.0.0.1:29094",
	}
	return kafka.NewProducer(
		logkafka.Config(ctx, kConfig, 6, &logger.Instance),
	)
}

func produce(
	ctx context.Context, p *kafka.Producer, order scheme.Order,
) (string, error) {
	message, err := createMessage(order)
	if err != nil {
		return "", err
	}

	deliveryCh := make(chan kafka.Event)
	defer close(deliveryCh)

	err = p.Produce(message, deliveryCh)
	if err != nil {
		return "", err
	}

	return waitDeliveryResult(ctx, deliveryCh)
}

func createMessage(order scheme.Order) (*kafka.Message, error) {
	value, err := json.Marshal(order)
	if err != nil {
		logger.Instance.Error().Caller().Err(err).Send()
		return nil, err
	}

	headers := []kafka.Header{
		{Key: "userID", Value: serializer.Int(order.UserID)},
	}

	topic := "orders_create"
	partition := kafka.TopicPartition{
		Topic: &topic, Partition: kafka.PartitionAny,
	}

	msg := &kafka.Message{
		TopicPartition: partition,
		Value:          value,
		Headers:        headers,
	}

	return msg, nil
}

func waitDeliveryResult(ctx context.Context, deliveryCh chan kafka.Event) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case event := <-deliveryCh:
		result, err := handleDeliveryEvent(event)
		if err != nil {
			return "", err
		}
		return result, nil
	}
}

func handleDeliveryEvent(event kafka.Event) (string, error) {
	msg, ok := event.(*kafka.Message)
	if !ok {
		return "", errors.New("event is not kafka message")
	}

	if msg.TopicPartition.Error != nil {
		return "", msg.TopicPartition.Error
	}

	result := fmt.Sprintf(
		"delivered message to topic %s [%d] at offset %v\n",
		*msg.TopicPartition.Topic, msg.TopicPartition.Partition, msg.TopicPartition.Offset)

	return result, nil
}
