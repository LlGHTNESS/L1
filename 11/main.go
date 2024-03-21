package main

import (
	"fmt"
)

// intersect принимает два "множества" в виде карт и возвращает их пересечение.
func intersect(set1, set2 map[int]struct{}) map[int]struct{} {
	// Инициализация результата пересечения.
	intersection := make(map[int]struct{})

	// Проходим по всем элементам первого множества.
	for element := range set1 {
		// Проверяем, существует ли элемент во втором множестве.
		if _, found := set2[element]; found {
			// Если элемент найден и в втором множестве, добавляем его в результат.
			intersection[element] = struct{}{}
		}
	}

	// Возвращаем множество, представляющее пересечение.
	return intersection
}

func main() {
	// Инициализируем два множества в виде карт.
	set1 := map[int]struct{}{1: {}, 2: {}, 3: {}}
	set2 := map[int]struct{}{2: {}, 3: {}, 4: {}}

	// Вызываем функцию intersect для нахождения пересечения двух множеств.
	result := intersect(set1, set2)

	// Выводим результат.
	fmt.Println("Пересечение множеств: ")
	for key := range result {
		fmt.Println(key)
	}
}
