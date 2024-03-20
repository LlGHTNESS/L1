package main

import (
	"fmt"
)

// SetBit устанавливает i-й бит числа n в 1
func SetBit(n int64, i uint) int64 {
	return n | (1 << i)
}

// ClearBit устанавливает i-й бит числа n в 0
func ClearBit(n int64, i uint) int64 {
	return n & ^(1 << i)
}

func main() {
	var number int64 = 0 // Начальное число
	var index uint = 3   // Битовый индекс

	// Устанавливаем i-й бит в 1
	number = SetBit(number, index)
	fmt.Printf("Число в бинарном виде: %b\n", number)

	// Устанавливаем i-й бит обратно в 0
	number = ClearBit(number, index)
	fmt.Printf("Число в бинарном виде: %b\n", number)
}
