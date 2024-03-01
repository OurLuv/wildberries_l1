package main

import "fmt"

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	mp := make(map[string]struct{})
	result := []string{}

	for _, v := range words {
		// проверяем есть ли такое слово в мапе
		// если есть, то возвращаемся в цикл
		if _, ok := mp[v]; ok {
			continue
		}
		// если нет, то добавляем слово в мапу
		// и в результат
		mp[v] = struct{}{}
		result = append(result, v)
	}
	fmt.Print(result)
}
