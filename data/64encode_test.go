package data_test

import (
	"testing"

	"github.com/fengdotdev/golibs-helperfuncs/data"
	"github.com/fengdotdev/golibs-testing/assert"
)

const (
	peter   = "peter"
	peter64 = "cGV0ZXI="
	mike    = "mike"
	mike64  = "bWlrZQo="
)

func TestEncode64(t *testing.T) {
	assert.Equal(t, peter64, data.Encode64(peter))
	assert.Equal(t, mike64, data.Encode64(mike))
}

func TestEncode64Bytes(t *testing.T) {
	assert.Equal(t, peter64, data.Encode64Bytes([]byte(peter)))
	assert.Equal(t, mike64, data.Encode64Bytes([]byte(mike)))
}

func TestDecode64(t *testing.T) {
	decodedPeter, _ := data.Decode64(peter64)
	decodedMike, _ := data.Decode64(mike64)

	assert.Equal(t, peter, decodedPeter)
	assert.Equal(t, mike, decodedMike)
}

func TestDecode64Bytes(t *testing.T) {
	decodedPeter, _ := data.Decode64Bytes(peter64)
	decodedMike, _ := data.Decode64Bytes(mike64)

	assert.Equal(t, peter, string(decodedPeter))
	assert.Equal(t, mike, string(decodedMike))
}
