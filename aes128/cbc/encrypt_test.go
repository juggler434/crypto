package cbc

import (
	"bytes"
	"github.com/juggler434/crypto/encoding/base64"
	"testing"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name                 string
		input                []byte
		key                  []byte
		initializationVector []byte
		expectedOutputBase64 []byte
		shouldError          bool
	}{
		{
			name:                 "base case",
			input:                []byte("This is a test, I repeat, this is a test"),
			key:                  []byte("YELLOW SUBMARINE"),
			initializationVector: []byte("0000000000000000"),
			expectedOutputBase64: []byte("IO0OI/27Zw9Q5MrfZZlw2F5StNTl7/hz+5mfE5BoijJrwu22mtXetJ/F2kQYFidQ"),
			shouldError:          false,
		}, {
			name:                 "incorrect key length",
			input:                []byte("This is a test, I repeat, this is a test"),
			key:                  []byte("YELLOW SUBMARIN"),
			initializationVector: []byte("0000000000000000"),
			expectedOutputBase64: nil,
			shouldError:          true,
		}, {
			name:                 "incorrect initialization vector length",
			input:                []byte("This is a test, I repeat, this is a test"),
			key:                  []byte("YELLOW SUBMARINE"),
			initializationVector: []byte("00000000000000"),
			expectedOutputBase64: nil,
			shouldError:          true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := Encrypt(test.input, test.key, test.initializationVector)
			if test.shouldError {
				if err == nil {
					t.Error("expected: error, got: nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected: nil error, got: %s", err)
				}
			}

			b64res := base64.Encode(res)
			if !bytes.Equal(b64res, test.expectedOutputBase64) {
				t.Errorf("expected: %s, got: %s", test.expectedOutputBase64, b64res)
			}
		})
	}
}
