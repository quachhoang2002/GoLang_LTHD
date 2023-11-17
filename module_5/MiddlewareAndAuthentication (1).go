package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var jwtKey = []byte("your-secret-key")

//main
func middleware() {
	r := mux.NewRouter()

	r.HandleFunc("/public", PublicHandler).Methods("GET")
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/protected", AuthMiddleware(ProtectedHandler)).Methods("GET")

	http.Handle("/", r)

	fmt.Println("Server is running on :8080...")
	http.ListenAndServe(":8080", nil)
}

func PublicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a public route.")
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a protected route.")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Xác thực người dùng và tạo JWT token
	tokenString, err := CreateToken()
	if err != nil {
		http.Error(w, "Unable to create token", http.StatusUnauthorized)
		return
	}
	w.Write([]byte("Login successful! Token: " + tokenString))
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := ExtractToken(r)
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if err := VerifyToken(tokenString); err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func ExtractToken(r *http.Request) string {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return ""
	}
	return tokenString
}

func CreateToken() (string, error) {
	// Tạo claims
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
		// Các thông tin khác có thể thêm vào claims
	}

	// Tạo token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Ký và chuyển token thành chuỗi
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return fmt.Errorf("Token is not valid")
	}

	return nil
}
