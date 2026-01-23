package main

import "fmt"

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
}
