package main

import "fmt"

type Human interface {
	Greeting()
}

// адаптируемый тип
type Russian struct {
	Name string
}

func (r *Russian) GreetingInRussian() {
	fmt.Printf("Привет! Я %s", r.Name)
}

// адаптер
type Adapter struct {
	rus Russian
}

// адаптируем метод так, чтобы адаптер реализовывал интерфейс Human
func (a *Adapter) Greeting() {
	a.rus.GreetingInRussian()
}

func main() {
	var human Human = &Adapter{
		rus: Russian{"Андрей"},
	}
	human.Greeting() // Привет! Я Андрей
}
