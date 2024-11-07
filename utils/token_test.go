package utils

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	secretKey := "testsecretkey"
	payload := "testuser"
	ttl := time.Minute * 5

	token, err := GenerateToken(ttl, payload, secretKey)

	// Ensure no error occurred and the token is not empty
	assert.NoError(t, err, "expected no error while generating token")
	assert.NotEmpty(t, token, "expected non-empty token")

	// Optional: Parse the token to verify claims directly (checking it has a correct "sub" and "exp")
	parsedToken, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	assert.NoError(t, err, "expected no error while parsing token")
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		assert.Equal(t, payload, claims["sub"], "expected payload to match")
		assert.Greater(t, int64(claims["exp"].(float64)), time.Now().Unix(), "expected token expiration to be in the future")
	} else {
		t.Errorf("failed to parse token claims")
	}
}

func TestValidateToken(t *testing.T) {
	secretKey := "testsecretkey"
	payload := "testuser"
	ttl := time.Minute * 5

	// Generate a valid token
	token, err := GenerateToken(ttl, payload, secretKey)
	assert.NoError(t, err, "expected no error while generating token")

	// Validate the token
	result, err := ValidateToken(token, secretKey)
	assert.NoError(t, err, "expected no error while validating token")
	assert.Equal(t, payload, result, "expected result to match payload")

	// Test with an expired token
	expiredToken, _ := GenerateToken(-time.Minute*5, payload, secretKey)
	_, err = ValidateToken(expiredToken, secretKey)
	assert.Error(t, err, "expected error for expired token")

	// Test with an incorrect secret key
	_, err = ValidateToken(token, "wrongsecretkey")
	assert.Error(t, err, "expected error for token validated with wrong secret key")
}
