package main

import (
	"fmt"
)

func main() {
	array := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)
	new := grouping(groups, array)

	for key, value := range new {

		fmt.Printf("%d: [", key)
		for _, t := range value {
			fmt.Printf("%f, ", t)
		}
		fmt.Printf("]\n")
	}
}
func grouping(groups map[int][]float64, array []float64) map[int][]float64 {
	for _, temperature := range array {
		group := (int(temperature) - (int(temperature) % 10))
		groups[group] = append(groups[group], temperature)
	}
	return groups
}
