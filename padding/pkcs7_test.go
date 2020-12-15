package padding

import (
	"bytes"
	"testing"
)

func TestPKCS7(t *testing.T) {
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
			res := PKCS7(test.input, test.blockSize)
			if !bytes.Equal(res, test.expectedOutput) {
				t.Errorf("expected: %+q, got: %+q", test.expectedOutput, res)
			}
		})
	}
}
