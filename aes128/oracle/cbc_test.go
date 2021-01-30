package oracle

import (
	"bytes"
	"crypto/rand"
	"github.com/juggler434/crypto/aes128/cbc"
	"github.com/juggler434/crypto/encoding/base64"
	"io/ioutil"
	"testing"
)

const cbcTestFile = "./test_files/cbc_oracle_test.txt"
const cbcBadTestFile = "./test_files/cbc_bad_data.txt"

func readTestFile(t *testing.T, inputFile string) [][]byte {
	f, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	l := bytes.Split(f, []byte("\n"))
	return l
}

func TestNewCBCOracle(t *testing.T) {
	l := readTestFile(t, cbcTestFile)

	r, err := NewCBCOracle(l)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	v := r.(*CBCOracle)
	if v.Key == nil {
		t.Error("expected a key to be initialized, got nil")
	}

	if v.Plaintext == nil || len(v.Plaintext) == 0 {
		t.Error("expected plaintext lines, got nil")
	}
}

func TestCBCOracle_Encrypt(t *testing.T) {
	tests := []struct {
		name      string
		inputFile string
		shouldErr bool
	}{
		{
			name:      "base case",
			inputFile: cbcTestFile,
			shouldErr: false,
		}, {
			name:      "malformed base64",
			inputFile: cbcBadTestFile,
			shouldErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := readTestFile(t, test.inputFile)
			r, err := NewCBCOracle(l)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}

			o := r.(*CBCOracle) // Do this so we can get the key to decrypt test data

			iv := make([]byte, 16)
			_, err = rand.Read(iv)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}

			res, err := r.Encrypt(iv)

			if test.shouldErr {
				if err == nil {
					t.Error("expected: err, got: nil")
				}
				return
			} else {
				if err != nil {
					t.Errorf("expected: nil error, got: %s", err)
				}
			}

			v, err := base64.Decode(res)
			if err != nil {
				t.Errorf("unexptected error: %s", err)
			}

			riv := v[:16]
			if !bytes.Equal(iv, riv) {
				t.Errorf("expected first 16 bytes to equal: %s, got %s", iv, riv)
			}

			e := v[17:]

			v, err = cbc.Decrypt(e, o.Key, iv)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}

			if !comparePlainText(v) {
				t.Errorf("expected decrypted value to match, one of tests values, got: %s", v)
			}
		})
	}
}

func TestCBCOracle_Decrypt(t *testing.T) {
	tests := []struct {
		name           string
		input          []byte
		expectedOutput []byte
		shouldErr      bool
	}{
		{
			name:           "base case",
			input:          []byte("MDAwMDAwMDAwMDAwMDAwMDol3JBQKbFhLvYynuFYy13Q"),
			expectedOutput: []byte("This is a test"),
			shouldErr:      false,
		}, {
			name:           "input too short",
			input:          []byte("c2hvcnQ="),
			expectedOutput: nil,
			shouldErr:      true,
		}, {
			name:           "invalid base64",
			input:          []byte("this isn't valid base 64 so should error"),
			expectedOutput: nil,
			shouldErr:      true,
		}, {
			name:           "malformed initialization vector",
			input:          []byte("JdyQUCmxYS72Mp7hWMtd0A=="),
			expectedOutput: nil,
			shouldErr:      true,
		}, {
			name:           "invalid padding",
			input:          []byte("MDAwMDAwMDAwMDAwMDAwMDr7kUiKh+eBOC/nkqFuSUOg"),
			expectedOutput: nil,
			shouldErr:      true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			o, err := NewCBCOracle([][]byte{})
			if err != nil {
				t.Errorf("unexpected err: %s", err)
			}

			co := o.(*CBCOracle)
			co.Key = []byte("YELLOW SUBMARINE") // set key so that we know what we are encrypting under

			res, err := co.Decrypt(test.input)
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

func comparePlainText(plaintext []byte) bool {
	testInputs := [][]byte{[]byte("This line is really long and is more than one block"),
		[]byte("Tacos are cool"),
		[]byte("Third input"),
	}

	for _, input := range testInputs {
		if bytes.Equal(input, plaintext) {
			return true
		}
	}
	return false
}
