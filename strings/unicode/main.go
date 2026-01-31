package main

import (
	"fmt"
	"unicode"
)

func main() {
	username := "test" // 4 bytes - 4 * 1
	fmt.Println(len(username))

	gujarati_name := "વિત" // 9 bytes - 3 * 3
	fmt.Println(len(gujarati_name))
	fmt.Printf("%c\n", gujarati_name[0])         // à
	fmt.Printf("%c\n", []rune(gujarati_name)[0]) // વ

	for _, v := range gujarati_name {
		fmt.Printf("%s\n", string(v))
	}

	data := []rune{'વ', 'ત', '1'}
	for _, v := range data {
		fmt.Println(string(v), unicode.IsDigit(v), unicode.IsLetter(v))
	}
}
