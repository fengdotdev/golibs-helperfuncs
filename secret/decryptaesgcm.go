package secret

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func DecryptAESGCM(key, iv, data, additionalData []byte) ([]byte, error) {

	err := AssertAESGCM_RequirementsOrErr(key, iv, data, additionalData)
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
		return nil, errors.New("el IV tiene un tamaño incorrecto")
	}



	plainText, err := aesGCM.Open(nil, iv, data, additionalData)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}
