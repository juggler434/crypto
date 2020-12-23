package aes128

import (
	"github.com/juggler434/crypto/aes128/oracle"
	"testing"
)

func TestBitflipAttack(t *testing.T) {
	tests := []struct {
		name            string
		expectedOutcome bool
		shouldErr       bool
	}{
		{
			name:            "base case",
			expectedOutcome: true,
			shouldErr:       false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := oracle.NewCommentEncrypter()
			r, err := BitflipAttack(e)
			if test.shouldErr {
				if err == nil {
					t.Error("expected: err, got: nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected: nil error, got: %s", err)
				}
			}

			if r != test.expectedOutcome {
				t.Errorf("expected: %t, got %t", test.expectedOutcome, r)
			}
		})
	}
}
