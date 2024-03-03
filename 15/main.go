package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func createHugeString(size int) string {
	var b strings.Builder
	for i := 0; i < size; i++ {
		fmt.Fprint(&b, "🚗")
	}
	return b.String()
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	// в данном случае мы срезаем по количеству байт, а не по количеству рун
	justString = v[:100]

	// преобразовываем строку в слайс рун
	r := []rune(v)
	// в даннам случае мы срезаем по количеству рун
	justString = string(r[:100])
}
func main() {
	fmt.Println(utf8.RuneLen('🚗')) // 3
	//someFunc()
}
