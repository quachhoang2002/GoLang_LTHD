package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
    // Đường dẫn đến tệp văn bản cần ghi
    filePath := "D:/temp/count.txt"

    // Đọc nội dung của tệp
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        fmt.Println("Lỗi khi đọc tệp:", err)
        return
    }

    // In nội dung tệp
    fmt.Println("Nội dung của tệp:")
    fmt.Println(string(data))
}
