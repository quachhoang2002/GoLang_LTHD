package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
    // Đường dẫn đến tệp văn bản cần ghi
    filePath := "D:/temp/count.txt"

     // Nội dung cần ghi vào tệp
     content := []byte("Nội dung mẫu để ghi vào tệp văn bản")

     // Ghi nội dung vào tệp
     err := ioutil.WriteFile(filePath, content, 0644)
     if err != nil {
         fmt.Println("Lỗi khi ghi tệp:", err)
         return
     }
 
     fmt.Println("Đã ghi thành công vào tệp:", filePath)
}
