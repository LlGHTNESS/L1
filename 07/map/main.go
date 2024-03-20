package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	m := sync.Map{}

	// Запускаем несколько горутин для конкурентной записи в map
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(i, i*i) // Безопасно записываем данные
		}(i)
	}

	// Ждём завершения всех горутин
	wg.Wait()

	// Печатаем результаты
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true // Продолжаем итерацию
	})
}
