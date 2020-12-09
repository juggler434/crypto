package aes128

import (
	"bytes"
	"testing"
)

func TestBreakECBSimple(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		shouldErr bool
	}{
		{
			name:      "base case",
			input:     []byte("this is a test"),
			shouldErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ret, err := BreakECBSimple(test.input)
			if test.shouldErr {
				if err == nil {
					t.Error("expected: err, got: nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected: nil error, got: %s", err)
				}
			}

			if !bytes.Equal(ret, test.input) {
				t.Errorf("expected: %s, got: %s", test.input, ret)
			}
		})
	}
}
