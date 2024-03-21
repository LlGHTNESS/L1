package main

import (
	"fmt"
)

func main() {
	// Инициализация последовательности строк.
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создание множества из последовательности.
	set := make(map[string]struct{})

	for _, value := range sequence {
		// Если строка ещё не содержится в множестве, она добавляется.
		set[value] = struct{}{}
	}

	// Вывод множества на экран.
	fmt.Println("Множество уникальных строк:")
	for key := range set {
		fmt.Println(key)
	}
}
