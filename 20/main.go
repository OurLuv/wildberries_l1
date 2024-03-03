package main

import (
	"fmt"
	"strings"
)

func reverse(str string) string {
	// разбиваем строку на слова
	words := strings.Fields(str)
	result := []string{}

	// проходимся по слайсу слов с конца до начала
	for i := len(words) - 1; i >= 0; i-- {
		// кладём в результат текущее слово
		result = append(result, words[i])
	}

	// объединяем все слова в string, разделяем слова пробелом
	return strings.Join(result, " ")
}

func main() {
	fmt.Print(reverse("I like  the 🚗")) //🚗 the like I
}
