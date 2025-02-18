package secret

import "errors"

var (
	ErrPlaintextNil = errors.New("Plaintext cannot be nil")
)

func AssertPlaintextOrErr(plaintext []byte) error {
	if plaintext == nil {
		return ErrPlaintextNil
	}
	return nil
}
