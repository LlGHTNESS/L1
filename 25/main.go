package main

//Реализовать собственную функцию sleep.

import (
	"fmt"
	"time"
)

// Занимает поток, но работает
func ssleep(duration time.Duration) {
	start := time.Now()
	for time.Since(start) < duration {

	}
}

// Работает не занимая поток
func sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println("Начало сна...")
	ssleep(2 * time.Second) // Задержка на 2 секунды
	fmt.Println("Проснулись!")
	fmt.Println("Начало сна...")
	sleep(2 * time.Second) // Задержка на 2 секунды
	fmt.Println("Проснулись!")
}
