package loyalty

import (
	"context"
	"niksmo/online-store/pkg/di"
	"niksmo/online-store/pkg/logger"
)

type LoyaltyService struct {
	consumer di.Consumer
}

func NewService(consumer di.Consumer) LoyaltyService {
	return LoyaltyService{consumer: consumer}
}

func (s LoyaltyService) Run(ctx context.Context) {
	logger.Instance.Info().Msg("loyalty service is running")

	s.consumer.Run(ctx)
}

func (s LoyaltyService) Close() error {
	logger.Instance.Info().Msg("close loyalty service")
	return s.consumer.Close()
}
