package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("Работа завершена...")
				return
			default:
				fmt.Println("Работает...")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	ch <- 1
}
