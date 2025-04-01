package di

import (
	"context"
	"niksmo/online-store/pkg/scheme"
)

type ProductGetter interface {
	GetRandProduct() scheme.Product
}

type OrderCreater interface {
	CreateOrder(ctx context.Context, order scheme.Order)
}

// return result if error is nil
type OrderProducer interface {
	Produce(ctx context.Context, order scheme.Order) (string, error)
	Close()
}

type Consumer interface {
	Run(ctx context.Context)
	Close() error
}
