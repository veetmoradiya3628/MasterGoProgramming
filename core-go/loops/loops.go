package main

import "fmt"

func main() {
	// while, do-while, foreach, for - other programming language
	// for - only this in go

	for i := 1; i <= 10; i++ {
		fmt.Println("C style loop ", i)
	}

	k := 3
	for k > 0 {
		fmt.Println("While style ", k)
		k--
	}

	count := 0
	for {
		fmt.Println("Infinite loop")
		count++
		if count >= 5 {
			break
		}
	}

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	items := [3]string{"Go", "Python", "Java"}
	for index, value := range items {
		fmt.Println(index, value)
	}
}
