package account

import (
	"bytes"
	"testing"
)

func TestCreateCookie(t *testing.T) {
	tests := []struct {
		name           string
		input          *Account
		expectedOutput []byte
	}{
		{
			name: "base case",
			input: &Account{
				email: []byte("test@test.com"),
				uid:   []byte("1"),
				role:  []byte("user"),
			},
			expectedOutput: []byte("email=test@test.com&uid=1&role=user"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := CreateCookie(test.input)
			if !bytes.Equal(res, test.expectedOutput) {
				t.Errorf("expected: %s, got %s", test.expectedOutput, res)
			}
		})
	}
}
