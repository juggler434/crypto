package xor

import (
	"bytes"
	"github.com/juggler434/crypto/encoding/hex"
	"io/ioutil"
	"strings"
	"testing"
)

const TestFile = "./test_files/xor_test.txt"
const TestString = "test case"

func TestDetectSingleCharEncryption(t *testing.T) {
	tests := []struct {
		name           string
		inputFile      string
		expectedOutput []byte
	}{
		{
			name:           "base case",
			inputFile:      TestFile,
			expectedOutput: []byte(TestString),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var input [][]byte
			f, err := ioutil.ReadFile(test.inputFile)
			if err != nil {
				t.Errorf("failed to read input file: %s", err)
			}

			lns := strings.Split(string(f), "\n")
			for _, ln := range lns {
				dhl, err := hex.Decode([]byte(ln))
				if err != nil {
					t.Errorf("failed to decode hex from test file: %s", err)
				}

				input = append(input, dhl)
			}
			res := DetectSingleCharEncryption(input)

			if !bytes.Equal(res, test.expectedOutput) {
				t.Errorf("expected: %s, got: %s", test.expectedOutput, res)
			}
		})
	}
}
