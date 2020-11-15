package cryptopals

import (
	"errors"
	"github.com/juggler434/crypto/encoding/hex"
)

// FixedXor takes 2 hex encoded byte slices, and returns a hex encoded byte slice of their xor combinations
func FixedXor(input1, input2 []byte) ([]byte, error) {
	dc1, err := hex.Decode(input1)
	if err != nil {
		return nil, err
	}
	dc2, err := hex.Decode(input2)
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

	hret := hex.Encode(ret)
	return hret, nil
}
