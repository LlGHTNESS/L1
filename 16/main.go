package main

import "fmt"

// Функция quickSort реализует алгоритм быстрой сортировки.
func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		// Разделяем массив, и получаем индекс опорного элемента.
		arr, p = partition(arr, low, high)
		// Рекурсивно применяем быструю сортировку для подмассивов до и после опорного элемента.
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

// quickSortStart - функция для запуска алгоритма быстрой сортировки.
func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1) // Запуск сортировки со всеми элементами слайса.
}

// partition - функция для разделения массива.
func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high] // Опорный элемент выбран как последний элемент слайса.
	i := low           // Индекс меньшего элемента и индекс разделения.

	// Перемещаем элементы меньше опорного элемента к началу слайса.
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i] // Меняем местами, если текущий элемент меньше опорного.
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i] // Меняем местами опорный элемент с элементом на индексе разделения.
	return arr, i
}

func main() {
	array := []int{1, 5, 3, 8, 5, 2, 6, 4}
	sortedArray := quickSortStart(array)

	// Выводим отсортированный массив.
	for _, value := range sortedArray {
		fmt.Println(value)
	}
}
