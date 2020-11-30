package hex

import (
	"bytes"
	"testing"
)

func TestHexToBase64(t *testing.T) {

	tests := []struct {
		name           string
		input          []byte
		expectedOutput []byte
		shouldError    bool
	}{
		{
			name:           "with valid input",
			input:          []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"),
			expectedOutput: []byte("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"),
			shouldError:    false,
		},
		{
			name:           "with non-valid hex string",
			input:          []byte("This is not a valid hex string"),
			expectedOutput: nil,
			shouldError:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ret, err := ToBase64(test.input)
			if test.shouldError {
				if err == nil {
					t.Error("expected: error, got: nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected: nil error, got: %s", err)
				}
			}

			if !bytes.Equal(ret, test.expectedOutput) {
				t.Errorf("expected: %s, got: %s", test.expectedOutput, ret)
			}
		})
	}
}
