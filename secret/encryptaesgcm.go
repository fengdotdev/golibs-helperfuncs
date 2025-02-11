package secret

import (
	//"crypto/aes"

)

type EncryptorAES256GCM struct {
	key []byte
	aad []byte
}


// NewEncryptorAES256GCM creates a new EncryptorAES256GCM with the given password and salt.
func NewEncryptorAES256GCM(password string, salt []byte) *EncryptorAES256GCM {
	panic("Not implemented NewEncryptorAES256GCM")
}

func NewEncryptorAES256GCMWithAAD(password string, salt []byte, aad []byte) *EncryptorAES256GCM {
	panic("Not implemented NewEncryptorAES256GCMWithAAD")
}

func NewEncryptorAES256GCMNoSalt(password string) (encryptor *EncryptorAES256GCM, salt []byte) {
	panic("Not implemented NewEncryptorAES256GCMNoSalt")

}

func NewEncryptorAES256GCMNoSaltWithAAD(password string, aad []byte) (encryptor *EncryptorAES256GCM, salt []byte) {
	panic("Not implemented NewEncryptorAES256GCMNoSaltWithAAD")

}

func (e *EncryptorAES256GCM) Encrypt(data []byte) ([]byte, error) {
	panic("Not implemented Encrypt")

}

func (e *EncryptorAES256GCM) Decrypt(data []byte) ([]byte, error) {
	panic("Not implemented Decrypt")
	return nil, nil
}

func (e *EncryptorAES256GCM) GenNonce() ([]byte, error) {

	panic("Not implemented GenNonce")
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
