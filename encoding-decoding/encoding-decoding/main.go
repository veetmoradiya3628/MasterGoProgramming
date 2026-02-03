package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"full_name" xml:"full_name"`
	Age      int    `json:"age" xml:"age"`
	Phone    string `json:"phone" xml:"tel"`
	Password string `json:"-" xml:"-"`
	IsActive bool   `json:"is_active" xml:"is_active"`
}

var payload = `{"full_name":"John Doe 2","age":12,"phone":"123-456-7890","is_active":true}`

func main() {
	u := User{
		Name:     "John Doe",
		Age:      30,
		Phone:    "123-456-7890",
		Password: "secret",
		IsActive: true,
	}
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(u); err != nil {
		panic(err)
	}
	fmt.Println(buf.String())

	var u2 User
	dec := json.NewDecoder(bytes.NewBufferString(payload))
	if err := dec.Decode(&u2); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", u2)
}
