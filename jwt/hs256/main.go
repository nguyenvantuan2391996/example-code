package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	JWTSecret            = []byte("123456")
	JWTAlgorithm         = jwt.SigningMethodHS256
	JWTExpirationSeconds = 10 * 60
)

func GenerateJWT(userID string) (string, error) {
	// Define token claims
	claims := jwt.MapClaims{
		"exp":        time.Now().Add(time.Second * time.Duration(JWTExpirationSeconds)).Unix(),
		"token_type": "access",
		"user_id":    17,
		"username":   "test",
		"role":       "normal",
		"jti":        uuid.NewString(),
	}

	// Create token
	token := jwt.NewWithClaims(JWTAlgorithm, claims)

	// Sign token with secret
	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func main() {
	token, err := GenerateJWT("")
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}
	fmt.Println(token)
}
