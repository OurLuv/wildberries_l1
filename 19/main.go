package main

import (
	"fmt"
	"strings"
)

func reverse(str string) string {
	var b strings.Builder
	// преобразовываем нашу строку в слайс рун
	runes := []rune(str)
	// проходимся по слайсу рун от конца до начала
	for i := len(runes) - 1; i >= 0; i-- {
		// кладём в билдер текущую руну
		b.WriteRune(runes[i])
	}
	// преобразовываем в string
	return b.String()
}

func main() {
	fmt.Print(reverse("🚗-car")) // rac-🚗
}
