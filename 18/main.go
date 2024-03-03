package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Т.к. для счётчика нужны простые операции, то мы можем использовать не mutex,
// а атомики
type Counter struct {
	counter int64
}

func (c *Counter) Add(n int64) {
	atomic.AddInt64(&c.counter, n)
}

func (c *Counter) Get() int64 {
	return atomic.LoadInt64(&c.counter)
}

func startWorker(wg *sync.WaitGroup, c *Counter) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		c.Add(1)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	counter := &Counter{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go startWorker(wg, counter)
	}
	wg.Wait()
	fmt.Print(counter.Get())
}
