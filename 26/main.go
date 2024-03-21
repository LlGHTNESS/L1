package main

/*Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false*/
import (
	"fmt"
	"strings"
)

// IsUniqueInsensitive проверяет, что все символы в строке уникальны, не учитывая регистр.
func IsUniqueInsensitive(s string) bool {
	charMap := make(map[rune]bool)
	for _, char := range strings.ToLower(s) {
		if _, exists := charMap[char]; exists {
			// Если символ уже есть в карте, значит он повторяется.
			return false
		}
		// Добавляем символ в карту.
		charMap[char] = true
	}
	return true
}

func main() {
	testStrings := []string{"abcd", "abCdefAaf", "	aabcd"}
	for _, s := range testStrings {
		fmt.Println(IsUniqueInsensitive(s))
	}
}
