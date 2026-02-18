package main

import (
	"fmt"
	"time"
)

func main() {
	// %Y-%m-%d - 2006-01-02:05-04
	now := time.Now()
	fmt.Printf("Current time (default Go format): %s\n", now)

	fmt.Printf("Formatted as YYYY-MM-DD : %s\n", now.Format("2006-01-02"))

	fmt.Printf("Formatted as MM/DD/YYYY hh:mm:ss PM: %s\n", now.Format("01/02/2006 03:04:05 PM"))

	fmt.Printf("Formatted RFC3339 : %s\n", time.RFC3339)

	fmt.Printf("Formatted DateTime : %s\n", time.DateTime)

	dateStr1 := "2025-07-15"
	layout1 := "2006-01-02"
	parsedTime1, err := time.Parse(layout1, dateStr1)
	if err != nil {
		fmt.Printf("Error parsing %s with layout %s: %v\n", dateStr1, layout1, err)
	} else {
		fmt.Printf("Parsed %s", parsedTime1)
	}
}
