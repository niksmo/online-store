package counter

import "sync/atomic"

type Counter struct {
	aint32 atomic.Int32
}

func New() *Counter {
	return &Counter{}
}

func (c *Counter) NextInt32() int32 {
	return c.aint32.Add(1)
}

func (c *Counter) ValueInt32() int32 {
	return c.aint32.Load()
}
