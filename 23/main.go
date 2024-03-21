package main

//Удалить i-ый элемент из слайса.
import (
	"fmt"
)

func remove(slice []int, i int) []int {
	// Соединяем два среза: один до i-го элемента, второй после i-го элемента.
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	// Инициализируем слайс чисел.
	mySlice := []int{1, 2, 3, 4, 5}

	i := 2

	// Вызываем функцию для удаления элемента.
	newSlice := remove(mySlice, i)

	// Выводим новый слайс.
	fmt.Println(newSlice)
}
