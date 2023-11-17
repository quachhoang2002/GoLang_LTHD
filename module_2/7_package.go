package main

import (
	"fmt"
	"math/rand"
	"module_2/mypackage"
	"time"
)

func Pakacage() {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100)
	fmt.Println("Random number:", num)
	testNum := mypackage.Canculate(num, 10)
	fmt.Println("Test number:", testNum)
}
