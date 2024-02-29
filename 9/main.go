package main

import "fmt"

func main() {
	nums := []int{5, 2, 10, 3, 8}
	ch1 := make(chan int)
	ch2 := make(chan int)

	// создаём горутину, в которой кладём в канал данные
	go func() {
		for _, num := range nums {
			ch1 <- num
		}
		close(ch1)
	}()

	// создаём горутину, в которой берём данные из первого канала
	// кладём во второй
	go func() {
		for num := range ch1 {
			ch2 <- num * 2
		}
		close(ch2)
	}()

	// нам не нужно использовать здесь никаких WaitGroup
	// т.к. мы используем range, программа не выполнится пока не обработает
	// все данные во втором канале
	for res := range ch2 {
		fmt.Println(res)
	}
}
