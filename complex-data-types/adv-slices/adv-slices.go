package main

import (
	"fmt"
	"slices"
)

func main() {
	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Original: %+v, Len: %d, Cap: %d\n", original, len(original), cap(original))

	s1 := original[2:5]
	fmt.Printf("s1: %+v, Len: %d, Cap: %d\n", s1, len(s1), cap(s1))

	s2 := original[:4]
	fmt.Printf("s2: %+v, Len: %d, Cap: %d\n", s2, len(s2), cap(s2))

	s3 := original[6:]
	fmt.Printf("s3: %+v, Len: %d, Cap: %d\n", s3, len(s3), cap(s3))

	fmt.Println(slices.Contains(s3, 10))
	s3 = append(s3, 10)
	fmt.Println(slices.Contains(s3, 10))

}
