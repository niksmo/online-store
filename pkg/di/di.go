package di

import (
	"context"
	"niksmo/online-store/pkg/scheme"
)

type ProductGetter interface {
	GetRandProduct() scheme.Product
}

type OrderCreater interface {
	CreateOrder(ctx context.Context, order scheme.Order) error
}
