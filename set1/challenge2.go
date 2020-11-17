package cryptopals

import (
	"errors"
)

// FixedXor takes 2 byte slices and returns the XORd result
func FixedXor(input1, input2 []byte) ([]byte, error) {
	if len(input1) != len(input2) {
		return nil, errors.New("the inputs have mismatched lengths")
	}
	ret := make([]byte, len(input1))
	for i := 0; i < len(input1); i++ {
		ret[i] = input1[i] ^ input2[i]
	}

	return ret, nil
}
