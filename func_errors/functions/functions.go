package main

import "fmt"

func greet(name string) {
	fmt.Printf("%s\n", "Hello, "+name)
}

func add(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}

func calculateArea(width, height float64) float64 {
	if width < 0 || height < 0 {
		fmt.Println("Error: height and width must be > 0")
		return 0
	}
	return width * height
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	greet("Veet")
	add(2, 3)
	fmt.Printf("Area : %.2f\n", calculateArea(2.0, 3.0))
	fmt.Println(factorial(5))
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	logger := func(msg string) {
		fmt.Println(msg)
	}
	logger("Hello world")
	logger("Logger")
}
