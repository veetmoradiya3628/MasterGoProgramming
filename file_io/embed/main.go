package main

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed hello.txt
var data string

//go:embed public
var publicFiles embed.FS

func main() {
	fmt.Println(data)
	content, _ := publicFiles.ReadFile("public/data.txt")
	fmt.Println(string(content))
}
