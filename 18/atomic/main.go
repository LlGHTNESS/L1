package main

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int64 // Текущее значение счетчика. Используется int64 для атомарных операций.
}

func (c *Counter) Increment() {
	atomic.AddInt64(&c.value, 1) // Атомарно увеличиваем значение счетчика на единицу.
}

func (c *Counter) Value() int64 {
	return atomic.LoadInt64(&c.value) // Атомарно получаем текущее значение счетчика.
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{} // Создаем экземпляр счетчика.

	// Запускаем 1000 горутин, каждая из которых будет инкрементировать счетчик.
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()                    // Ожидаем завершения всех горутин.
	fmt.Println(counter.Value()) // Выводим итоговое значение счетчика.
}
