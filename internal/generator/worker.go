package generator

import (
	"context"
	"math/rand/v2"
	"niksmo/online-store/pkg/logger"
	"niksmo/online-store/pkg/scheme"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	minSendInterval = 400
	maxSendInterval = 5000
)

func OrderSendersPool(
	ctx context.Context, n int, orderStream <-chan scheme.Order, URL string,
) {
	for id := range n {
		go sendOrderWorker(ctx, id, orderStream, URL)
	}
}

func sendOrderWorker(
	ctx context.Context,
	workerID int,
	orderStream <-chan scheme.Order,
	URL string,
) {
	for {
		select {
		case <-ctx.Done():
			return
		case order := <-orderStream:
			sendOrder(workerID, URL, order)
			wait()
		}
	}
}

func sendOrder(workerID int, URL string, order scheme.Order) {
	log := logger.Instance.With().
		Int("workerID", workerID).
		Int("userID", order.UserID).
		Int("orderID", order.OrderID).
		Logger()

	statusCode, _, errs := fiber.Post(URL).JSON(order).Bytes()

	if len(errs) != 0 {
		log.Error().Caller().Errs("errors", errs).Send()
		return
	}

	log.Info().Int("statusCode", statusCode).Msg("send order")
}

func wait() {
	interval := minSendInterval + rand.IntN(maxSendInterval-minSendInterval+1)
	time.Sleep(time.Duration(interval) * time.Millisecond)
}
