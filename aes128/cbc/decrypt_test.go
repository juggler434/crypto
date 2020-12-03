package cbc

import (
	"bytes"
	"github.com/juggler434/crypto/encoding/base64"
	"testing"
)

func TestDecrypt(t *testing.T) {
	tests := []struct {
		name                 string
		input                []byte
		key                  []byte
		initializationVector []byte
		expectedOutput       []byte
		shouldError          bool
	}{
		{
			name:                 "base case",
			input:                []byte("IO0OI/27Zw9Q5MrfZZlw2F5StNTl7/hz+5mfE5BoijJrwu22mtXetJ/F2kQYFidQ"),
			key:                  []byte("YELLOW SUBMARINE"),
			initializationVector: []byte("0000000000000000"),
			expectedOutput:       []byte("This is a test, I repeat, this is a test"),
			shouldError:          false,
		},
		{
			name:                 "wrong key size",
			key:                  []byte("YELLOW SUBMARIN"),
			input:                []byte("IO0OI/27Zw9Q5MrfZZlw2F5StNTl7/hz+5mfE5BoijJrwu22mtXetJ/F2kQYFidQ"),
			initializationVector: []byte("0000000000000000"),
			expectedOutput:       nil,
			shouldError:          true,
		}, {
			name:                 "initialization vector doesn't match block size",
			key:                  []byte("YELLOW SUBMARINE"),
			input:                []byte("IO0OI/27Zw9Q5MrfZZlw2F5StNTl7/hz+5mfE5BoijJrwu22mtXetJ/F2kQYFidQ"),
			initializationVector: []byte("00000000000000"),
			expectedOutput:       nil,
			shouldError:          true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			hi, err := base64.Decode(test.input)
			if err != nil {
				t.Errorf("failed to decode base64 input for test: %s", err)
			}
			res, err := Decrypt(hi, test.key, test.initializationVector)
			if test.shouldError {
				if err == nil {
					t.Error("expected: error, got: nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected: nil error, got: %s", err)
				}
			}

			if !bytes.Equal(test.expectedOutput, res) {
				t.Errorf("expected: %s, got: %s", test.expectedOutput, res)
			}

		})
	}
}
