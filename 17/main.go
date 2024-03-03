package main

import (
	"fmt"
)

// бинарный поиск будет работать только с отсортированным массивом данных
func binarySearch(arr []int, n int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		// находм середину рассматриваемого массива
		mid := (l + r) / 2
		// если наше значение больше, то рассматриваем подмассив справа от середины
		if arr[mid] < n {
			l = mid + 1
			// если меньше, то слева
		} else if arr[mid] > n {
			r = mid - 1
			// если значение равно искомому, то возвращаем его
		} else if arr[mid] == n {
			return n
		}
	}
	// возвращаем -1 если не смогли найти значение в массиве
	return -1
}

func main() {
	arr := []int{2, 3, 5, 10, 15, 20, 25}
	res := binarySearch(arr, 20)
	fmt.Print(res)
}
