package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "Welcome to the wonderful world of Go!"
	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded:", encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	fmt.Println("Decoded:", string(decoded))

	rawData := []byte{0xDE, 0xAD, 0xBE, 0xEF}
	encodedRaw := base64.StdEncoding.EncodeToString(rawData)
	fmt.Println("Encoded raw bytes:", encodedRaw)
	decodedRaw, err := base64.StdEncoding.DecodeString(encodedRaw)
	if err != nil {
		fmt.Println("Error decoding raw bytes:", err)
		return
	}
	fmt.Printf("Decoded raw bytes: % X\n", decodedRaw)

	if data == string(decoded) {
		fmt.Println("Success: Decoded data matches original data.")
	} else {
		fmt.Println("Failure: Decoded data does not match original data.")
	}
}
