package main

import "fmt"

type Number interface {
	int | float64 | float32
}

// Generic function
func Sum[T Number](numbers ...T) T {
	var total T
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	grades := []int{90, 85}
	people := []string{"Jane", "John", "Mark"}

	fmt.Println(len(grades), len(people))
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(1.2, 1.5))
	fmt.Println(Sum(1, 2, 3.2))
}
