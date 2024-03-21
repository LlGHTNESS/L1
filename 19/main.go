package main

//Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

import (
	"fmt"
)

// reverseString переворачивает строку и возвращает результат.
func reverseString(s string) string {
	// Преобразуем строку в слайс рун для поддержки Unicode символов.
	runes := []rune(s)

	// Переворачиваем слайс рун.
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Возвращаем перевернутую строку.
	return string(runes)
}

func main() {
	input := "главрыба"
	// Вызов функции для переворачивания строки и вывод результата.
	reversed := reverseString(input)
	fmt.Println(reversed)
}
