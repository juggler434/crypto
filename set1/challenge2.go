package cryptopals

import (
	"encoding/hex"
	"errors"
)

// FixedXor takes 2 hex encoded byte slices, and returns a hex encoded byte slice of their xor combinations
func FixedXor(input1, input2 []byte) ([]byte, error) {
	dc1, err := decodeHexBytes(input1)
	if err != nil {
		return nil, err
	}
	dc2, err := decodeHexBytes(input2)
	if err != nil {
		return nil, err
	}

	if len(dc1) != len(dc2) {
		return nil, errors.New("the inputs have mismatched lengths")
	}
	ret := make([]byte, len(dc1))
	for i := 0; i < len(dc1); i++ {
		ret[i] = dc1[i] ^ dc2[i]
	}

	hret := make([]byte, hex.EncodedLen(len(ret)))
	hex.Encode(hret, ret)
	return hret, nil
}
