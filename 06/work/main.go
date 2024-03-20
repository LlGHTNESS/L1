package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("Работа завершена...")
		for i := 0; i < 5; i++ {
			fmt.Println("Работает...")
		}
	}()
	time.Sleep(1 * time.Second)
}
