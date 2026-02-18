package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println("Random number between 1 to 99", rand.IntN(100))

	fmt.Println("Random IntN(100) [1, 100):")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d\n", rand.IntN(100))
	}
	// Seed is super important (This is currently using the default seed)

	numbers := []int{10, 20, 30, 40, 50}
	fmt.Printf("\nOriginal slice: %v\n", numbers)
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	fmt.Printf("\nShuffled slice: %v\n", numbers)

	perm := rand.Perm(5)
	fmt.Printf("\nRandom permutation of 5 elements %v\n", perm)
}
