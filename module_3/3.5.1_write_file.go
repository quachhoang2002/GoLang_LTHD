package main

import (
	"fmt"
	"os"
)

func main() {
    // Đường dẫn đến tệp văn bản cần ghi
    filePath := "D:/temp/count.txt"

    file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Lỗi khi mở tệp:", err)
        return
    }
    defer file.Close() // Đảm bảo đóng tệp sau khi hoàn thành

    // Nội dung cần ghi thêm
    content := []byte("Nội dung mới được ghi thêm vào tệp.\n")

    // Ghi nội dung vào tệp
    _, err = file.Write(content)
    if err != nil {
        fmt.Println("Lỗi khi ghi thêm vào tệp:", err)
        return
    }

    fmt.Println("Đã ghi thêm nội dung vào tệp:", filePath)
}
