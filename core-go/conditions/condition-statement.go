package main

import (
	"fmt"
	"time"
)

func main() {
	temp := 35
	if temp > 30 {
		fmt.Println("Temprature is greater than 30")
	} else {
		fmt.Println("Temprature less than 30")
	}

	score := 85
	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else if score >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}

	userAccess := map[string]bool{
		"jane": false,
		"john": false,
	}

	hasAccess, ok := userAccess["jane"]

	if ok && hasAccess {
		fmt.Println("Jane can access the system")
	} else {
		fmt.Println("Jane can not access the system")
	}

	day := "Sunday"
	fmt.Println("Today is ", day)

	switch day {
	case "Sunday", "Saturday":
		fmt.Println("Weekend...")
	case "Monday", "Tuesday":
		fmt.Println("Work day...")
	default:
		fmt.Println("Mid week")
	}

	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good Morning")
	case hour < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good evening")
	}

	checkType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("Integer: %d\n", v)
		case string:
			fmt.Printf("String: %s\n", v)
		case bool:
			fmt.Printf("Boolean: %t\n", v)
		default:
			fmt.Printf("Unknown type %T\n", v)
		}
	}
	checkType(21)
	checkType("Test")
	checkType(true)
	checkType(312.12)
}
