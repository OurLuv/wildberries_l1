package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	if _, ok := a.SetString("480000000000000000000000", 10); !ok {
		panic("can't create a number")
	}

	b := new(big.Int)
	if _, ok := b.SetString("240000000000000000000000", 10); !ok {
		panic("can't create a number")
	}

	res := new(big.Int)
	fmt.Println("умножение:", res.Mul(a, b))
	fmt.Println("деление:", res.Div(a, b))
	fmt.Println("сложение:", res.Add(a, b))
	fmt.Println("вычитание:", res.Sub(a, b))
}
