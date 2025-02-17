package gointernal

import "encoding/base64"

type EncodeAESGCM_Object struct {
	DATA64    string `json:"data64"`
	IV64      string `json:"iv64"`
	AUTHTAG64 string `json:"authtag64"`
}

func NewEncodeAESGCM_Object(data []byte, iv []byte, authtag []byte) EncodeAESGCM_Object {
	return EncodeAESGCM_Object{
		DATA64:    base64.StdEncoding.EncodeToString(data),
		IV64:      base64.StdEncoding.EncodeToString(iv),
		AUTHTAG64: base64.StdEncoding.EncodeToString(authtag),
	}
}

func (e *EncodeAESGCM_Object) Data() ([]byte, error) {
	return base64.StdEncoding.DecodeString(e.DATA64)
}

func (e *EncodeAESGCM_Object) IV() ([]byte, error) {
	return base64.StdEncoding.DecodeString(e.IV64)
}

func (e *EncodeAESGCM_Object) AuthTag() ([]byte, error) {
	return base64.StdEncoding.DecodeString(e.AUTHTAG64)
}
