package ex8

import (
	"sync"
	"testing"
)

func TestCounterConcurrent(t *testing.T) {
	const goroutinesCount = 1000

	c := Counter{}
	var wg sync.WaitGroup
	wg.Add(goroutinesCount)

	for i := 0; i < goroutinesCount; i++ {
		go func() {
			c.Increment()
			wg.Done()
		}()
	}

	wg.Wait()

	if c.value != goroutinesCount {
		t.Errorf("got: %d, want: %d", c.value, goroutinesCount)
	}
}
