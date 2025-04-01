package dashboard

import (
	"context"
	"niksmo/online-store/pkg/logger"
)

type DashboardService struct {
	consumer BatchMessageConsumer
}

func NewService(consumer BatchMessageConsumer) DashboardService {
	return DashboardService{consumer: consumer}
}

func (s DashboardService) Run(ctx context.Context) {
	logger.Instance.Info().Msg("dashboard service is running")

	s.consumer.Run(ctx)
}

func (s DashboardService) Close() error {
	logger.Instance.Info().Msg("close dashboard service")
	return s.consumer.kafkaC.Close()
}
