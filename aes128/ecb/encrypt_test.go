package ecb

import (
	"bytes"
	"github.com/juggler434/crypto/encoding/base64"
	"testing"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name           string
		input          []byte
		key            []byte
		expectedOutput []byte
		shouldErr      bool
	}{
		{
			name:           "base case",
			input:          []byte("THIS IS A TEST!!"),
			key:            []byte("YELLOW SUBMARINE"),
			expectedOutput: []byte("dgdvTfIWW2IBN+w+g8xLcw=="),
			shouldErr:      false,
		},
		{
			name:           "pads shorter input",
			input:          []byte("THIS IS A TEST"),
			key:            []byte("YELLOW SUBMARINE"),
			expectedOutput: []byte("Q6bquZeGNENhBmZSdXREIQ=="),
			shouldErr:      false,
		},
		{
			name:           "improper key length",
			input:          []byte("THIS IS A TEST!!"),
			key:            []byte("YELLOW SUBMARIN"),
			expectedOutput: nil,
			shouldErr:      true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := Encrypt(test.input, test.key)
			if test.shouldErr {
				if err == nil {
					t.Error("expected: error, got: nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected: nil error, got: %s", err)
				}
			}

			result := base64.Encode(res)
			if !bytes.Equal(result, test.expectedOutput) {
				t.Errorf("expected: %s, got %s", test.expectedOutput, result)
			}
		})
	}
}
