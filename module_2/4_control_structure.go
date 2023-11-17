package main

import "fmt"

func Controlstructure() {
	// If-else
	num := 10
	if num > 5 {
		fmt.Println("Greater than 5")
	} else {
		fmt.Println("Less than or equal to 5")
	}

	// Switch
	day := "Sunday"
	switch day {
	case "Monday", "Tuesday":
		fmt.Println("It's a weekday")
	case "Saturday", "Sunday":
		fmt.Println("It's a weekend")
	default:
		fmt.Println("Not a valid day")
	}
}
