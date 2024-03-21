package main

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.

import (
	"fmt"
	"sync"
)

// Функция для вычисления квадрата числа и добавления результата в канал
func sqr(wg *sync.WaitGroup, ch chan int, num int) {
	defer wg.Done() // Уменьшаем счётчик группы ожидания
	ch <- num * num // Отправляем квадрат числа в канал
}

func main() {
	// Инициализируем последовательность чисел и создаем канал и группу ожидания
	numbers := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(numbers))
	var wg sync.WaitGroup

	// Передаем группу ожидания и канал в функцию
	for _, number := range numbers {
		wg.Add(1)
		go sqr(&wg, ch, number)
	}

	// Считаем сумму квадратов
	sum := 0
	for square := range ch {
		sum += square
	}

	// Выводим результат
	fmt.Print(sum)

	// Закрываем канал, когда все горутины завершат работу
	go func() {
		wg.Wait()
		close(ch)
	}()
}
