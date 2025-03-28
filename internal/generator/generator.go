package generator

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"math/rand/v2"
	"niksmo/online-store/pkg/counter"
	"niksmo/online-store/pkg/di"
	"niksmo/online-store/pkg/scheme"
	"time"
)

const (
	minUserID = 1
	maxUserID = 1_000
	minItems  = 1
	maxItems  = 3
)

type OrderGenerator struct {
	URL     string
	store   di.ProductGetter
	counter *counter.Counter
}

func NewOrderGenerator(URL string, store di.ProductGetter) *OrderGenerator {
	return &OrderGenerator{
		URL:     URL,
		store:   store,
		counter: counter.New(),
	}
}

func (og *OrderGenerator) Run(ctx context.Context, nWorkers int) {
	for workerID := range nWorkers {
		go func(ctx context.Context, orderGenerator *OrderGenerator) {
			for {
				data, err := json.MarshalIndent(og.MakeOrder(), "", "  ")
				if err != nil {
					panic(err)
				}
				fmt.Println(
					fmt.Sprintf("workerID=%d\n", workerID), string(data),
				)
				sleep := 100 + rand.IntN(1001-100)
				time.Sleep(time.Duration(sleep) * time.Millisecond)
			}
		}(ctx, og)
	}
}

func (og *OrderGenerator) MakeOrder() scheme.Order {
	items := og.randItems()
	totalPrice := og.sumItemsPrice(items)

	return scheme.Order{
		UserID:     og.randUserID(),
		OrderID:    og.getOrderID(),
		Items:      items,
		TotalPrice: totalPrice,
		Date:       time.Now(),
	}
}

func (og *OrderGenerator) randUserID() int {
	return minUserID + rand.IntN(maxUserID)
}

func (og *OrderGenerator) getOrderID() int {
	return int(og.counter.NextInt32())
}

func (og *OrderGenerator) randItems() []scheme.Product {
	itemsLen := rand.IntN(maxItems) + 1
	items := make([]scheme.Product, 0, itemsLen)
	for range itemsLen {
		items = append(items, og.store.GetRandProduct())
	}
	return items
}

func (og *OrderGenerator) sumItemsPrice(orderItems []scheme.Product) float64 {
	var sum float64
	for _, product := range orderItems {
		sum += product.Price
	}
	sum *= 100
	sum = math.Trunc(sum)
	sum /= 100
	return sum
}
