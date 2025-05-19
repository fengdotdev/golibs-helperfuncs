package datum_test

import (
	"bytes"
	"encoding/hex"
	"os"
	"strings"
	"testing"

	"github.com/fengdotdev/golibs-helperfuncs/v2/datum"
	"github.com/fengdotdev/golibs-testing/assert"
)

const (
	txt64Foo       = "foo"
	encode64Foo    = "Zm9v"
	txt64Empty     = ""
	encode64Empty  = ""
	txt64Foobar    = "foobar"
	encode64Foobar = "Zm9vYmFy"
	txt64F         = "f"
	encode64F      = "Zg=="
)

func TestEncode64(t *testing.T) {
	assert.Equal(t, encode64Foo, datum.Encode64(txt64Foo))
	assert.Equal(t, encode64Empty, datum.Encode64(txt64Empty))
	assert.Equal(t, encode64Foobar, datum.Encode64(txt64Foobar))
	assert.Equal(t, encode64F, datum.Encode64(txt64F))
}

func TestEncode64Bytes(t *testing.T) {

	b := []byte(txt64Foo)
	assert.Equal(t, encode64Foo, datum.Encode64Bytes(b))

	b = []byte(txt64Empty)
	assert.Equal(t, encode64Empty, datum.Encode64Bytes(b))

	b = []byte(txt64Foobar)
	assert.Equal(t, encode64Foobar, datum.Encode64Bytes(b))

	//hex sample
	{

		s := "Hello Gopher!"
		src := []byte(s)
		encodehex := make([]byte, hex.EncodedLen(len(src)))
		hex.Encode(encodehex, src)

		//encode
		encodeHelloGopher := datum.Encode64Bytes(encodehex)

		//decode
		decodeHelloGopher, err := datum.Decode64Bytes(encodeHelloGopher)
		assert.Nil(t, err)

		if !bytes.Equal(encodehex, decodeHelloGopher) {
			t.Errorf("Not equal")
		}

		//decode hex
		decodehex := make([]byte, hex.DecodedLen(len(decodeHelloGopher)))
		_, err = hex.Decode(decodehex, decodeHelloGopher)
		assert.Nil(t, err)

		//to string
		s2 := string(decodehex)
		assert.Equal(t, s, s2)
	}

	//file sample this May dont work in test if file not found and just skip
	{
		//read file
		rawpath := "test_assets/file-type-go-gopher.svg"
		wd, _ := os.Getwd()
		//remove /data
		toRemove := "/data"
		wd = strings.TrimSuffix(wd, toRemove)
		file := wd + "/" + rawpath
		filelen, err := os.Stat(file)

		// may be file not found in test
		if err == nil {
			filedata := make([]byte, filelen.Size())
			f, err := os.Open(file)
			assert.Nil(t, err)
			_, err = f.Read(filedata)
			assert.Nil(t, err)
			f.Close()

			//encode
			encodeFile := datum.Encode64Bytes(filedata)

			//decode
			decodeFile, err := datum.Decode64Bytes(encodeFile)
			assert.Nil(t, err)

			if !bytes.Equal(filedata, decodeFile) {
				t.Errorf("Not equal")
			}
		}

	}

}

func TestDecode64(t *testing.T) {
	decode, err := datum.Decode64(encode64Foo)
	assert.Nil(t, err)
	assert.Equal(t, txt64Foo, decode)

	decode, err = datum.Decode64(encode64Empty)
	assert.Nil(t, err)
	assert.Equal(t, txt64Empty, decode)

	decode, err = datum.Decode64(encode64Foobar)
	assert.Nil(t, err)
	assert.Equal(t, txt64Foobar, decode)
}

func TestDecode64Bytes(t *testing.T) {
	decode, err := datum.Decode64Bytes(encode64Foo)
	assert.Nil(t, err)
	assert.Equal(t, txt64Foo, string(decode))

	decode, err = datum.Decode64Bytes(encode64Empty)
	assert.Nil(t, err)
	assert.Equal(t, txt64Empty, string(decode))

	decode, err = datum.Decode64Bytes(encode64Foobar)
	assert.Nil(t, err)
	assert.Equal(t, txt64Foobar, string(decode))

}
