package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var signingSecret = []byte("this is a secret")

type claims struct {
	Username   string    `json:"user"`
	Exp        time.Time `json:"exp"`
	Authorized bool      `json:"authorized"`
	jwt.StandardClaims
}

// Create a JWT token
func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["Exp"] = time.Now().Add(10 * time.Minute)
	claims["Authorized"] = true
	claims["User"] = username

	tokenString, err := token.SignedString(signingSecret)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

// Validate a JWT token
func ParseJWT(tokenString string) (*claims, error) {
	// Parse the JWT token with the provided secret key.
	token, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return signingSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid.
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Extract custom claims from the token.
	if claims, ok := token.Claims.(*claims); ok {
		return claims, nil
	}

	return nil, fmt.Errorf("failed to parse claims")
}
