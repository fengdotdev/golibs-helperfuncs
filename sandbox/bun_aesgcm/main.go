package main

import (
	"bunexample/gointernal"
	internal "bunexample/gointernal"
	"encoding/json"
	"fmt"
)

func main() {
	key, err := internal.Decode64(key64)
	if err != nil {
		panic(err)
	}
	bplainText := []byte(plainText)
	fmt.Println(string(bplainText))
	iv, err := internal.Decode64(iv64)
	if err != nil {
		panic(err)
	}

	additionalData := []byte("additionalData")


	cypherText, err := internal.EncodeAESGCM(key, bplainText, iv, additionalData)
	if err != nil {
		panic(err.Error() + "Error al EncodeAESGCM")
	}

	fmt.Println(internal.Encode64(cypherText), internal.Encode64(iv), internal.Encode64(additionalData))

	decoded, err := internal.DecodeAESGCM(key, cypherText, iv, additionalData)
	if err != nil {
		panic(err.Error() + "Error al DecodeAESGCM")
	}

	fmt.Println(string(decoded))

}

func Foo() {
	// Create a new instance of the server
	server := gointernal.NewServer()

	// Start the server
	server.Run("8080")
}

const (
	iv64      string = "Kfsl1u9tlqLiHudv"
	plainText string = "Hello from Bun with AES-256-GCM!"
	j         string = `{"data64":"RxrK2/BDLwz6bH6TSmHw/mdgNXzObnp5+6LJnhgX9nk=","iv64":"Kfsl1u9tlqLiHudv","authtag64":"ME6y99UzTK7TYxT1E7ShOw=="}`
	key64     string = "TH+35QFhuho+Cb4AR8IXtvbbLRtf+vsi4nSwFFEyxgY="
)

func Boo() {
	var obj internal.EncodeAESGCM_Object

	err := json.Unmarshal([]byte(j), &obj)
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)

	data, err := obj.Data()
	if err != nil {
		panic(err)
	}
	iv, err := obj.IV()
	if err != nil {
		panic(err)
	}

	tag, err := obj.AuthTag()
	if err != nil {
		panic(err)
	}
	fmt.Println(len(tag))

	key, err := internal.Decode64(key64)
	if err != nil {
		panic(err)
	}

	_, err = internal.AssertKey(key)
	if err != nil {
		panic(err)
	}

	plainText, err := internal.DecodeAESGCM(key, data, iv,nil)
	if err != nil {
		panic(err.Error() + "Error al DecodeAESGCM")
	}
	fmt.Println(string(plainText))
}
