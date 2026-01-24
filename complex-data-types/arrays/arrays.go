package main

import "fmt"

func main() {
	var numbers [2]int
	fmt.Printf("%+v\n", numbers)

	numbers[0] = 1
	numbers[1] = 2
	fmt.Printf("%+v\n", numbers)

	primes := [4]int{2, 3, 5, 7}
	fmt.Printf("%+v\n", primes)
	primes[3] = 11
	fmt.Printf("%+v\n", primes)

	for i := 0; i < len(primes); i++ {
		fmt.Printf("%d ", primes[i])
	}
	fmt.Println()

	var matrix [2][3]int
	matrix[0][1] = 2
	fmt.Printf("%+v\n", matrix)
}
