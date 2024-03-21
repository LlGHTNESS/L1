package main

import (
	"fmt"
)

func determineType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Переменная имеет тип int:", v)
	case string:
		fmt.Println("Переменная имеет тип string:", v)
	case bool:
		fmt.Println("Переменная имеет тип bool:", v)
	case chan int:
		fmt.Println("Переменная имеет тип chan int")
	default:
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
