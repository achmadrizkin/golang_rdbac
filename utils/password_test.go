package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "mysecretpassword"

	// Generate hash for the password
	hashedPassword, err := HashPassword(password)

	// Ensure the hash was generated without errors and is non-empty
	assert.NoError(t, err, "expected no error while hashing password")
	assert.NotEmpty(t, hashedPassword, "expected hashed password to be non-empty")

	// Ensure that hashing the same password twice produces different hashes
	hashedPassword2, err := HashPassword(password)
	assert.NoError(t, err, "expected no error while hashing password again")
	assert.NotEqual(t, hashedPassword, hashedPassword2, "expected different hashes for the same password due to salting")
}

func TestVerifyPassword(t *testing.T) {
	password := "mysecretpassword"
	wrongPassword := "wrongpassword"

	// Hash the password
	hashedPassword, err := HashPassword(password)
	assert.NoError(t, err, "expected no error while hashing password")

	// Verify that the correct password matches the hash
	isValid := VerifyPassword(hashedPassword, password)
	assert.True(t, isValid, "expected password verification to succeed for correct password")

	// Verify that an incorrect password does not match the hash
	isInvalid := VerifyPassword(hashedPassword, wrongPassword)
	assert.False(t, isInvalid, "expected password verification to fail for incorrect password")
}
