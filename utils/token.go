package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateAccessToken(id, email, role string) (int, string, error) {
	secret := os.Getenv("JWT_PRIVATE_KEY")
	privateKey, _ := os.ReadFile(secret)
	// 24 hours
	expiresin := 86400

	key, _ := jwt.ParseRSAPrivateKeyFromPEM(privateKey)

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"id":    id,
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Second * time.Duration(expiresin)).Unix(),
	})

	tokenString, err := token.SignedString(key)
	if err != nil {
		return 0, "", err
	}

	return expiresin, tokenString, nil
}

func GenerateRefreshToken(id, email, role string) (string, error) {
	secret := os.Getenv("JWT_PRIVATE_KEY")
	privateKey, _ := os.ReadFile(secret)

	key, _ := jwt.ParseRSAPrivateKeyFromPEM(privateKey)

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"id":    id,
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
