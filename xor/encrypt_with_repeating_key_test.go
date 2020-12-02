package xor

import (
	"bytes"
	"github.com/juggler434/crypto/encoding/hex"
	"testing"
)

const (
	RepeatingXorTestString         = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	RepeatingXorExpectedTestResult = "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
)

func TestEncryptWithRepeatingKey(t *testing.T) {
	tests := []struct {
		name           string
		input          []byte
		key            []byte
		expectedOutput []byte
	}{
		{
			name:           "base case",
			input:          []byte(RepeatingXorTestString),
			key:            []byte("ICE"),
			expectedOutput: []byte(RepeatingXorExpectedTestResult),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := EncryptWithRepeatingKey(test.input, test.key)

			tres := hex.Encode(res)

			if !bytes.Equal(tres, test.expectedOutput) {
				t.Errorf("expected: %s, got: %s", test.expectedOutput, res)
			}
		})
	}
}
