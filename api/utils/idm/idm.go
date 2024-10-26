package idm

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"hestia/api/utils/db"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

// RandomSalt generates a random salt for password hashing
func RandomSalt() string {
	// The length of the salt
	saltLength := 32
	bytes := make([]byte, saltLength)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Error().Err(err).Msg("Error generating random salt")
		return ""
	}
	return hex.EncodeToString(bytes)
}

// GetSalt retrieves the salt for a user from the database using the email
func GetSalt(ctx context.Context, email string) (string, error) {
	db, err := db.GetDBPoolConn()
	if err != nil {
		log.Error().Err(err).Msg("Error getting database connection")
		return "", err
	}

	var salt string
	err = db.QueryRow(ctx, "SELECT salt FROM users.users WHERE email = $1", email).Scan(&salt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", errors.New("no rows found")
		}
		log.Error().Err(err).Str("email", email).Msg("Error getting salt from database")
		return "", err
	}
	return salt, nil
}

// PasswordHash turns a password into a encrypted hash (using various methods, such as SHA-256, ...)
func PasswordHash(password string, salt string) string {
	// First we verify that the password is a SHA-256 hash
	if len(password) != 64 {
		log.Warn().Msg("Password is not a SHA-256 hash")
		return ""
	}
	// Verify the salt isn't empty
	if salt == "" {
		log.Warn().Msg("Salt is empty")
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
