package main

import "fmt"

/*

Other language :-

class User {}

class Admin extends User {}

Go

- Composition - HAS-A relationship
- Inheritance - IS-A relationship
Car -> is composed of many parts (Engine, Doors)

*/

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

func (a Address) FullAddress() string {
	if a.Street == "" && a.City == "" {
		return "No Address provider"
	}
	return fmt.Sprintf("%s, %s, %s, %s", a.Street, a.City, a.State, a.ZipCode)
}

type Customer struct {
	CustomerID      int
	Name            string
	Email           string
	BillingAddress  Address // embadded
	ShippingAddress Address
}

func (c Customer) PrintDetails() {
	fmt.Println("Customer ID:", c.CustomerID)
	fmt.Println("Name:", c.Name)
	fmt.Println("Email:", c.Email)
	fmt.Println("Billing Address:", c.BillingAddress.FullAddress())
	fmt.Println("Shipping Address:", c.ShippingAddress.FullAddress())
}

func main() {
	fmt.Println("------ Composition ---------")

	customer1 := Customer{
		CustomerID: 1,
		Name:       "John Doe",
		Email:      "john@example.com",
		BillingAddress: Address{
			Street:  "123 Main St",
			City:    "New York",
			State:   "NY",
			ZipCode: "10001",
		},
		ShippingAddress: Address{
			Street:  "456 Oak Ave",
			City:    "Boston",
			State:   "MA",
			ZipCode: "02101",
		},
	}
	customer1.PrintDetails()

	fmt.Println("------- Customer with same billing and shipping address ----------")
	mainAddress := Address{
		Street:  "789 Elm St",
		City:    "Chicago",
		State:   "IL",
		ZipCode: "60601",
	}

	customer2 := Customer{
		CustomerID:      2,
		Name:            "Jane Smith",
		Email:           "jane@example.com",
		BillingAddress:  mainAddress,
		ShippingAddress: mainAddress,
	}
	customer2.PrintDetails()
}
