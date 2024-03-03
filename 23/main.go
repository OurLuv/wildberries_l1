package main

import "fmt"

// если порядок не важен
func removeNoOrder[T any](arr []T, i int) []T {
	// копируем последний элемент в индекс
	arr[i] = arr[len(arr)-1]
	//возвращаем слайс без последнего элемента
	return arr[:len(arr)-1]
}

// есл порядок важен
func remove[T any](arr []T, i int) []T {
	// повторный слайсинг
	return append(arr[:i], arr[i+1:]...)
}

func main() {
	arr := []string{"a", "b", "c", "d"}
	fmt.Println(removeNoOrder(arr, 0)) // [d b c]

	arr1 := []string{"a", "b", "c", "d"}
	fmt.Println(remove(arr1, 1)) // [a c d]
}
