package hex

import (
	"bytes"
	"testing"
)

func TestEncode(t *testing.T) {
	testCases :=  []struct {
		description    string
		input          []byte
		expectedOutput []byte
	} {
		{
			description: "normal case",
			input: []byte("Hello, World"),
			expectedOutput: []byte("48656c6c6f2c20576f726c64"),
		},
	}

	for _, test := range testCases {
		t.Run(test.description, func(t *testing.T) {
			r := Encode(test.input)
			if !bytes.Equal(test.expectedOutput, r) {
				t.Errorf("expected: %s, got %s", test.expectedOutput, r)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	testCases := []struct {
		description string
		input []byte
		expectedOutput []byte
		shouldError bool
	} {
		{
			description: "passing case",
			input: []byte("48656c6c6f2c20576f726c64"),
			expectedOutput: []byte("Hello, World"),
			shouldError: false,
		},
		{
			description: "invalid hex input",
			input: []byte("This ins't hex"),
			expectedOutput: nil,
			shouldError: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.description, func(t *testing.T) {
			r, err := Decode(test.input)
			if test.shouldError {
				if err == nil {
					t.Error("expected an error, got: nil")
				}
			}

			if err != nil {
				t.Errorf("expected nil error, got: %s", err)
			}

			if !bytes.Equal(test.expectedOutput, r) {
				t.Errorf("expected: %s, got %s", test.expectedOutput, r)
			}
		})
	}

}
