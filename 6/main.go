package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg *sync.WaitGroup

// 1. с использованием контекста context.WithCancel / context.WithTimeout
func stopWithContext(ctx context.Context, in <-chan int) {
	defer wg.Done()
	for {
		select {
		case data := <-in:
			fmt.Printf("stopWithContext got: %d\n", data)
		case <-ctx.Done():
			fmt.Print("stopWithContext is shutting down \n")
			return
		}
	}
}

// 2. наша горутина будет проверять закрыт ли канал или нет
func stopCheckIfClosed(in <-chan int) {
	defer wg.Done()
	for {
		if _, ok := <-in; !ok {
			fmt.Print("stopCheckIfClosed is shutting down\n")
			return
		}
		fmt.Printf("stopCheckIfClosed got: %d\n", <-in)
	}
}

// 3. когда канал будет пуст - итерации в цикле закончатся и метод сам остановится
func stopChanRange(in <-chan int) {
	defer wg.Done()
	for v := range in {
		fmt.Printf("stopChanRange got: %d\n", v)
	}
	fmt.Print("stopChanRange is shutting down\n")
}

// 4. Используем отдельный канал для остановки горутины
func stopChanStop(in <-chan int, st <-chan struct{}) {
	defer wg.Done()
	for {
		select {
		case data := <-in:
			fmt.Printf("stopChanStop got: %d\n", data)
		case <-st:
			fmt.Printf("stopChanStop is shutting down\n")
			return
		}
	}
}

// отправка данных в канал
func publish(ctx context.Context, in chan int) {
	defer wg.Done()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			data := rand.Intn(100) + 1
			in <- data
		case <-ctx.Done():
			fmt.Println("publisher is shutting down")
			return
		}
	}
}

func main() {
	wg = &sync.WaitGroup{}
	in := make(chan int)

	//с контекстом
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(2)
	go publish(ctx, in)
	go stopWithContext(ctx, in)
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()

	// с проверкой на закрытый канал
	ctx, cancel = context.WithCancel(context.Background())
	in = make(chan int)
	wg.Add(2)
	go publish(ctx, in)
	go stopCheckIfClosed(in)
	time.Sleep(5 * time.Second)
	close(in)
	cancel()
	wg.Wait()

	// с range
	ctx, cancel = context.WithCancel(context.Background())
	in = make(chan int)
	wg.Add(2)
	go publish(ctx, in)
	go stopChanRange(in)
	time.Sleep(5 * time.Second)
	close(in)
	cancel()
	wg.Wait()

	// с отдельным каналом для остановки
	ctx, cancel = context.WithCancel(context.Background())
	in = make(chan int)
	st := make(chan struct{})
	wg.Add(2)
	go publish(ctx, in)
	go stopChanStop(in, st)
	time.Sleep(5 * time.Second)
	st <- struct{}{}
	close(in)
	close(st)
	cancel()
	wg.Wait()

}
