package main

//Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

import (
	"fmt"
)

// Определяем базовую структуру Human
type Human struct {
	Name string
	Age  int
}

// Метод для структуры Human
func (h Human) Greet() {
	fmt.Printf("Привет, меня зовут %s, мне %d лет.\n", h.Name, h.Age)
}

type Action struct {
	Human    // Встраиваем структуру Human
	Activity string
}

func (a Action) DoActivity() {
	fmt.Printf("%s занят %s.\n", a.Name, a.Activity)
}

func main() {
	a := Action{
		Human: Human{
			Name: "Иван",
			Age:  30,
		},
		Activity: "программированием",
	}

	a.Greet()      // Вызываем метод Greet() у встроенной структуры Human
	a.DoActivity() // Вызываем метод из Action
}
