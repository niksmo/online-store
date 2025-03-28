package generator

import (
	"math/rand/v2"
	"niksmo/online-store/pkg/scheme"
	"niksmo/online-store/pkg/strgen"
	"sync"
)

const (
	minProductID       = 1
	maxProductID       = 1_000
	minProductPriceExp = 100
	maxProductPriceExp = 100_000
	productNameLength  = 8
)

type RandProductGetter struct {
	mu    sync.RWMutex
	index map[int]scheme.Product
}

func NewProductStore() *RandProductGetter {
	productGetter := &RandProductGetter{
		index: make(map[int]scheme.Product),
	}
	productGetter.init()
	return productGetter
}

func (pg *RandProductGetter) GetRandProduct() scheme.Product {
	pg.mu.RLock()
	defer pg.mu.RUnlock()

	product := pg.index[pg.randID()]
	return product
}

func (pg *RandProductGetter) init() {
	for id := minProductID; id <= maxProductID; id++ {
		pg.index[id] = scheme.Product{
			ID:    id,
			Name:  pg.randName(),
			Price: pg.randPrice(),
		}
	}
}

func (pg *RandProductGetter) randID() int {
	return rand.IntN(maxProductID) + 1
}

func (pg *RandProductGetter) randName() string {
	return strgen.Len(productNameLength)
}

func (pg *RandProductGetter) randPrice() float64 {
	exp := float64(
		minProductPriceExp + rand.IntN(maxProductPriceExp-minProductPriceExp+1),
	)
	frac := float64(rand.IntN(10)) / 10
	return exp + frac
}
