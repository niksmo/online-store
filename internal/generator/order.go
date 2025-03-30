package generator

import (
	"context"
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
	store   di.ProductGetter
	counter *counter.Counter
}

func NewOrderGenerator(store di.ProductGetter) *OrderGenerator {
	return &OrderGenerator{
		store:   store,
		counter: counter.New(),
	}
}

func (og *OrderGenerator) Run(ctx context.Context) <-chan scheme.Order {
	orderStream := make(chan scheme.Order)
	go og.sendOrderToStream(ctx, orderStream)
	return orderStream
}

func (og *OrderGenerator) sendOrderToStream(
	ctx context.Context, stream chan<- scheme.Order,
) {
	defer close(stream)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			stream <- og.MakeOrder()
		}
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
