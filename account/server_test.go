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

func TestUserService_CheckAdminPermission(t *testing.T) {
	tests := []struct {
		name           string
		input          []byte
		expectedOutput bool
		shouldErr      bool
	}{
		{
			name:           "not admin account",
			input:          []byte("email=test@test.com&uid=10&role=user"),
			expectedOutput: false,
			shouldErr:      false,
		}, {
			name:           "admin account",
			input:          []byte("email=test@test.com&uid=10&role=admin"),
			expectedOutput: true,
			shouldErr:      false,
		}, {
			name:           "invalid input",
			input:          []byte("}["),
			expectedOutput: false,
			shouldErr:      true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			us := NewUserService()
			ei, err := ecb.Encrypt(test.input, us.key)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}
			res, err := us.CheckAdminPermission(ei)
			if test.shouldErr {
				if err == nil {
					t.Error("expected: error, got: nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected: nil error, got: %s", err)
				}
			}

			if res != test.expectedOutput {
				t.Errorf("expected: %t, got %t", test.expectedOutput, res)
			}
		})
	}
}
