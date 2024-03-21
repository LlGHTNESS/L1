package main

/*Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».*/
import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	// Разделение строки на слова.
	words := strings.Fields(s)

	// Переворачиваем порядок слов.
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Объединяем слова обратно в строку и возвращаем результат.
	return strings.Join(words, " ")
}

func main() {
	// Исходная строка для обработки.
	input := "snow dog sun"

	// Переворачиваем слова в строке.
	reversed := reverseWords(input)
	fmt.Println(reversed)
}
