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

	key, err := Decode64(queryValue)
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

	plainText, err := DecodeAESGCM(key, data, iv , nil)
	if err != nil {
		http.Error(w, "Error al desencriptar", http.StatusBadRequest)
		return
	}
	w.Write([]byte(plainText))
}
