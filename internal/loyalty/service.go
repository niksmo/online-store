package loyalty

import (
	"context"
	"niksmo/online-store/pkg/logger"
)

type LoyaltyService struct {
	consumer SingleMessageConsumer
}

func NewService(consumer SingleMessageConsumer) LoyaltyService {
	return LoyaltyService{consumer: consumer}
}

func (s LoyaltyService) Run(ctx context.Context) {
	logger.Instance.Info().Msg("loyalty service is running")

	s.consumer.Run(ctx)
}

func (s LoyaltyService) Close() error {
	logger.Instance.Info().Msg("close loyalty service")
	return s.consumer.kafkaC.Close()
}
