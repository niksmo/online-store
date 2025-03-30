package store

import (
	"context"
	"niksmo/online-store/pkg/scheme"
)

type StoreService struct{}

func NewService() StoreService {
	return StoreService{}
}

func (s StoreService) CreateOrder(ctx context.Context, order scheme.Order) error {
	// producer behavior
	return nil
}
