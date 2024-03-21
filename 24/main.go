package main

/*Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
 */
import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) Distance(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаем две точки с помощью конструктора.
	p1 := NewPoint(1.0, 1.0)
	p2 := NewPoint(4.0, 5.0)

	// Вычисляем расстояние между точками.
	distance := p1.Distance(p2)

	fmt.Println(distance)
}
