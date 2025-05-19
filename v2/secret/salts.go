package secret

import (
	"crypto/rand"
	"errors"
)

// Generates a cryptografic secure salt of the given size
func GenerateSalt(size int) ([]byte, error) {
	if size < 1 {
		return nil, errors.New("size must be greater than 0")
	}
	salt := make([]byte, size)
	_, err := rand.Read(salt) // Use crypto/rand to generate cryptographically secure bytes
	if err != nil {
		return nil, err
	}
	return salt, nil
}
