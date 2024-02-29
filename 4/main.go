package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func startWorker(wg *sync.WaitGroup, ctx context.Context, in <-chan int, n int) {
	defer wg.Done()
	var data int
	for {
		select {
		case data = <-in:
			fmt.Printf("worker %d recieved: %d\n\n", n, data)
		// останавливаем работу воркера, когда срабатывает cancel()
		case <-ctx.Done():
			fmt.Printf("shutting down worker №%d\n", n)
			return
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}
	// создаём флаг для количетсва воркеров -w=5
	workers := flag.Int("w", 1, "number of workers")
	flag.Parse()
	fmt.Printf("Number of workers: %d\n", *workers)
	wg.Add(*workers)

	// создаём в канал, в которым будем класть данные
	in := make(chan int)
	defer close(in)

	// создаём контекст, через который мы будем останавливать все наши горутины
	ctx, cancel := context.WithCancel(context.Background())

	// запускаем воркеры
	for i := 0; i < *workers; i++ {
		go startWorker(wg, ctx, in, i+1)
	}

	//кладём данные в канал
	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				data := rand.Intn(100) + 1
				in <- data
				fmt.Printf("message is sent: %d\n", data)
			// останавливаем работу messanger, когда срабатывает cancel()
			case <-ctx.Done():
				fmt.Println("shutting down messager")
				return
			}
		}
	}()

	// cоздаем канал shuttingDown для получения сигналов
	shuttingDown := make(chan os.Signal, 1)
	signal.Notify(shuttingDown, os.Interrupt, syscall.SIGTERM)
	// блокирующая операция, программа будет ждать сигнала перед завершением
	<-shuttingDown
	fmt.Println("application is shutting down")
	// отправляется сигнал остановки всем горутинам через контекст
	cancel()
	// ждём отработки всех горутин
	wg.Wait()
}
