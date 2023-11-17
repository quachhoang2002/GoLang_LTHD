package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	// Tạo một tệp CPU profile
	cpuProfile, err := os.Create("cpu_profile.pprof")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cpuProfile.Close()

	// Tạo một tệp memory profile
	memoryProfile, err := os.Create("memory_profile.pprof")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer memoryProfile.Close()

	// Bắt đầu ghi CPU profile
	if err := pprof.StartCPUProfile(cpuProfile); err != nil {
		fmt.Println(err)
		return
	}
	defer pprof.StopCPUProfile()

	// Tạo một slice và thêm dữ liệu vào nó
	data := make([]int, 1000000)
	for i := range data {
		data[i] = i
	}

	// Tạo một vòng lặp đơn giản để mô phỏng một thời gian chạy của chương trình
	for i := 0; i < 100000; i++ {
		fmt.Printf("Iteration %d\n", i)
		time.Sleep(10 * time.Millisecond)
	}

	// Ghi memory profile
	pprof.WriteHeapProfile(memoryProfile)
}
