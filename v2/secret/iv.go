package secret

import (
	"crypto/rand"
	"errors"
	"io"
)

var (
	ErrIVLength = errors.New("IV length must be 12 bytes")
)

func AssertIVOrErr(iv []byte) error {

	if iv == nil {
		return errors.New("IV cannot be nil")
	}

	if len(iv) != 12 {
		return ErrIVLength
	}

	return nil
}

// GenerateIV generates a 12 byte long IV or nonce
func GenerateIV() ([]byte, error) {
	iv := make([]byte, 12)
	_, err := rand.Read(iv)
	if err != nil {
		return nil, err
	}
	return iv, nil
}

// Generate a random N-byte nonce
func GenerateNonce(n int) ([]byte, error) {
	nonce := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return nonce, nil
}
