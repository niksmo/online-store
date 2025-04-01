package dashboard

import (
	"context"
	"niksmo/online-store/pkg/di"
	"niksmo/online-store/pkg/logger"
)

type DashboardService struct {
	consumer di.Consumer
}

func NewService(consumer di.Consumer) DashboardService {
	return DashboardService{consumer: consumer}
}

func (s DashboardService) Run(ctx context.Context) {
	logger.Instance.Info().Msg("dashboard service is running")

	s.consumer.Run(ctx)
}

func (s DashboardService) Close() error {
	logger.Instance.Info().Msg("close dashboard service")
	return s.consumer.Close()
}
