package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Giảm đếm của WaitGroup khi goroutine hoàn thành
	fmt.Printf("Worker %d is starting\n", id)
	time.Sleep(1 * time.Second) // Giả lập công việc
	fmt.Printf("Worker %d has finished\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Khởi tạo và khởi động hai goroutines
	wg.Add(2) // Đăng ký hai goroutines với WaitGroup

	go worker(1, &wg)
	go worker(2, &wg)

	// Đợi cho tất cả các goroutines hoàn thành
	wg.Wait()

	fmt.Println("All goroutines have finished.")
}
