package store

import (
	"context"
	"niksmo/online-store/pkg/di"
	"niksmo/online-store/pkg/logger"
	"niksmo/online-store/pkg/scheme"
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

func (s StoreService) CreatedOrdersStream(
	ctx context.Context, orderProducer di.OrderProducer,
) {
	defer orderProducer.Close()

	for {
		select {
		case <-ctx.Done():
			return
		case order := <-s.orderStream:
			result, err := orderProducer.Produce(ctx, order)
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
