package cryptopals

import (
	"encoding/hex"
)

//EncryptWithRepeatingXor XOR encrypts a given string with a given key.  Returns a hex encoded string.
func EncryptWithRepeatingXor(input, key []byte) ([]byte, error) {
	eb := make([]byte, len(input))

	for i := 0; i < len(input); i++ {
		eb[i] = input[i] ^ key[i%len(key)]
	}

	ret := make([]byte, hex.EncodedLen(len(eb)))
	hex.Encode(ret, eb)
	return ret, nil
}
