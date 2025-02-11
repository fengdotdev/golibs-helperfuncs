package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"

	"golang.org/x/crypto/pbkdf2"

	"github.com/fengdotdev/golibs-helperfuncs/asserty"
	"github.com/fengdotdev/golibs-helperfuncs/data"
	"github.com/fengdotdev/golibs-helperfuncs/secret"
)

type ExampleAES256GCM struct {
	Password string
	Salt     []byte
	Key      []byte
}

func main() {

	password := "password"

	s := "Hello, playground"
	fmt.Println(s)

	//to base64
	dataforcypher := data.Encode64(s)
	fmt.Println("data to cypher", dataforcypher)

	salt, err := secret.GenerateSalt256()
	asserty.AssertNoError(err)

	key := GenerateKey256(password, salt)



	// Encrypt the data

	// convert key to [32]byte
	var key32 [32]byte
	copy(key32[:], key)
	fmt.Println("key32", data.Encode64Bytes(key32[:]))
	r, err := Encrypt([]byte(s), &key32)
	asserty.AssertNoError(err)
	fmt.Println("encrypted data", r)
	fmt.Println("encrypted data base64", data.Encode64Bytes(r))
}

type AES256GCMResult struct {
	Nonce64      string
	CypherData64 string
	Tag64        string
	Combine64    string
}

func EncryptAES256GCM(data string, key []byte) (AES256GCMResult, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return AES256GCMResult{}, err
	}

	// 2. Wrap block with GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return AES256GCMResult{}, err
	}

	// 3. Generate a random nonce
	nonce, err := GenerateNonce(gcm.NonceSize())

	if err != nil {
		return AES256GCMResult{}, err
	}
	// 4. Encrypt the data
	// Seal appends the ciphertext and GCM tag to the first argument (dst),
	// which we've passed as nil to create a new slice.
	ciphertext := gcm.Seal(nil, nonce, []byte(data), nil)

	// 5. Prepend the nonce to the ciphertext
	combined := append(nonce, ciphertext...)

	combined64 := base64.StdEncoding.EncodeToString(combined)
	nonce64 := base64.StdEncoding.EncodeToString(nonce)
	cypherData64 := base64.StdEncoding.EncodeToString(ciphertext)

	return AES256GCMResult{
		Nonce64:      nonce64,
		CypherData64: cypherData64,
		Tag64:        "",
		Combine64:    combined64,
	}, nil
}

// Generate a random N-byte nonce
func GenerateNonce(n int) ([]byte, error) {
	nonce := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return nonce, nil
}

// GenerateKey256 derives a 256-bit key using PBKDF2 with SHA-256.
// Parameters:
//   - password: The password to derive the key from
//   - salt: Random salt value (should be at least 16 bytes)
//
// Returns:
//   - []byte: A 32-byte (256-bit) derived key
func GenerateKey256(password string, salt []byte) []byte {
	if len(password) == 0 {
		panic("password cannot be empty")
	}
	if len(salt) < 16 {
		panic("salt must be at least 16 bytes")
	}

	bytepassword := []byte(password)
	iterations := 100000 // is a reasonable default; increase if performance allows
	keylen := 32         // 256 bits
	hasher := sha256.New
	result := pbkdf2.Key(bytepassword, salt, iterations, keylen, hasher)
	AssetSize256(result)
	return result
}

// AssetSize256 asserts that the given byte slice is 32 bytes long.

func AssetSize256(data []byte) {
	if len(data) != 32 {
		panic("Assertion failed")
	}
}

func Encrypt(plaintext []byte, key *[32]byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}
