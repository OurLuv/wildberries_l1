package main

import "fmt"

type Human struct {
	Name  string
	Email string
}

func (h *Human) GetName() string {
	return h.Name
}

func (h *Human) GetEmail() string {
	return h.Email
}

type Action struct {
	Name string
	Human
}

func (a *Action) GetName() string {
	return a.Name
}

func main() {
	action := Action{
		Name: "Andrew",
		Human: Human{
			Name:  "Gabriel",
			Email: "gabriel@gmail.com",
		},
	}

	//* Обращение к полям объектов

	// поле Name присутствует в обеих структурах
	// если существуют одинаковые поля, приоритет будет отдан полю на более высоком уровне.
	fmt.Printf("Action's name: %s\n", action.Name) // Andrew

	// Чтобы обратиться к полю вложенной структуры, необходимо указать имя этого поля в самой структуре (Human),
	// а затем обратиться к нужному полю вложенной структуры (Name).
	fmt.Printf("Human's name: %s\n", action.Human.Name) // Gabriel

	//* Вызовы методов

	// структура Action имеет доступ к методам встроенной структуры Human. Данное свойство часто применяется в композиции
	fmt.Printf("Action's email: %s\n", action.GetEmail()) // gabriel@gmail.com

	// если есть методы с одинаковыми названиям, то приорите отдается методу из структуры выше.
	fmt.Printf("Action's name: %s\n", action.GetName()) // Andrew

	// так же как и с полями, если мы хотим вызвать метод встроенной стурктуры,то нам нужно указать само поле структуры
	fmt.Printf("Human's name: %s\n", action.Human.GetName()) // Gabriel

}
