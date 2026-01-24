package main

import "fmt"

// dynamic array - slices
func main() {
	names := []string{"Alice", "John", "Mark"}
	fmt.Println(names)

	items := make([]int, 3, 5)
	fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", items, len(items), cap(items))

	items = append(items, 100)
	items = append(items, 101)
	fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", items, len(items), cap(items))

	items = append(items, 102)
	fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", items, len(items), cap(items))

	fmt.Printf("%+v\n", items[3:6])
}
