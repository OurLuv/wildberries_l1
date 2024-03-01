package main

import "fmt"

func intersection(set1 []int, set2 []int) []int {
	mp := make(map[int]int)
	res := []int{}
	for _, v := range set1 {
		mp[v]++
	}
	for _, v := range set2 {
		mp[v]++
		// если один из элементов множества встречается 2 раза -
		// добавляем его в результат
		if mp[v] == 2 {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	set1 := []int{5, 2, 1, 7, 6}
	set2 := []int{7, 2, 10, 4, 3}
	result := intersection(set1, set2)
	fmt.Print(result)
}
