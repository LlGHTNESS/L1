package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	m := make(map[int]int)

	// Функция для записи данных в map
	writeData := func(key int, value int) {
		mu.Lock() // Получаем блокировку перед записью данных
		m[key] = value
		mu.Unlock() // Освобождаем блокировку после записи данных
	}

	// Запускаем несколько горутин для конкурентной записи в map
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			writeData(i, i*i)
		}(i)
	}

	// Ждём завершения всех горутин
	wg.Wait()

	// Печатаем результаты
	for k, v := range m {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}
}
