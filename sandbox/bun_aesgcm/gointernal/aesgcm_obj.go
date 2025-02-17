package gointernal

import (
	"encoding/json"

	traits "github.com/fengdotdev/golibs-traits"
)

var _ traits.JSONTrait = &AESGCM_Object{}

var _ AESGCM_DataInterface = &AESGCM_Object{}

// AESGCM_DTO is a struct that represents a AESGCM object and holds the data in the cypheredForm, iv and additional data (optional and uncyphered) in base64
// this is the struct that is used to parse the json object
// this json object is used to send the data in the body of the request
type AESGCM_DTO struct {
	Data64           string `json:"data64"`           // this is cypher and tag encoded in base64
	IV64             string `json:"iv64"`             // this is the iv or nonce encoded in base64
	AdditionalData64 string `json:"additionalData64"` // thi
}

func NewAESGCM_Object(data64, iv64, additionalData64 string) *AESGCM_Object {
	return &AESGCM_Object{
		data64:           data64,
		iv64:             iv64,
		additionalData64: additionalData64,
	}
}

// AESGCM_Object is a struct that represents a AESGCM object and holds the data in the cypheredForm, iv and additional data (optional and uncyphered) in base64
type AESGCM_Object struct {
	data64           string // this is cypher and tag encoded in base64
	iv64             string // this is the iv or nonce encoded in base64
	additionalData64 string // this is the additional data encoded in base64 (optional) and uncyphered
}

// JSONTrait interface implementation

func (a *AESGCM_Object) FromJSON(data string) error {
	var dto AESGCM_DTO
	err := json.Unmarshal([]byte(data), &dto)
	if err != nil {
		return err
	}
	a.data64 = dto.Data64
	a.iv64 = dto.IV64
	a.additionalData64 = dto.AdditionalData64

	return nil
}

func (a *AESGCM_Object) ToJSON() (string, error) {

	var dto AESGCM_DTO
	dto.Data64 = a.data64
	dto.IV64 = a.iv64
	dto.AdditionalData64 = a.additionalData64

	data, err := json.Marshal(dto)
	if err != nil {
		return "", err
	}
	// TODO CHECK THIS
	panic("not implemented")
	str := string(data)

	return str, nil
}

// AESGCM_DataInterface interface implementation

// Data returns the data in the object
func (a *AESGCM_Object) Data() ([]byte, error) {
	return nil, nil
}

func (a *AESGCM_Object) IV() ([]byte, error) {
	return nil, nil
}

func (a *AESGCM_Object) AdditionalData() ([]byte, error) {
	return nil, nil
}

func (a *AESGCM_Object) IV64() string {
	return ""
}

func (a *AESGCM_Object) Data64() string {
	return ""
}

func (a *AESGCM_Object) AdditionalData64() string {
	return ""
}
