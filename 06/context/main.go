package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				// Очистить ресурсы, завершить работу
				return
			default:
				fmt.Println("Работает...")
				// Обычная работа
			}
		}
	}()
	time.Sleep(1 * time.Second)
	cancel() // Отправить сигнал для выхода
}
