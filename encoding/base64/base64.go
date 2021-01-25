package base64

import (
	"bytes"
	"encoding/base64"
)

func Encode(plaintext []byte) []byte {
	ret := make([]byte, base64.StdEncoding.EncodedLen(len(plaintext)))
	base64.StdEncoding.Encode(ret, plaintext)
	return ret
}

func Decode(encodedBytes []byte) ([]byte, error) {
	ret := make([]byte, base64.StdEncoding.DecodedLen(len(encodedBytes)))
	_, err := base64.StdEncoding.Decode(ret, encodedBytes)
	if err != nil {
		return nil, err
	}
	return bytes.Trim(ret, "\x00"), nil
}
