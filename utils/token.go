package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateAccessToken(id, email, role string) (int, string, error) {
	secret := os.Getenv("JWT_PRIVATE_KEY")
	// 24 hours
	expiresin := 86400
	fmt.Println("secret", secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Duration(expiresin) * time.Second).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return 0, "", err
	}

	return expiresin, tokenString, nil
}

func GenerateRefreshToken(id, email, role string) (string, error) {
	secret := os.Getenv("JWT_PRIVATE_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
