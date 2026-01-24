package main

import (
	"fmt"
	"strings"
)

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextID int = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addContact(name, email, phone string) {
	if _, exists := contactIndexByName[name]; exists {
		fmt.Printf("Contact already exists : %+v\n", name)
		return
	}
	newContact := Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	nextID++
	contactList = append(contactList, newContact)
	contactIndexByName[name] = len(contactList) - 1
	fmt.Printf("Contact added: %v\n", name)
}

func findContactByName(name string) *Contact {
	index, exists := contactIndexByName[name]
	if exists {
		return &contactList[index]
	}
	return nil
}

func listContacts() {
	if len(contactList) == 0 {
		fmt.Println("No contacts found.")
		return
	}
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("CONTACT LIST")
	fmt.Println(strings.Repeat("=", 50))
	for _, contact := range contactList {
		fmt.Printf("ID: %d\n", contact.ID)
		fmt.Printf("Name: %s\n", contact.Name)
		fmt.Printf("Email: %s\n", contact.Email)
		fmt.Printf("Phone: %s\n", contact.Phone)
		fmt.Println(strings.Repeat("-", 50))
	}
}

func main() {
	listContacts()
	for {
		fmt.Print("\n=== Contact Management ===\n1. Add Contact\n2. Find Contact\n3. List Contacts\n4. Exit\nChoose an option: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var name, email, phone string
			fmt.Print("Enter name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter email: ")
			fmt.Scanln(&email)
			fmt.Print("Enter phone: ")
			fmt.Scanln(&phone)
			addContact(name, email, phone)
		case 2:
			var name string
			fmt.Print("Enter name to find: ")
			fmt.Scanln(&name)
			contact := findContactByName(name)
			if contact != nil {
				fmt.Printf("Found: %+v\n", contact)
			} else {
				fmt.Println("Contact not found.")
			}
		case 3:
			listContacts()
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}
