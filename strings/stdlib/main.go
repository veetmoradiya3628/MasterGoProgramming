package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "  abc "
	s2 := strings.Clone(s1)

	fmt.Println(s1, s2)

	b := strings.Builder{}
	b.Write([]byte("Here is the example"))
	fmt.Println(b.String())

	b.WriteString(" Ok this is write string")
	fmt.Println(b.String())

	fmt.Println(strings.ToUpper(s1))
	fmt.Println(strings.TrimSpace(s1))

	fmt.Println(strings.HasSuffix("veet@google.com", "google.com"))

	fmt.Println(strings.ReplaceAll("hey hi how are you hi", "hi", "hello"))

	parts := strings.Split("veet@gmail.com", "@")
	username, domain := parts[0], parts[1]
	fmt.Println(username, domain)

	parts2 := strings.Fields("test example.com")
	username, domain = parts2[0], parts2[1]
	fmt.Println(username, domain)

	fmt.Println(strings.Join(parts, ":"))
}
