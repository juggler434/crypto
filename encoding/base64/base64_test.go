package base64

import (
	"bytes"
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		input []byte
		expectedOutput []byte
	}{
		{
			name: "normal case",
			input: []byte("Hello, World"),
			expectedOutput: []byte("SGVsbG8sIFdvcmxk"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ret := Encode(test.input)
			if !bytes.Equal(ret, test.expectedOutput) {
				t.Errorf("expected %s, got %s", test.expectedOutput, ret)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name string
		input []byte
		expectedOutput []byte
		shouldError bool
	}{
		{
			name: "normal case",
			input: []byte("SGVsbG8sIFdvcmxk"),
			expectedOutput: []byte("Hello, World"),
			shouldError: false,
		},
		{
			name: "invalid input",
			input: []byte("ae432kl;ao3"),
			expectedOutput: nil,
			shouldError: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ret, err := Decode(test.input)
			if test.shouldError {
				if err == nil {
					t.Error("expected: error, got: nil")
				}
			}else {
				if err != nil {
					t.Errorf("expected: nil error, got: %s", err)
				}
			}

			if !bytes.Equal(ret, test.expectedOutput) {
				t.Errorf("expected: %s, got %s", test.expectedOutput, ret)
			}

		})
	}
}

