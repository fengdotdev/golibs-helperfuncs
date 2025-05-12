package main

import (
	"fmt"

	"github.com/fengdotdev/golibs-helperfuncs/datum"
	"github.com/fengdotdev/golibs-helperfuncs/secret"
)

func main() {

	fmt.Println("AES-GCM")

	key, err := secret.Generate256Key()
	if err != nil {
		panic(err)
	}

	iv, err := secret.GenerateIV()
	if err != nil {
		panic(err)
	}

	iv64 := datum.Encode64Bytes(iv)
	key64 := datum.Encode64Bytes(key)

	fmt.Println("Key:", key64)
	fmt.Println("IV:", iv64)

	text := "Hola Mundo"

	d := []byte(text)

	ciphertext, err := secret.EncryptAESGCM(key, iv, d, nil)
	if err != nil {
		panic("cipher err" + err.Error())
	}

	ciphertext64 := datum.Encode64Bytes(ciphertext)

	fmt.Println("Ciphertext:", ciphertext64)

	decrypted, err := secret.DecryptAESGCM(key, iv, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Decrypted:", string(decrypted))

}
