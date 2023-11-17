package main

import "fmt"

func Operators() {
	// Arithmetic operators
	a, b := 10, 5
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)

	// Comparison operators
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)

	// Logical operators
	x, y := true, false
	fmt.Println(x && y)
	fmt.Println(x || y)
	fmt.Println(!x)
}
