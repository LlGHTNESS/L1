package main

//Реализовать паттерн «адаптер» на любом примере.

import (
	"fmt"
)

type OldPrinter interface {
	PrintOld(s string)
}

type MyOldPrinter struct{}

func (p *MyOldPrinter) PrintOld(s string) {
	fmt.Println("Старый принтер:", s)
}

type NewPrinter interface {
	PrintNew(s string)
}

type PrinterAdapter struct {
	OldPrinter OldPrinter
}

func (a *PrinterAdapter) PrintNew(s string) {
	// Адаптер изменяет вызов нового интерфейса на вызов старого интерфейса.
	a.OldPrinter.PrintOld(s)
}

func main() {
	// Создаем экземпляр старого принтера.
	oldPrinter := &MyOldPrinter{}

	// Адаптируем старый принтер к новому интерфейсу.
	adapter := PrinterAdapter{
		OldPrinter: oldPrinter,
	}

	// Теперь мы можем использовать старый принтер через новый интерфейс.
	adapter.PrintNew("Привет")
}
