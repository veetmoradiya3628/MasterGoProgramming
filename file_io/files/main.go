package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "text.txt"

	data := "Welcome to Go Programming language!, lots of love for go"
	err := os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("file created")

	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

	// file2, err := os.Create("file-via-create.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// // defer func() {
	// // 	if err := file2.Close(); err != nil {
	// // 		panic(err)
	// // 	}
	// // }()
	// defer file2.Close()
	// _, err = file2.WriteString("Welcome all Java developers to Go lang")
	// if err != nil {
	// 	panic(err)
	// }

	// Reading the file line by line using bufio package and Scanner type from os package
	fileName := "file-via-create.txt"
	printContent(fileName)

	newFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	_, err = newFile.WriteString("\nGo is awesome. Let's learn Go")
	if err != nil {
		panic(err)
	}
	printContent(fileName)
}

func printContent(filePath string) {
	newFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	scanner := bufio.NewScanner(newFile)
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("Line %d: %s\n", lineNum, scanner.Text())
		lineNum++
	}
}
