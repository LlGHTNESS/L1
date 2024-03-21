package main

//Реализовать все возможные способы остановки выполнения горутины.
import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Работает...")
	}()
	wg.Wait()
	fmt.Println("Не работает.")
}
