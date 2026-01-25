package main

import "fmt"

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

type ContactInfo struct {
	Email string
	Phone string
}

func (c ContactInfo) DisplayContact() string {
	if c.Email == "" && c.Phone == "" {
		return "No Contact Info provided"
	}
	return fmt.Sprintf("Email: %s, Phone: %s", c.Email, c.Phone)
}

type Company struct {
	Name string
	Address
	ContactInfo
	BusinessType string
}

func (c Company) GetProfile() string {
	return fmt.Sprintf("Company Name: %s\nAddress: %s\nContact Info: %s\nBusiness Type: %s",
		c.Name, c.FullAddress(), c.DisplayContact(), c.BusinessType)
}

func main() {
	company := Company{
		Name: "Tech Innovations",
		Address: Address{
			Street:  "123 Tech Lane",
			City:    "Silicon Valley",
			State:   "CA",
			ZipCode: "94043",
		},
		ContactInfo: ContactInfo{
			Email: "info@techinnovations.com",
			Phone: "123-456-7890",
		},
		BusinessType: "Software Development",
	}

	fmt.Println(company.GetProfile())
}
