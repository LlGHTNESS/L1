package main

//Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

import (
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Println("Введите время работы (секунд):")
	fmt.Scan(&n)

	ch := make(chan int)
	timer := time.NewTimer(time.Duration(n) * time.Second)
	// Запускаем горутину producer, которая будет отправлять данные в канал
	go produce(ch, timer.C)

	// Печатаем числа, полученные из канала, до его закрытия
	for val := range ch {
		fmt.Println(val)
	}

	fmt.Println("Время истекло, программа завершила работу.")
}

func produce(ch chan int, timerC <-chan time.Time) {
	for i := 0; ; i++ {
		select {
		case <-timerC:
			close(ch) // Закрыть канал для завершения работы
			return
		case ch <- i:
			// Отправить значение в канал
			time.Sleep(500 * time.Millisecond) // Пауза перед следующей отправкой
		}
	}
}
