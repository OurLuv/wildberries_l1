package main

import (
	"fmt"
	"reflect"
)

// 1ый способ
func checkType1(v interface{}) {
	// внутри switch мы можем использовать конструкция var.(type) которая возвращает тип
	switch v.(type) {
	case int:
		fmt.Printf("%v is int\n", v)
	case bool:
		fmt.Printf("%v is bool\n", v)
	case string:
		fmt.Printf("%v is string\n", v)
	case chan int:
		fmt.Printf("%v is chan int\n", v)
	}
}

// 2ой способ
func checkType2(v interface{}) {
	// создаем объект типа Value, представляющий значение интерфейса
	val := reflect.ValueOf(v)
	// возвращаеv тип значения, хранящегося внутри объекта Value
	t := val.Kind()
	fmt.Printf("%v is %s\n", v, t.String())
}

func main() {
	a := 5
	str := "Some text"
	checked := true
	var b *int
	ch := make(chan int)
	// 1
	checkType1(a)
	checkType1(str)
	checkType1(checked)
	checkType1(b)
	checkType1(ch)
	// 2
	checkType2(a)
	checkType2(str)
	checkType2(checked)
	checkType2(b)
	checkType2(ch)
}
