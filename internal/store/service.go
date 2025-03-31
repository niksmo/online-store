package store

import (
	"context"
	"fmt"
	"niksmo/online-store/pkg/logger"
	"niksmo/online-store/pkg/logkafka"
	"niksmo/online-store/pkg/scheme"

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
	kConfig := &kafka.ConfigMap{
		"bootstrap.servers": "127.0.0.1:19094,127.0.0.1:29094",
	}

	producer, err := kafka.NewProducer(
		logkafka.Config(ctx, kConfig, 6, &logger.Instance),
	)

	if err != nil {
		logger.Instance.Error().Err(err).Msg("broken producer config")
		return
	}

	logger.Instance.Info().Str("producer", fmt.Sprintf("%v", producer)).Msg("create producer")

	for {
		select {
		case <-ctx.Done():
			return
		case order := <-s.orderStream:
			logger.Instance.Info().Int("orderID", order.OrderID).Msg("produce order")
		}
	}
}
