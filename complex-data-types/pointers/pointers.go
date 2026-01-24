package main

import "fmt"

func modifyValue(val int) {
	val = val * 10
	fmt.Printf("modifyValue: %+v\n", val)
}

func modifyPointer(val *int) {
	if val == nil {
		return
	}
	*val = *val * 10 // de referencing
	fmt.Printf("modifyPointer: %+v\n", *val)
}

func main() {
	age := 10 //
	agePtr := &age
	fmt.Printf("age address: %d\n", &age)
	fmt.Printf("agePtr: %d\n", agePtr)

	num := 10
	modifyValue(num)
	fmt.Println(num)

	modifyPointer(&num)
	fmt.Println(num)

	grade := 50
	gradePtr := &grade
	fmt.Printf("gradePtr grade: %+v\n", gradePtr)
	fmt.Printf("gradePtr : %+v\n", *(&gradePtr)) // for understanding
}
