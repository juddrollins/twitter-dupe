package util

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var signingSecret = []byte("this is a secret")

// Create a JWT token
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = "username"

	tokenString, err := token.SignedString(signingSecret)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

// Validate a JWT token
func ValidateJWT() (bool, error) {
	tokenString, err := GenerateJWT()
	if err != nil {
		return false, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingSecret, nil
	})

	if err != nil {
		return false, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, nil
	}

	return false, nil
}
