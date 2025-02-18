package gointernal

import (
	"encoding/base64"
	"encoding/json"
)

type AdditionalData struct {
	Algorithm string `json:"algorithm"` // AES
	Mode      string `json:"mode"`      // GCM
	Strength  int    `json:"strength"`  // 256
	IV64      string `json:"iv64"`      // ex: 32bVr0KW+Cj5pPLB
}

func (a *AdditionalData) GetIV() ([]byte, error) {
	return base64.StdEncoding.DecodeString(a.IV64)
}

type Payload struct {
	Cypher64       string         `json:"cypher64"`
	Additionaldata AdditionalData `json:"additionaldata"`
}

func (p *Payload) GetCypher() ([]byte, error) {
	return base64.StdEncoding.DecodeString(p.Cypher64)
}

func (p *Payload) GetAdditionalDataAsBinary() ([]byte, error) {
	jsonData, err := json.Marshal(p.Additionaldata)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

type PayloadWithKey struct {
	Key64   string  `json:"key64"`
	Payload Payload `json:"payload"`
}

func (p *PayloadWithKey) GetKey() ([]byte, error) {
	return base64.StdEncoding.DecodeString(p.Key64)
}
