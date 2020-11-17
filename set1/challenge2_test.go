package cryptopals

import (
	"bytes"
	"github.com/juggler434/crypto/encoding/hex"
	"testing"
)

func TestFixedXor(t *testing.T) {
	tests := []struct{
		name string
		input1 []byte
		input2 []byte
		expectedResult []byte
		shouldError bool
	} {
		{
			name: "base case",
			input1: []byte("1c0111001f010100061a024b53535009181c"),
			input2: []byte("686974207468652062756c6c277320657965"),
			expectedResult: []byte("746865206b696420646f6e277420706c6179"),
			shouldError: false,
		},
		{
			name: "mismatched input lengths",
			input1: []byte("1c0111001f010100061a024b53535009181c"),
			input2: []byte("686974207468652062756c6c27732065796568"),
			expectedResult: nil,
			shouldError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			di1, err := hex.Decode(test.input1)
			if err != nil {
				t.Fatal("Error decoding hex input 1")
			}
			di2 ,err := hex.Decode(test.input2)
			if err != nil {
				t.Fatal("Error decoding hex input 2")
			}

			res, err := FixedXor(di1, di2)
			if test.shouldError {
				if err == nil {
					t.Error("expected: error, got: nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected nil error, got: %s", err)
				}
			}

			hres := hex.Encode(res)
			if !bytes.Equal(hres, test.expectedResult) {
				t.Errorf("expected: %s, got %s", test.expectedResult, hres)
			}

		})
	}
}
