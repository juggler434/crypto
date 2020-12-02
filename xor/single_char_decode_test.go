package xor

import (
	"bytes"
	"github.com/juggler434/crypto/encoding/hex"
	"testing"
)

func TestSingleCharDecode(t *testing.T) {

	tests := []struct {
		name           string
		hexInput       []byte
		expectedOutput []byte
		expectedScore  int
		key            byte
		shouldError    bool
	}{
		{
			name:           "base case",
			hexInput:       []byte("2c45160d0a1009014502001145110d0c16"),
			expectedOutput: []byte("I should get this"),
			expectedScore:  130,
			key:            byte('e'),
			shouldError:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			eb, err := hex.Decode(test.hexInput)
			if err != nil {
				t.Errorf("unexpected error decoding hex test case: %s", err)
			}
			res, score := SingleCharDecode(eb)

			if !bytes.Equal(res, test.expectedOutput) {
				t.Errorf("expected: %s, got: %s", test.expectedOutput, res)
			}

			if score != test.expectedScore {
				t.Errorf("expected score: %d, got: %d", test.expectedScore, score)
			}
		})
	}
}
