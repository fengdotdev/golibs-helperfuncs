package data

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

// ValidateSHA256 validates if the hash is the SHA256 of the string data
// ex: foo, 2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae -> nil
func ValidateSHA256(data string, hash string) error {
	computed := GetSHA256(data)
	if computed != hash {
		return errors.New("invalid SHA256 hash")
	}
	return nil
}

// ValidateSHA256Bytes validates if the hash is the SHA256 of the data bytes
func ValidateSHA256Bytes(data []byte, hash string) error {
	computed := GetSHA256Bytes(data)
	if computed != hash {
		return errors.New("invalid SHA256 hash")
	}
	return nil
}

// GetSHA256 returns the SHA256 hash of the string data
// ex: foo -> 2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae
func GetSHA256(data string) string {
	return GetSHA256Bytes([]byte(data))
}

// GetSHA256Bytes returns the SHA256 hash of the data bytes
func GetSHA256Bytes(data []byte) string {
	hash := sha256.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}
