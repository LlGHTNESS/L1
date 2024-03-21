package main

//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

import (
	"fmt"
)

// Функция вычисляет квадрат числа и отправляет его в канал
func sqr(num int, ch chan int) {
	ch <- num * num
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	ch := make(chan int) // Создаем канал для передачи результатов

	for _, number := range numbers {
		go sqr(number, ch) // Запускаем горутину
	}

	for range numbers {
		fmt.Println(<-ch) // Читаем результат из канала
	}
	close(ch) // Закрываем канал
}
