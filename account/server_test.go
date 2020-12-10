package account

import (
	"bytes"
	"github.com/juggler434/crypto/aes128/ecb"
	"testing"
)

func TestNewUserService(t *testing.T) {
	us := NewUserService()
	if us.key == nil || len(us.key) != 16 {
		t.Error("failed to initiate user service with correct key")
	}
}

func TestUserService_GetUser(t *testing.T) {
	tests := []struct {
		name           string
		input          []byte
		expectedOutput []byte
		shouldErr      bool
	}{
		{
			name:           "base case",
			input:          []byte("john@example.com"),
			expectedOutput: []byte("email=john@example.com&uid=10&role=user"),
			shouldErr:      false,
		}, {
			name:           "invalid input",
			input:          []byte("joh@example.com&role=admin"),
			expectedOutput: nil,
			shouldErr:      true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			us := NewUserService()
			o, err := us.GetUser(test.input)
			if test.shouldErr {
				if err == nil {
					t.Error("expected: err, got: nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected: nil error, got:%s", err)
				}
			}

			res, err := ecb.Decrypt(o, us.key)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}

			if !bytes.Equal(res, test.expectedOutput) {
				t.Errorf("expected: %s, got:%s", test.expectedOutput, res)
			}
		})
	}
}
