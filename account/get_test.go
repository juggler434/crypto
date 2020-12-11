package account

import (
	"bytes"
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name           string
		input          []byte
		expectedOutput []byte
		shouldError    bool
	}{
		{
			name:           "base case",
			input:          []byte("joe@example.com"),
			expectedOutput: []byte("email=joe@example.com&uid=10&role=user"),
			shouldError:    false,
		}, {
			name:           "& character in email",
			input:          []byte("joe@example.com&admin"),
			expectedOutput: nil,
			shouldError:    true,
		}, {
			name:           "= character in email",
			input:          []byte("joe@example.com=admin"),
			expectedOutput: nil,
			shouldError:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := Get(test.input)
			if test.shouldError {
				if err == nil {
					t.Error("expected: err, got: nil err")
				}
			} else {
				if err != nil {
					t.Errorf("expected: nil error, got: %s", err)
				}
			}

			if !bytes.Equal(res, test.expectedOutput) {
				t.Errorf("expected: %s, got: %s", test.expectedOutput, res)
			}
		})
	}
}
