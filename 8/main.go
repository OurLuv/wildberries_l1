package main

import (
	"flag"
	"fmt"
)

func main() {
	var n, shift int64
	var b, i int
	// создаём нужные флаги
	flag.Int64Var(&n, "n", 10, "number")
	flag.IntVar(&b, "b", 1, "bit")
	flag.IntVar(&i, "i", 2, "index")
	flag.Parse()

	fmt.Printf("%064b\n", n) // ...00001010
	// свдигаем 1 на i влево
	shift = 1 << i
	fmt.Printf("%064b\n", shift) // ...00000100
	// если b = 1, то нужно применить побитовое исключающее ИЛИ (^)
	if b == 1 {
		n = n ^ shift
		// если b = 0, то нужно применить побитовое И-НЕТ (&^)
	} else if b == 0 {
		n = n &^ shift
	}
	fmt.Printf("%064b\n", n) //  финальный резульат
}
