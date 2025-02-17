package gointernal

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

func Decode64(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

func Encode64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func AssertKey(key []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("la clave debe tener 32 bytes para AES-256")
	}
	return key, nil
}

func EncodeAESGCM(key, data, iv, additionalData []byte) (cypherdata []byte, err error) {
	if len(key) != 32 {
		return nil, errors.New("la clave debe tener 32 bytes para AES-256")
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
		return nil, errors.New("el IV tiene un tamaño incorrecto")
	}
	var buf []byte
	tag := aesGCM.Seal(buf, iv, data, additionalData)
	dataAndTag := append(buf, tag...)
	return dataAndTag, nil
}

func DecodeAESGCM(key, data, iv, additionalData []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("la clave debe tener 32 bytes para AES-256")
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
		return nil, errors.New("el IV tiene un tamaño incorrecto")
	}

	// Decrypt with combined ciphertext

	plainText, err := aesGCM.Open(nil, iv, data, additionalData)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}
