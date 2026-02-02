package main

import (
	"log"
	"os"
)

func main() {
	if err := os.Mkdir("Downloads", 0755); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("Downloads/subfolder", 0755); err != nil {
		log.Fatal(err)
	}

	dirtyPath := "Downloads/subfolder"
	if err := os.Remove(dirtyPath); err != nil {
		log.Printf("Remove error: %v", err)
	}

}
