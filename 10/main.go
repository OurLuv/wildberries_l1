package main

import (
	"fmt"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// создаём мапу в которой ключ - число(десятки) и значение - температура
	groups := make(map[int][]float64)
	for _, v := range temperatures {
		//округляем значение до ближайшего кратного 10ти
		k := int(v/10) * 10
		groups[k] = append(groups[k], v)
	}

	fmt.Printf("%v\n", groups)
}
