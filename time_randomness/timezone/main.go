package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("Current local time: %s\n", now)

	newYork, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Printf("Error loading New York timezone: %v\n", err)
	}
	fmt.Printf("New York time now : %v\n", now.In(newYork))
	fmt.Printf("New York time now : %v\n", now.In(newYork))

	tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Printf("Error loading Tokyo timezone: %v\n", err)
	}
	fmt.Printf("Tokyo time now : %v\n", now.In(tokyo))
	fmt.Printf("Tokyo time now : %v\n", now.In(tokyo))

}
