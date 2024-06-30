package idm

import (
	"testing"
)

func TestRandomSalt(t *testing.T) {
	salt := RandomSalt()

	if salt == "" {
		t.Errorf("RandomSalt() = %s, expected a non-empty string", salt)
	}
}

func TestPasswordHash(t *testing.T) {
	// SHA-256 hash of "password"
	password := "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"
	salt := "mySalt"

	expectedHash := "4a1a4f8362882bba5f9a8d484846319e9c82a601e87b3a29d72f16ba703d88d4"
	hash := PasswordHash(password, salt)

	if hash != expectedHash {
		t.Errorf("PasswordHash(%s, %s) = %s, expected %s", password, salt, hash, expectedHash)
	}
}
