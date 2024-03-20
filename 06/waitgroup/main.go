package main

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
