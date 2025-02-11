package gointernal

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func decode64(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

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

func ENCODED(w http.ResponseWriter, r *http.Request) {

	///----ONLY FOR TESTING PURPOSES
	queryValue := r.URL.Query().Get("key")

	fmt.Println(queryValue)


	key, err := decode64(queryValue)
	if err != nil {
		http.Error(w, "Error al decodificar la clave", http.StatusBadRequest)
		return
	}

	///----ONLY FOR TESTING PURPOSES

	var dataobj EncodeAESGCM_Object

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &dataobj)
	if err != nil {
		http.Error(w, "Error al parsear JSON", http.StatusBadRequest)
		return
	}


	fmt.Println(dataobj)
	//decode

	data, err := dataobj.Data()
	if err != nil {
		http.Error(w, "Error al decodificar data", http.StatusBadRequest)
		return
	}

	iv, err := dataobj.IV()
	if err != nil {
		http.Error(w, "Error al decodificar iv", http.StatusBadRequest)
		return
	}

	//decrypt

	plainText, err := DecodeAESGCM(data, iv, key)
	if err != nil {
		http.Error(w, "Error al desencriptar", http.StatusBadRequest)
		return
	}
	w.Write([]byte(plainText))
}

func DecodeAESGCM(data []byte, iv []byte, key []byte) ([]byte, error) {
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

	// Desencriptar los datos
	plainText, err := aesGCM.Open(nil, iv, data, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}
