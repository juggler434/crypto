package account

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name           string
		input          []byte
		expectedOutput Account
		shouldErr      bool
	}{
		{
			name:  "base case",
			input: []byte("email=foo@bar.com&uid=10&role=user"),
			expectedOutput: Account{
				email: []byte("foo@bar.com"),
				uid:   []byte("10"),
				role:  []byte("user"),
			},
			shouldErr: false,
		}, {
			name:           "invalid input",
			input:          []byte(")["),
			expectedOutput: Account{},
			shouldErr:      true,
		}, {
			name:           "unkown field",
			input:          []byte("e-mail=foo@bar.com&uid=10&role=user"),
			expectedOutput: Account{},
			shouldErr:      true,
		}, {
			name:           "multiple key/values per entry",
			input:          []byte("email=foo@bar.com=penny-acrade.com&uid=10&role=user"),
			expectedOutput: Account{},
			shouldErr:      true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := Parse(test.input)
			if test.shouldErr {
				if err == nil {
					t.Error("expected: error, got: nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected: nil error, got: %s", err)
				}
				if reflect.DeepEqual(res, test.expectedOutput) {
					t.Errorf("expected: %x, got %x", test.expectedOutput, *res)
				}
			}
		})
	}
}
