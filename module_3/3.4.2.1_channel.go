package main

import (
	"fmt"
	"time"
)

func main() {
	// Tạo một channel kiểu int
	ch := make(chan int)

	// Goroutine 1: Gửi dữ liệu vào channel
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i // Gửi giá trị i vào channel
			time.Sleep(100 * time.Millisecond)
		}
		close(ch) // Đóng channel sau khi gửi xong
	}()

	// Goroutine 2: Nhận dữ liệu từ channel và in ra màn hình
	go func() {
		for num := range ch {
			fmt.Println("Received:", num)
		}
	}()

	// Đợi một chút để cho các goroutine chạy xong
	time.Sleep(1 * time.Second)
}

func main2() {
	// Tạo một channel kiểu int
	ch := make(chan int)

	// Goroutine 1: Gửi dữ liệu vào channel
	go func() {
		// for i := 1; i <= 5; i++ {
		// 	ch <- i // Gửi giá trị i vào channel
		time.Sleep(1000 * time.Millisecond)
		// }

		ch <- 100
	}()

	go func() {
		// for i := 1; i <= 5; i++ {
		// 	ch <- i // Gửi giá trị i vào channel
		time.Sleep(2000 * time.Millisecond)
		// }

		ch <- 10000
	}()

	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)

	// Đợi một chút để cho các goroutine chạy xong
	time.Sleep(1 * time.Second)
}
