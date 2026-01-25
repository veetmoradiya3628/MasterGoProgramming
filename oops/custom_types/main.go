package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActive  bool
	JoinedAt  time.Time
}

func NewEmployee(id int, firstName, lastName, position string, isActive bool) Employee {
	return Employee{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Position:  position,
		IsActive:  isActive,
		Salary:    10,
		JoinedAt:  time.Now(),
	}
}

func main() {
	jane := Employee{
		ID:        1,
		FirstName: "Jane",
		LastName:  "Doe",
		Position:  "Night",
		Salary:    1000,
		IsActive:  true,
		JoinedAt:  time.Now(),
	}

	fmt.Printf("%+v\n", jane)
	fmt.Println(jane.ID)
	fmt.Println(jane.FirstName)

	joe := NewEmployee(1, "John", "Doe", "Jane", true)
	fmt.Printf("%+v\n", joe)
	joe.Salary = 10000
	fmt.Printf("%+v\n", joe)

	joePtr := &joe
	fmt.Println(joePtr)
	fmt.Println(*joePtr)

	jane.LastName = "Adam"
	fmt.Printf("%+v\n", jane)
}
