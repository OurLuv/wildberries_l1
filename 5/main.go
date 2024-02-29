package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func startWorker(wg *sync.WaitGroup, in <-chan int) {
	defer wg.Done()
	for v := range in {
		fmt.Printf("worker got: %d\n\n", v)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	// создаём флаг duration -d=10s
	var duration time.Duration
	flag.DurationVar(&duration, "d", 30*time.Second, "a duration value")
	flag.Parse()

	workers := 2
	wg.Add(workers)

	// создаём в канал, в который будем класть данные
	in := make(chan int)
	// запускаем воркеры
	for i := 0; i < workers; i++ {
		go startWorker(wg, in)
	}

	//кладём данные в канал
	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()
		// создаём таймер, в который кладём наш duration
		timer := time.NewTimer(duration)
		defer timer.Stop()

		for {
			select {
			case <-ticker.C:
				data := rand.Intn(100) + 1
				in <- data
				fmt.Printf("message is sent: %d\n", data)
			// когда приходит сигнал таймера - закрываем канал и завершаем работу горутины
			//
			case <-timer.C:
				fmt.Println("shutting down an application")
				close(in)
				return
			}
		}
	}()
	// ждём отработки всех горутин
	wg.Wait()
}
