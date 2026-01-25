package main

import (
	"fmt"
)

type Person interface {
	GetName() string
}

type Employee struct {
	ID   int
	Name string
}

type BusinessPerson struct {
	ID   int
	Name string
}

// GetName implements [Person].
func (b BusinessPerson) GetName() string {
	return b.Name
}

func (e Employee) GetName() string {
	return e.Name
}

// similar to toString() functionality
func (e BusinessPerson) String() string {
	return fmt.Sprintf("Person[ID:%d, Name: %s]", e.ID, e.Name)
}

type ID int

func (e ID) String() string {
	return fmt.Sprintf("Coming from here ID:%d", e)
}

func main() {

	jane := BusinessPerson{
		ID:   1,
		Name: "Jane",
	}
	fmt.Println(jane)

	var myId ID
	myId = 30
	fmt.Println(myId)
}
