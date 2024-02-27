package main

import (
	"fmt"
	"sync"
)

// первый способ - для каждого числа вызываем горутину, которая будет считать квадрат числа
func square(nums []int) {
	// создаём WaitGroup для синхронизации приложения
	wg := &sync.WaitGroup{}
	// количество отработанных итераций/wg.Done(), которых будет ждать wg.Wait() будет равно число элементов в массиве
	wg.Add(len(nums))

	// проходимся по каждому элементу
	for _, v := range nums {
		// вызываем горутину и передаём в неё число из массива
		go func(n int) {
			//по заверщению работы горутины выполняем wg.Done()
			defer wg.Done()
			// выводит квадрат числа
			fmt.Printf("%d\n", n*n)
		}(v)
	}

	// ждём вызова всех wg.Done()
	wg.Wait()
}

// второй способ - вызов специальных воркеров. Данные приходят в канал, после чего производятся нужные операции в методе
func startWorker(c <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range c {
		fmt.Printf("%d\n", i*i)
	}
}

func squareWorker(nums []int) {
	amountOfWorkers := 2
	// создаём канал
	in := make(chan int)
	wg := &sync.WaitGroup{}
	// wg.Add() должен быть не больше количества наших воркеров иначе будет deadlock
	wg.Add(amountOfWorkers)
	// в цикле запускаем amountOfWorkers воркеров
	for i := 0; i < amountOfWorkers; i++ {
		go startWorker(in, wg)
	}
	// кладём все числа из массива в канал
	for i := range nums {
		in <- nums[i]
	}

	// закрываем канал и ждём, когда все наши воркеры отработают
	close(in)
	wg.Wait()
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	square(nums)
	squareWorker(nums)
}
