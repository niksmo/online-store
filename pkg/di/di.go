package di

import (
	"niksmo/online-store/pkg/scheme"
)

type ProductGetter interface {
	GetRandProduct() scheme.Product
}
