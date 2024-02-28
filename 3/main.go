package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// первый способ - использование мьютексов
func withMutex(nums []int) {
	var sum int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, num := range nums {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			// блокируем доступ к переменной sum
			mu.Lock()
			sum += n * n
			// деблокируем доступ
			mu.Unlock()
		}(num)
	}
	wg.Wait()

	fmt.Printf("Сумма квадратов чисел: %d", sum)
}

// второй способ - использование atomic
func withAtomic(nums []int64) {
	var wg sync.WaitGroup
	var res int64
	for _, num := range nums {
		wg.Add(1)
		go func(n int64) {
			defer wg.Done()
			//потокобезопасная операция
			atomic.AddInt64(&res, n*n)
		}(num)
	}
	wg.Wait()
	fmt.Printf("Result: %d\n", res)
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	withMutex(nums)
	nums64 := []int64{2, 4, 6, 8, 10}
	withAtomic(nums64)
}
