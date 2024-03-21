package main

//Реализовать все возможные способы остановки выполнения горутины.
import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				// Завершение работы из-за истечения таймаута
				return
			default:
				fmt.Println("Работает...")
				// Обычная работа
			}
		}
	}()
	time.Sleep(3 * time.Second)
}
