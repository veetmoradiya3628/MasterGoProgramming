package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() // current time
	fmt.Printf("Current time : %s\n", now)

	fmt.Printf("Year: %d, Month: %d, Day : %d\n", now.Year(), now.Month(), now.Day())
	fmt.Printf("Hour: %d, Minute: %d, Second: %d\n", now.Hour(), now.Minute(), now.Second())
	fmt.Printf("Weekday: %s\n", now.Weekday())

	launchDate := time.Date(2026, time.February, 28, 19, 0, 0, 0, time.UTC)
	fmt.Printf("Launch date: %s\n", launchDate)

	diff := launchDate.Sub(now)
	fmt.Printf("Total hours difference: %f\n", diff.Hours())
	fmt.Printf("Total minutes difference: %f\n", diff.Minutes())
	fmt.Printf("Total seconds difference: %f\n", diff.Seconds())
	fmt.Printf("Duration in default format: %s\n", diff.String())

	futureTime := now.Add(time.Hour * 2)
	fmt.Printf("Future time: %s\n", futureTime)

	fmt.Printf("Before 10 seconds : %s\n", time.Now())
	time.Sleep(time.Second * 10)
	fmt.Printf("After 10 seconds : %s\n", time.Now())
}
