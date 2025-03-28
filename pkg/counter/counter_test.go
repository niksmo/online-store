package counter_test

import (
	"niksmo/online-store/pkg/counter"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	cntr := counter.New()

	n := 1_000
	nWorkers := 10
	ch := make(chan struct{}, n)

	var wg sync.WaitGroup
	for range nWorkers {
		wg.Add(1)
		go func(cntr *counter.Counter, ch <-chan struct{}) {
			defer wg.Done()
			for {
				select {
				case <-ch:
					cntr.NextInt32()
				default:
					return
				}
			}
		}(cntr, ch)
	}

	for range n {
		ch <- struct{}{}
	}
	wg.Wait()

	assert.Equal(t, int32(n), cntr.ValueInt32())
}
