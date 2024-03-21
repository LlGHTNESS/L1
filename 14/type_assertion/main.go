package main

import (
	"fmt"
)

func determineType(i interface{}) {
	if v, ok := i.(int); ok {
		fmt.Println("Переменная имеет тип int:", v)
	} else if v, ok := i.(string); ok {
		fmt.Println("Переменная имеет тип string:", v)
	} else if v, ok := i.(bool); ok {
		fmt.Println("Переменная имеет тип bool:", v)
	} else if _, ok := i.(chan int); ok {
		fmt.Println("Переменная имеет тип chan int")
	} else {
		fmt.Println("Неизвестный тип")
	}
}

func main() {
	var myInt int = 42
	var myString string = "Hello, world"
	var myBool bool = true
	var myChannel chan int = make(chan int)

	// Проверяем тип каждой переменной с помощью функции determineType.
	determineType(myInt)
	determineType(myString)
	determineType(myBool)
	determineType(myChannel)
}
