package main

import "fmt"

// hashmap, dictionary, associative array
// zero value is nil
func main() {
	studentGrads := map[string]int{
		"Alice": 90,
		"James": 85,
		"Dan":   60,
	}
	fmt.Println(studentGrads)

	studentGrads["Alice"] = 95
	fmt.Println(studentGrads)

	alice, ok := studentGrads["Alice"]
	if ok {
		fmt.Println(alice)
	}

	if _, ok := studentGrads["Bob"]; ok {
		fmt.Println("Bob exists")
	}

	if _, ok := studentGrads["Dan"]; ok {
		fmt.Println("Dan exists")
	}

	delete(studentGrads, "Alice")
	fmt.Println(studentGrads)

	configs := make(map[string]int)
	fmt.Println(configs)
}
