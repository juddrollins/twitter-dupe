package util

import (
	"fmt"
	"testing"
	"time"

	"github.com/juddrollins/twitter-dupe/db"
)

func TestGenerateJWT(t *testing.T) {
	user := db.Entry{
		PK:   "testuser",
		SK:   "testuser",
		Data: "testuser::testpassword",
	}

	username := "testuser"
	expectedExpiresAt := time.Now().Add(10 * time.Minute).Unix()

	tokenString, err := GenerateJWT(user)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Verify that the token is not empty
	if tokenString == "" {
		t.Error("token string is empty")
	}

	// Verify that the token can be parsed
	claims, err := ParseJWT(tokenString)
	if err != nil {
		fmt.Print(err)
		t.Errorf("unexpected error: %v", err)
	}

	// Verify that the subject in the token matches the username
	if claims.Subject != username {
		t.Errorf("expected subject to be %s, got %s", username, claims.Subject)
	}

	// Verify that the token expires after 10 minutes
	if claims.ExpiresAt != expectedExpiresAt {
		t.Errorf("expected expiresAt to be %d, got %d", expectedExpiresAt, claims.ExpiresAt)
	}
}

func TestParseJWT(t *testing.T) {
	tokenString := "your_token_string_here"

	_, err := ParseJWT(tokenString)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	invalidTokenString := "invalid_token_format"
	_, err = ParseJWT(invalidTokenString)
	if err == nil {
		t.Error("expected error, but got nil")
	}

	// Test case: Expired token
	expiredTokenString := "expired_token_string"
	_, err = ParseJWT(expiredTokenString)
	if err == nil {
		t.Error("expected error, but got nil")
	}

	// Test case: Token with invalid signature
	invalidSignatureTokenString := "token_with_invalid_signature"
	_, err = ParseJWT(invalidSignatureTokenString)
	if err == nil {
		t.Error("expected error, but got nil")
	}
}
