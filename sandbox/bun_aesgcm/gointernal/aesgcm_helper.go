package gointernal

import (
	"encoding/base64"
	"errors"
)

const (
	ErrorKeySize = "the key must have 32 bytes for AES-256"
	ErrorIVSize  = "the IV has an incorrect size"
)

var _ AESGCM_HelperInterface = &AESGCM_Helper{}

type AESGCM_Helper struct {
	key []byte
	iv  []byte
}

func NewAESGCM_HelperFromSeed(seed string) (AESGCM_Helper, error) {
	panic("not implemented")
}

func NewAESGCM_HelperFromSeedData(seed []byte) (AESGCM_Helper, error) {
	panic("not implemented")
}

// NewAESGCM_HelperFromBase64 returns an AESGCM_Helper object and a error from base64  encoded key and iv
func NewAESGCM_HelperFromBase64(key64, iv64 string) (AESGCM_Helper, error) {

	key, err := base64.StdEncoding.DecodeString(key64)
	if err != nil {
		return AESGCM_Helper{}, err
	}

	iv, err := base64.StdEncoding.DecodeString(iv64)
	if err != nil {
		return AESGCM_Helper{}, err
	}

	return NewAESGCM_Helper(key, iv)
}

// returns a new AESGCM_Helper object from key and iv or an error
func NewAESGCM_Helper(key, iv []byte) (AESGCM_Helper, error) {

	if len(key) != 32 {
		return AESGCM_Helper{}, errors.New(ErrorKeySize)
	}

	if len(iv) != 12 {
		return AESGCM_Helper{}, errors.New(ErrorIVSize)
	}

	return AESGCM_Helper{
		key: key,
		iv:  iv,
	}, nil
}

// returns the key in base64
func (a *AESGCM_Helper) GetKey64() string {

	return base64.StdEncoding.EncodeToString(a.key)
}

// returns the iv in base64
func (a *AESGCM_Helper) GetIV64() string {
	return base64.StdEncoding.EncodeToString(a.iv)
}

// returns the key in bytes
func (a *AESGCM_Helper) Getkey() []byte {
	return a.key
}

// returns the iv in bytes
func (a *AESGCM_Helper) GetIV() []byte {
	return a.iv
}
