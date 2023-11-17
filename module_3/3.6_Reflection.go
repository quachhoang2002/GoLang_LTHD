package main

import (
	"fmt"
	"reflect"
)

func main() {
    // Sử dụng reflection để lấy thông tin kiểu dữ liệu của một biến
    var x int = 42
    fmt.Println("Kiểu dữ liệu của x:", reflect.TypeOf(x))

    // Sử dụng reflection để xem giá trị của một biến
    fmt.Println("Giá trị của x:", reflect.ValueOf(x).Int())

    // Tạo một biến và giá trị mới dựa trên kiểu dữ liệu
    type MyStruct struct {
        Name  string
        Value int
    }
    myValue := reflect.New(reflect.TypeOf(MyStruct{})).Elem()
    myValue.FieldByName("Name").SetString("Hello")
    myValue.FieldByName("Value").SetInt(100)
    fmt.Println(myValue.Interface().(MyStruct))
}
