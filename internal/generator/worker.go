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

func sendOrderWorker(ctx context.Context, workerID int, orderStream <-chan scheme.Order, URL string) {
	for {
		select {
		case <-ctx.Done():
			return
		case order := <-orderStream:
			sendOrder(URL, order)

			logger.Instance.Info().
				Int("workerID", workerID).
				Int("userID", order.UserID).
				Int("orderID", order.OrderID).
				Msg("Send order")

			wait()
		}
	}
}

func sendOrder(URL string, order scheme.Order) {
	statusCode, _, errs := fiber.Post(URL).JSON(order).Bytes()
	if len(errs) != 0 {
		logger.Instance.Error().
			Caller().
			Int("statusCode", statusCode).
			Errs("errors", errs).Send()
	}
}

func wait() {
	interval := minSendInterval + rand.IntN((maxSendInterval+1)-minSendInterval)
	time.Sleep(time.Duration(interval) * time.Millisecond)
}
