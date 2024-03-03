package main

import (
	"fmt"
	//"math/rand"
)

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	l, r := 0, len(arr)-1

	// разбиваем массив на элементы меньше и больше опорного. Опорный элемент - последний
	for i := range arr {
		if arr[i] < arr[r] {
			arr[i], arr[l] = arr[l], arr[i]
			l++
		}
	}

	// возвращаем опорный элемент на свое место, к "стене"
	// таким образом у нас получается два массива,
	// которые можно сортировать независимо друг от друго
	arr[l], arr[r] = arr[r], arr[l]

	// вызываем рекурсивно для двух частей массива
	quicksort(arr[:l])
	quicksort(arr[l+1:])
	return arr
}

func main() {
	arr := []int{8, 3, 4, 1, 5, 9, 7, 2, 6}
	fmt.Print(quicksort(arr))
}
