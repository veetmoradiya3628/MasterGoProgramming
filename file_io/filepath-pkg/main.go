package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// path = C:\Users\Documents - puts path separator according to OS
	path1 := filepath.Join("C:", "Users", "Documents")
	fmt.Println(path1)

	path2 := filepath.Join("folder1", "folder2", "file.txt")
	fmt.Println(path2)

	fmt.Println(filepath.Base(path2))
	fmt.Println(filepath.Ext(path2))

	dirtyDir := "folder1/../folder2/./file.txt"
	cleanDir := filepath.Clean(dirtyDir)
	fmt.Println("Dirty Path:", dirtyDir)
	fmt.Println("Cleaned Path:", cleanDir)
}
