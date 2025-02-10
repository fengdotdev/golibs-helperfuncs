package data

import "crypto/aes"
import "golang.org/x/crypto/scrypt"


type EncryptorAES256GCM struct {
	key []byte
	aad []byte
}

func generateSalt() ([]byte, error) {
	return nil, nil
}

func generateKey(password string, salt []byte) ([]byte, error) {
	const (
		N      = 32768 // CPU/memory cost parameter (must be a power of 2, typically >= 2^14)
		r      = 8     // Block size parameter
		p      = 1     // Parallelization parameter
		keyLen = 32    // The desired key length in bytes (256-bit key)
	)

	key, err := scrypt.Key([]byte(password), salt, N, r, p, keyLen)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// NewEncryptorAES256GCM creates a new EncryptorAES256GCM with the given password and salt.
func NewEncryptorAES256GCM(password string, salt []byte) *EncryptorAES256GCM {

	return &EncryptorAES256GCM{password: password, salt: salt, aad: nil}
}

func NewEncryptorAES256GCMWithAAD(password string, salt []byte, aad []byte) *EncryptorAES256GCM {
	return &EncryptorAES256GCM{password: password, salt: salt, aad: aad}
}

func NewEncryptorAES256GCMNoSalt(password string) (encryptor *EncryptorAES256GCM, salt []byte) {
	salt, _ = generateSalt()
	return &EncryptorAES256GCM{password: password, salt: salt, aad: nil}, salt
}

func NewEncryptorAES256GCMNoSaltWithAAD(password string, aad []byte) (encryptor *EncryptorAES256GCM, salt []byte) {
	salt, _ = generateSalt()
	return &EncryptorAES256GCM{password: password, salt: salt, aad: aad}, salt
}

func (e *EncryptorAES256GCM) Encrypt(data []byte) ([]byte, error) {

	block, err := aes.NewCipher()
	if err != nil {
		return nil, nil, err
	}

}

func (e *EncryptorAES256GCM) Decrypt(data []byte) ([]byte, error) {
	return nil, nil
}

func (e *EncryptorAES256GCM) GenNonce() ([]byte, error) {

}

func (e *EncryptorAES256GCM) GenKey() ([]byte, error) {
	return nil, nil
}

func (e *EncryptorAES256GCM) GenPassword() ([]byte, error) {
	return nil, nil
}

func (e *EncryptorAES256GCM) GenIV() ([]byte, error) {
	return nil, nil
}

func (e *EncryptorAES256GCM) GenSalt() ([]byte, error) {
	return nil, nil
}
