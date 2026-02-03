package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string  `json:"full_name" xml:"full_name"`
	Age      int     `json:"age" xml:"age"`
	Phone    string  `json:"phone" xml:"tel"`
	Password string  `json:"-" xml:"-"`
	IsActive bool    `json:"is_active" xml:"is_active"`
	Role     string  `json:"role" xml:"role"`
	Profile  Profile `json:"profile" xml:"profile"`
}

type Profile struct {
	URL string `json:"url" xml:"url"`
}

var payload = `{
	"full_name": "Jane",
	"age": 30,
	"phone": "123-456-789",
	"is_active": true,
	"profile": {
		"url": "https://example.com/jane"
	}
}
`

func main() {
	// marshalling
	jane := User{
		Name:     "Jane",
		Age:      30,
		IsActive: true,
		Phone:    "123-456-789",
		Profile:  Profile{URL: "https://example.com/jane"},
	}
	byteSlice, err := json.MarshalIndent(jane, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byteSlice))

	// unmarshalling
	var user User
	err = json.Unmarshal([]byte(payload), &user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", user)
}
