package pkcs7

import (
	"bytes"
	"testing"
)

func TestPad(t *testing.T) {
	tests := []struct {
		name           string
		input          []byte
		blockSize      int
		expectedOutput []byte
	}{
		{
			name:           "block size greater than input",
			input:          []byte("YELLOW SUBMARINE"),
			blockSize:      20,
			expectedOutput: []byte("YELLOW SUBMARINE\x04\x04\x04\x04"),
		}, {
			name:           "block size smaller than input",
			input:          []byte("YELLOW SUBMARINE"),
			blockSize:      6,
			expectedOutput: []byte("YELLOW SUBMARINE\x02\x02"),
		}, {
			name:           "input multiple of block size",
			input:          []byte("YELLOW SUBMARINE"),
			blockSize:      4,
			expectedOutput: []byte("YELLOW SUBMARINE\x04\x04\x04\x04"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := Pad(test.input, test.blockSize)
			if !bytes.Equal(res, test.expectedOutput) {
				t.Errorf("expected: %+q, got: %+q", test.expectedOutput, res)
			}
		})
	}
}

func TestUnpad(t *testing.T) {
	tests := []struct {
		name           string
		input          []byte
		expectedOutput []byte
		shouldErr      bool
	}{
		{
			name:           "base case",
			input:          []byte("YELLOW SUBMARINE\x04\x04\x04\x04"),
			expectedOutput: []byte("YELLOW SUBMARINE"),
			shouldErr:      false,
		}, {
			name:           "too few pad bytes",
			input:          []byte("YELLOW SUBMARINE\x04\x04\x04"),
			expectedOutput: nil,
			shouldErr:      true,
		}, {
			name:           "unmatching pad bytes",
			input:          []byte("YELLOW SUBMARINE\x01\x02\x03\x04"),
			expectedOutput: nil,
			shouldErr:      true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := Unpad(test.input)
			if test.shouldErr {
				if err == nil {
					t.Error("expected: error, got: nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected: nil error, got: %s", err)
				}
			}

			if !bytes.Equal(res, test.expectedOutput) {
				t.Errorf("expected: %s, got %s", test.expectedOutput, res)
			}
		})
	}
}
