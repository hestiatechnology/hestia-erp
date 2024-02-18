package idm

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hestia/api/utils/db"
	"log"
)

// RandomSalt generates a random salt for password hashing
func RandomSalt() string {
	// The length of the salt
	saltLength := 32
	bytes := make([]byte, saltLength)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Println("An error occured while generating random bytes:", err)
		return ""
	}
	return hex.EncodeToString(bytes)
}

// GetSalt retrieves the salt for a user from the database using the email
func GetSalt(ctx context.Context, email string) (string, error) {
	db, err := db.GetDbPoolConn()
	if err != nil {
		log.Println(err)
		return "", err
	}

	var salt string
	err = db.QueryRow(ctx, "SELECT salt FROM users.users WHERE email = $1", email).Scan(&salt)
	if err != nil {
		log.Println("Unable to get salt for "+email+": ", err)
		return "", err
	}
	return salt, nil
}

// PasswordHash turns a password into a encrypted hash (using various methods, such as SHA-256, ...)
func PasswordHash(password string, salt string) string {
	// First we verify that the password is a SHA-256 hash
	if len(password) != 64 {
		return ""
	}
	// Verify the salt isn't empty
	if salt != "" {
		return ""
	}

	// Make a SHA-256 hash of the password and salt
	hasher := sha256.New()
	hasher.Write([]byte(password + salt))
	sha256Hash := hasher.Sum(nil)

	// Encrypt using SHA-512 and the salt
	hasher = sha512.New()
	hasher.Write(append(sha256Hash, []byte(salt)...))
	sha512Hash := hasher.Sum(nil)

	// Re-encrypt using SHA-256 and the salt
	hasher = sha256.New()
	hasher.Write(append(sha512Hash, []byte(salt)...))
	finalHash := hasher.Sum(nil)

	// Return the final hash as a hexadecimal string
	return hex.EncodeToString(finalHash)
}
