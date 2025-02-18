package gointernal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ENCODED(w http.ResponseWriter, r *http.Request) {

	///----ONLY FOR TESTING PURPOSES
	queryValue := r.URL.Query().Get("key")

	fmt.Println(queryValue)

	key, err := Decode64Bytes(queryValue)
	if err != nil {
		http.Error(w, "Error al decodificar la clave", http.StatusBadRequest)
		return
	}

	///----ONLY FOR TESTING PURPOSES

	var payload Payload

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, "Error al parsear JSON", http.StatusBadRequest)
		return
	}

	fmt.Println(payload)
	//decode

	data, err := payload.GetCypher()
	if err != nil {
		http.Error(w, "Error al decodificar data", http.StatusBadRequest)
		return
	}

	iv, err := payload.Additionaldata.GetIV()
	if err != nil {
		http.Error(w, "Error al decodificar iv", http.StatusBadRequest)
		return
	}

	addjson, err := payload.GetAdditionalDataAsBinary()
	if err != nil {
		http.Error(w, "Error al decodificar addjson", http.StatusBadRequest)
		return
	}

	//decrypt

	plainText, err := DecryptAESGCM(key, iv, data, addjson)
	if err != nil {
		http.Error(w, "Error al desencriptar", http.StatusBadRequest)
		return
	}
	w.Write([]byte(plainText))
}
