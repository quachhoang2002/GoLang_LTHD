package main

import "fmt"

func Loop() {
	// For loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// While loop (Go doesn't have a while keyword)
	j := 1
	for j <= 5 {
		fmt.Println(j)
		j++
	}
}
