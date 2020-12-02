package ecb

import (
	"bufio"
	"bytes"
	"github.com/juggler434/crypto/encoding/hex"
	"io/ioutil"
	"testing"
)

const (
	ECBEncryptedLine        = 2
	ValidTestFile           = "./test_files/ecb_detection_test.txt"
	ECBEncryptedLineReverse = 1
	ValidReverseTestFile    = "./test_files/ecb_detection_reverse_test.txt"
)

func TestDetect(t *testing.T) {
	tests := []struct {
		name           string
		inputFile      string
		expectedOutput int
	}{
		{
			name:           "base case",
			inputFile:      ValidTestFile,
			expectedOutput: ECBEncryptedLine,
		},
		{
			name:           "reverse lines for min test",
			inputFile:      ValidReverseTestFile,
			expectedOutput: ECBEncryptedLineReverse,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var input [][]byte
			f, err := ioutil.ReadFile(test.inputFile)
			if err != nil {
				t.Errorf("failed to read test file: %s", err)
			}

			scanner := bufio.NewScanner(bytes.NewReader(f))
			for scanner.Scan() {
				hdl, err := hex.Decode(scanner.Bytes())
				if err != nil {
					t.Errorf("failed to decode hex bytes: %s", err)
				}

				input = append(input, hdl)
			}
			res := Detect(input)

			if res != test.expectedOutput {
				t.Errorf("Expected %d, got %d", test.expectedOutput, res)
			}
		})

	}
}
