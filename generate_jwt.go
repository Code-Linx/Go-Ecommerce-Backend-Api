package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	secret := []byte("not-secret-secret-anymore") // ✅ Your JWT secret

	claims := jwt.MapClaims{
		"userID": "1",                                   // ✅ Your custom payload
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // ✅ Expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}

	fmt.Println("Your JWT token is:", signedToken)
}
