package oracle

import (
	"testing"
)

func TestRandom_Encrypt(t *testing.T) {
	tests := []struct {
		name        string
		input       []byte
		shouldError bool
	}{
		{
			name:        "base case",
			input:       []byte("12345678901234561234567890123456"),
			shouldError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			o := NewRandom()
			res, et, err := o.Encrypt(test.input)
			if test.shouldError {
				if err == nil {
					t.Error("expected: error, got: nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected: nil error, got: %s", err)
				}
			}

			if et != ECB && et != CBC {
				t.Errorf("expected: 0 or 1, got: %d", et)
			}

			if len(res) <= len(test.input) {
				t.Errorf("expected: encrypted length greater than %d, got: %d", len(test.input), len(test.input))
			}

		})
	}
}
