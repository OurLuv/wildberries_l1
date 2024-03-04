package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func createHugeString(size int) string {
	var b strings.Builder
	for i := 0; i < size; i++ {
		fmt.Fprint(&b, "⛄")
	}
	return b.String()
}

func someFunc() {
	// создаём огромную строку
	v := createHugeString(1 << 10)

	justString := string(v[:100])
	fmt.Print(justString)
}
func main() {
	fmt.Println(utf8.RuneLen('⛄')) // 3
	someFunc()
}
