package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// Signing secret is a uuid
var signingSecret = []byte("B433A9B4-24BF-4583-84F2-040E0F7763B0")

// Create a JWT token
func GenerateJWT(username string) (string, error) {
	claims := jwt.StandardClaims{
		Subject:   username,
		ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		IssuedAt:  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(signingSecret)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

// Validate a JWT token
func ParseJWT(tokenString string) (*jwt.StandardClaims, error) {
	// Parse the JWT token with the provided secret key.
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Extract custom claims from the token.
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if ok {
		return claims, nil
	}

	return nil, fmt.Errorf("failed to parse claims")
}
