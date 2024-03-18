package main

import (
	"fmt"
	"sync"
)

func sqr(num int, wg *sync.WaitGroup) {
	result := num * num
	fmt.Println(result) // Выводим результат
	defer wg.Done()     // Помечаем выполненным при завершении
}

func main() {
	array := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup // Группа ожидания

	for _, num := range array {
		wg.Add(1)        // Уведомление о новой горутине
		go sqr(num, &wg) // Запуск горутины
	}

	wg.Wait() // Ожидание завершения всех горутин
}
