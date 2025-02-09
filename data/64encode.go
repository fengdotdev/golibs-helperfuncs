package data

import "encoding/base64"

// Encode64Bytes returns the base64 encoding of the given byte slice.
func Encode64Bytes(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Encode64 returns the base64 encoding of the given string.
func Encode64(data string) string {
	return Encode64Bytes([]byte(data))
}

// Decode64Bytes decodes the base64 encoded string and returns the corresponding byte slice.
func Decode64Bytes(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

// Decode64 decodes the base64 encoded string and returns the corresponding decoded string.
func Decode64(data string) (string, error) {
	decodedBytes, err := Decode64Bytes(data)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}
