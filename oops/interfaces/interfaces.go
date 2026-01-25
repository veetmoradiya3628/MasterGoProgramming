package main

import (
	"fmt"
)

// To support interface implementation all method to be implemented

type Person interface {
	GetName() string
	GetId() int
}

type Employee struct {
	ID   int
	Name string
}

// GetId implements [Person].
func (e Employee) GetId() int {
	return e.ID
}

type BusinessPerson struct {
	ID   int
	Name string
}

// GetId implements [Person].
func (b BusinessPerson) GetId() int {
	return b.ID
}

// GetName implements [Person].
func (b BusinessPerson) GetName() string {
	return b.Name
}

func (e Employee) GetName() string {
	return e.Name
}

func displayPerson(p Person) {
	fmt.Println(p.GetName())
}

func main() {
	joe := Employee{
		ID:   1,
		Name: "Joe",
	}
	// displayPerson(joe)

	jane := BusinessPerson{
		ID:   1,
		Name: "Jane",
	}
	// displayPerson(jane)
	fmt.Println(jane)
	fmt.Println(joe)
}
