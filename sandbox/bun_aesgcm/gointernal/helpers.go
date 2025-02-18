package gointernal

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// Encode64Bytes returns the base64 encoding of the given byte slice.
func Encode64Bytes(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Encode64 returns the base64 encoding of the given string.
func Encode64(data string) string {
	return Encode64Bytes([]byte(data))
}

// Decode64Bytes decodes the base64 encoded string and returns the corresponding byte slice.
func Decode64Bytes(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

// Decode64 decodes the base64 encoded string and returns the corresponding decoded string.
func Decode64(data string) (string, error) {
	decodedBytes, err := Decode64Bytes(data)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}

func DecryptAESGCM(key, iv, data, additionalData []byte) ([]byte, error) {

	//err := AssertAESGCM_RequirementsOrErr(key, iv, data, additionalData)

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

	plainText, err := aesGCM.Open(nil, iv, data, additionalData)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

