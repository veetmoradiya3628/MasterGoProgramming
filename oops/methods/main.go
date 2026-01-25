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

// A value receiver
func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}

// A pointer receiver
func (e *Employee) DeActivate() {
	e.IsActive = false
}

func (e *Employee) SetJoinedAt(t time.Time) {
	e.JoinedAt = t
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
	fmt.Println(jane.FullName())

	fmt.Printf("%+v\n", jane)
	jane.DeActivate()
	fmt.Printf("%+v\n", jane)

	jane.SetJoinedAt(time.Now().Add(10 * time.Second))
	fmt.Printf("%+v\n", jane)

	// jade := &Employee{}
	// jade = nil
	// jade.FullName() - error since its nil // bad idea
}
