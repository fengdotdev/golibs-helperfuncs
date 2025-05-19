package secret

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)





var (
	ErrIVLength12 = errors.New("In AES256GCM IV or nonce length must be 12 bytes")
)

// EncryptAESGCM encrypts the plaintext using AES-GCM with the provided key.
// The key must be 16, 24 or 32 bytes long to select AES-128, AES-192 or AES-256.
// The iv must be 12 bytes long.
// The plaintext is the data that will be encrypted.
// AdditionalData is optional and can be nil, this is the data that will be authenticated but not encrypted
// The function returns the encrypted data.
func EncryptAESGCM(key, iv, plaintext, AdditionalData []byte) ([]byte, error) {

	err := AssertAESGCM_RequirementsOrErr(key, iv, plaintext, AdditionalData)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(iv) != aesGCM.NonceSize() {
		return nil, ErrIVLength12
	}
	//var buf []byte
	tag := aesGCM.Seal(nil, iv, plaintext, AdditionalData)
	//dataAndTag := append(buf, tag...)
	return tag, nil

}
