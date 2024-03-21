// Поменять местами два числа без создания временной переменной.
package main

import (
	"fmt"
)

func main() {
	// + и -
	a := 1
	b := 2

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("%d %d\n", a, b)

	// Через умножение (с 0 не работает)
	a = 1
	b = 2

	a = a * b
	b = a / b
	a = a / b

	fmt.Printf("%d %d\n", a, b)

	// Xor
	a = 1
	b = 2

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("%d %d\n", a, b)
}
