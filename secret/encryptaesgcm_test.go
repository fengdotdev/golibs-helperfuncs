package secret_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/fengdotdev/golibs-helperfuncs/data"
	"github.com/fengdotdev/golibs-helperfuncs/secret"
	"github.com/fengdotdev/golibs-testing/assert"
)

type additionalData struct {
	Algorithm string `json:"algorithm"` // AES
	Mode      string `json:"mode"`      // GCM
	Strength  int    `json:"strength"`  // 256
	IV64      string `json:"iv64"`      // ex: 32bVr0KW+Cj5pPLB
}

const (
	text1          = "Hola Mundo"
	key64_1        = "ZOkodKmzHIMwBI3RtvRlSo4dKQWU5bM3+lKKIvmSy3w="
	iv64_1         = "32bVr0KW+Cj5pPLB"
	ciphertext64_1 = "WKBqzxm+x6R2sg5+0e2XLXGpC9QuY68wfiQ="

	text2              = "to go server"
	key64_2            = "T26jQLcmz49b/UU0exWOblxdEaBlSED96TPlnl89U9k="
	iv64_2             = "p2C0G98HwCajYXur"
	ciphertext64_2     = "AjIfZOH5FGwvGea9LMCRpal6DjyDRPksZpXJ0A=="
	additionalData64_2 = "eyJhbGdvcml0aG0iOiJBRVMiLCJtb2RlIjoiR0NNIiwic3RyZW5ndGgiOjI1NiwiaXY2NCI6InAyQzBHOThId0NhallYdXIifQ=="
)



func TestAESGCM(t *testing.T) {

	text := "Hola Mundo"

	data := []byte(text)

	key, err := secret.Generate256Key()
	assert.Nil(t, err)

	iv, err := secret.GenerateIV()
	assert.Nil(t, err)

	ciphertext, err := secret.EncryptAESGCM(key, iv, data, nil)
	assert.Nil(t, err)

	assert.False(t, bytes.Equal(data, ciphertext))
	assert.NotNil(t, ciphertext)

	decrypted, err := secret.DecryptAESGCM(key, iv, ciphertext, nil)
	assert.Nil(t, err)

	assert.NotNil(t, decrypted)

	result := bytes.Equal(data, decrypted)

	assert.True(t, result)

	assert.Equal(t, text, string(decrypted))

}

type additionaldata struct {
	Algorythm string `json:"algorythm"` // AES
	Mode      string `json:"mode"`      // GCM
	Strength  int    `json:"strength"`  // 256
	IV64      string `json:"iv64"`      // ex: 32bVr0KW+Cj5pPLB
}

type packagedata struct {
	Cypher64       string         `json:"cypher64"`
	Additionaldata additionaldata `json:"additionaldata"`
}

func TestAESGCMwithADD(t *testing.T) {
	text := "lorem ipsum dolor sit amet"

	binarydata := []byte(text)

	key, err := secret.Generate256Key()
	assert.NilWithMessage(t, err, "error generating key")

	iv, err := secret.GenerateIV()
	assert.NilWithMessage(t, err, "error generating iv")

	iv64 := data.Encode64Bytes(iv)

	add := additionaldata{
		Algorythm: "AES",
		Mode:      "GCM",
		Strength:  256,
		IV64:      iv64,
	}

	addjson, err := json.Marshal(add)
	assert.Nil(t, err)

	ciphertext, err := secret.EncryptAESGCM(key, iv, binarydata, addjson)
	assert.NilWithMessage(t, err, "error encrypting")

	assert.FalseWithMessage(t, bytes.Equal(binarydata, ciphertext), "binarydata and ciphertext are equal")
	assert.NotNil(t, ciphertext)

	// sent over the wire

	pkg := packagedata{
		Cypher64:       data.Encode64Bytes(ciphertext),
		Additionaldata: add,
	}

	var buf bytes.Buffer

	enc := json.NewEncoder(&buf)
	err = enc.Encode(pkg)
	assert.NilWithMessage(t, err, "error encoding at sender")

	// received from the wire

	var pkg1 packagedata

	dec := json.NewDecoder(&buf)
	err = dec.Decode(&pkg1)
	assert.NilWithMessage(t, err, "error decoding at receiver")

	reciveCypher, err := data.Decode64Bytes(pkg1.Cypher64)
	assert.NilWithMessage(t, err, "error decoding cypher at receiver")
	reciveadd := pkg1.Additionaldata
	reciveaddjson, err := json.Marshal(reciveadd)
	assert.NilWithMessage(t, err, "error marshalling additional data at receiver")

	assert.Equal(t, add.Algorythm, reciveadd.Algorythm)
	assert.Equal(t, add.Mode, reciveadd.Mode)
	assert.Equal(t, add.Strength, reciveadd.Strength)
	assert.Equal(t, add.IV64, reciveadd.IV64)

	ivprime, err := data.Decode64Bytes(reciveadd.IV64)
	assert.Nil(t, err)

	assert.TrueWithMessage(t, bytes.Equal(iv, ivprime), "iv and ivprime are not equal")

	decrypted, err := secret.DecryptAESGCM(key, ivprime, reciveCypher, reciveaddjson)
	assert.NilWithMessage(t, err, "error decrypting")

	assert.NotNilWithMessage(t, decrypted, "decrypted is nil")

	result := bytes.Equal(binarydata, decrypted)

	assert.TrueWithMessage(t, result, "binarydata and decrypted are not equal")

	assert.EqualWithMessage(t, text, string(decrypted), "text and decrypted are not equal")
}

func TestAESGCM_WithBase64Kwowns1(t *testing.T) {
	key1, err := data.Decode64Bytes(key64_1)
	assert.Nil(t, err)
	iv1, err := data.Decode64Bytes(iv64_1)
	assert.Nil(t, err)
	ciphertext1, err := data.Decode64Bytes(ciphertext64_1)
	assert.Nil(t, err)

	decrypted, err := secret.DecryptAESGCM(key1, iv1, ciphertext1, nil)
	assert.Nil(t, err)
	assert.NotNil(t, decrypted)

	decryptedtext := string(decrypted)
	assert.Equal(t, text1, decryptedtext)

}

func TestAESGCM_WithBase64Kwowns2(t *testing.T) {
	key2, err := data.Decode64Bytes(key64_2)
	assert.Nil(t, err)
	iv2, err := data.Decode64Bytes(iv64_2)
	assert.Nil(t, err)
	ciphertext2, err := data.Decode64Bytes(ciphertext64_2)
	assert.Nil(t, err)
	additionalData2,err:= data.Decode64Bytes(additionalData64_2)
	assert.Nil(t, err)
	decrypted, err := secret.DecryptAESGCM(key2, iv2, ciphertext2, additionalData2)
	assert.Nil(t, err)
	assert.NotNil(t, decrypted)

	decryptedtext := string(decrypted)
	assert.Equal(t, text2, decryptedtext)

}
