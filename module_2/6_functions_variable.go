package main

import "fmt"

// Function
func add(a, b int) int {
	return a + b
}

// Method (attached to a struct)
type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func Function() {
	// Using the function
	result := add(3, 4)
	fmt.Println("Result:", result)

	// Using the method
	rect := Rectangle{Width: 5, Height: 3}
	area := rect.Area()
	fmt.Println("Area:", area)
}
