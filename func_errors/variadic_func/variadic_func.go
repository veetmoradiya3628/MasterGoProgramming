package main

import (
	"errors"
	"fmt"
	"strings"
)

func sum(a, b int) int {
	return a + b
}

func sumOfNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

func splitName(fullName string) (firstName, lastName string) {
	parts := strings.Split(fullName, " ")
	return parts[0], parts[1]
}

func main() {
	fmt.Println(sum(2, 3))
	fmt.Println(sumOfNumbers(1, 2, 3, 4))

	value, err := divide(10, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	firstName, lastName := splitName("Veet Moradiya")
	fmt.Println(firstName, lastName)
}
