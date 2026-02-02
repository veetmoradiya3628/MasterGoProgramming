package main

import (
	"fmt"
	"os"
)

func main() {
	tempFile, err := os.CreateTemp("", "logs.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		fmt.Println("removing tempFile", tempFile.Name())
		// _ = os.Remove(tempFile.Name()) - if not removing, it will remain in the temp folder and it will be cleaned up by the OS later
	}()

	_, err = tempFile.Write([]byte("Hello world\n"))
	if err != nil {
		panic(err)
	}
	fmt.Println("tempFile created:", tempFile.Name())

	tempDir, err := os.MkdirTemp("", "my_app_logs")
	if err != nil {
		panic(err)
	}
	defer func() {
		fmt.Println("removing tempDir", tempDir)
		_ = os.RemoveAll(tempDir)
	}()
	fmt.Println("tempDir created:", tempDir)
}
