package secret

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"

	"golang.org/x/crypto/pbkdf2"
)

var (
	ErrKeyLength  = errors.New("Key length must be 16, 24 or 32 bytes")
	ErrKeyLength32 = errors.New("Key length must be 32 bytes")
	ErrKeyNil     = errors.New("Key cannot be nil")
)

func AssertKeyOrErr(key []byte) error {

	if key == nil {

		return ErrKeyNil
	}

	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return ErrKeyLength
	}

	return nil
}

func AssertSize256(key []byte) error {
	if len(key) != 32 {
		return ErrKeyLength32
	}
	return nil
}

func Generate256Key() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func GenerateKey256(password string, salt []byte) []byte {
	if len(password) == 0 {
		panic("password cannot be empty")
	}
	if len(salt) < 16 {
		panic("salt must be at least 16 bytes")
	}

	bytepassword := []byte(password)
	iterations := 100000 // is a reasonable default; increase if performance allows
	keylen := 32         // 256 bits
	hasher := sha256.New
	result := pbkdf2.Key(bytepassword, salt, iterations, keylen, hasher)
	AssertSize256(result)
	return result
}
