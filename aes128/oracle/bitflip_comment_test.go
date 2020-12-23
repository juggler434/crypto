package oracle

import (
	"bytes"
	"github.com/juggler434/crypto/aes128/cbc"
	"testing"
)

func TestCommentEncrypter_Encrypt(t *testing.T) {
	tests := []struct {
		name           string
		input          []byte
		expectedOutput []byte // Not sure if this is the best way to test this, but works for now!
		shouldErr      bool
	}{
		{
			name:           "base case",
			input:          []byte("Hello, World!!!"),
			expectedOutput: []byte("Hello, World!!!"),
			shouldErr:      false,
		}, {
			name:           "should quote out = in input",
			input:          []byte("admin=true"),
			expectedOutput: []byte("admin\"=\"true"),
			shouldErr:      false,
		}, {
			name:           "should quoate out ; in input",
			input:          []byte("endingtheblock;"),
			expectedOutput: []byte("endingtheblock\";\""),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ce := NewCommentEncrypter()
			res, err := ce.Encrypt(test.input)
			if test.shouldErr {
				if err == nil {
					t.Error("expected: error, got: nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected: nil err, got: %s", err)
				}
			}

			res, err = cbc.Decrypt(res, ce.key, ce.initializationVector)
			if err != nil {
				t.Errorf("unexepcted error: %s", err)
			}

			expectedOutput := append([]byte(commentPrefix), test.expectedOutput...)
			expectedOutput = append(expectedOutput, []byte(commentPostfix)...)

			if !bytes.Equal(res, expectedOutput) {
				t.Errorf("expected: %s, got: %s", expectedOutput, res)
			}
		})
	}
}

func TestCommentEncrypter_CheckIsAdmin(t *testing.T) {
	tests := []struct {
		name            string
		input           []byte
		expectedOutcome bool
		shouldErr       bool
	}{
		{
			name:            "base case, is admin",
			input:           []byte("something;admin=true;topic"),
			expectedOutcome: true,
			shouldErr:       false,
		}, {
			name:            "base case, is not admin",
			input:           []byte("totally random input"),
			expectedOutcome: false,
			shouldErr:       false,
		}, {
			name:            "sanitized input",
			input:           []byte("something\";\"admin\"=\"true\";\"topic"),
			expectedOutcome: false,
			shouldErr:       false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := NewCommentEncrypter()

			input := append([]byte(commentPrefix), test.input...)
			input = append(input, []byte(commentPostfix)...)

			//we can't use our encrypter to encrypt, because it will comment out the = and ; characters, which we don't want
			ei, err := cbc.Encrypt(input, e.key, e.initializationVector)
			if err != nil {
				t.Errorf("unexepected error: %s", err)
			}

			res, err := e.CheckIsAdmin(ei)
			if test.shouldErr {
				if err == nil {
					t.Error("expected: error, got : nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected: nil error, got: %s", err)
				}
			}
			if res != test.expectedOutcome {
				t.Errorf("expected: %t, got %t", test.expectedOutcome, res)
			}
		})
	}
}
