package main

//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20
import (
	"fmt"
	"math/big"
)

func main() {
	// int64 может хранить такие числа (до 2^62)
	var a, b int64
	a = int64(1 << 22)
	b = int64(1 << 22)

	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a + b)
	fmt.Println(a - b)

	// Для больших чем 2^62 чисел придуман модуль big
	c := big.NewInt(1 << 22)
	d := big.NewInt(1 << 22)

	fmt.Println(new(big.Int).Mul(c, d))
	fmt.Println(new(big.Int).Div(c, d))
	fmt.Println(new(big.Int).Add(c, d))
	fmt.Println(new(big.Int).Sub(c, d))

	// Можно использовать строку

	e := new(big.Int)
	e.SetString("100000000000000000000000000000000", 10)
	f := new(big.Int)
	f.SetString("100000000000000000000000000000000", 10)

	fmt.Println(new(big.Int).Mul(e, f))
	fmt.Println(new(big.Int).Div(e, f))
	fmt.Println(new(big.Int).Add(e, f))
	fmt.Println(new(big.Int).Sub(e, f))
}
