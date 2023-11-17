package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// User là cấu trúc dữ liệu biểu diễn người dùng, bao gồm Tên và Email.
type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

// HTTPResponse là cấu trúc cho phản hồi HTTP tuỳ chỉnh.
type HTTPResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

var validate = validator.New() // Tạo đối tượng validator

func main() {
	http.HandleFunc("/signup", signupHandler) // Định tuyến cho đường dẫn đăng ký người dùng
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// signupHandler xử lý các yêu cầu đăng ký người dùng.
func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpResponse(w, http.StatusMethodNotAllowed, "Phương thức không được cho phép", nil)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		httpResponse(w, http.StatusBadRequest, "Dữ liệu yêu cầu không hợp lệ", nil)
		return
	}

	// Kiểm tra dữ liệu đầu vào dựa trên các quy tắc đã định nghĩa trong struct User.
	if err := validate.Struct(user); err != nil {
		httpResponse(w, http.StatusBadRequest, "Dữ liệu không hợp lệ: "+err.Error(), nil)
		return
	}

	// Trả về phản hồi thành công với mã trạng thái 201.
	httpResponse(w, http.StatusCreated, "Người dùng đã được tạo thành công", user)
}

// httpResponse giúp xử lý việc trả về phản hồi HTTP.
func httpResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json") // Thiết lập kiểu nội dung là JSON
	w.WriteHeader(statusCode)                          // Thiết lập mã trạng thái HTTP

	response := HTTPResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}

	// Mã hóa và gửi phản hồi dưới dạng JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Lỗi khi gửi phản hồi: %v\n", err)
	}
}
