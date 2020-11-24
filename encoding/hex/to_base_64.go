// Package cryptopals houses all the business logic for the crypto pals challenges.
package hex

import (
	"github.com/juggler434/crypto/encoding/base64"
)


//ToBase64 takes a slice of hex encoded bytes, and returns a slice of base64 encoding bytes
func ToBase64(hexBytes []byte) ([]byte, error) {
	b, err := Decode(hexBytes)
	if err != nil {
		return nil, err
	}

	ret := base64.Encode(b)
	return ret, nil
}
