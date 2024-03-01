package main

import "fmt"

func main() {
	a := 5
	b := 10

	a, b = b, a
	fmt.Printf("a: %d, b: %d", a, b) // a: 10, b: 5
}
