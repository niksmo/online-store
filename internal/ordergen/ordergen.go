package ordergen

import (
	"context"
)

type OrderGen struct {
	URL string
}

func New(URL string) OrderGen {
	return OrderGen{URL: URL}
}

func (g OrderGen) Run(ctx context.Context) {
	// agent := fiber.Post(g.URL).JSON()
}
