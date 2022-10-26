package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateAccessToken(id, email string) (string, error) {
	secret := os.Getenv("JWT_PRIVATE_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
