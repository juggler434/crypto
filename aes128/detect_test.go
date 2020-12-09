package aes128

import (
	"github.com/juggler434/crypto/encoding/base64"
	"testing"
)

func TestDetectMode(t *testing.T) {
	tests := []struct {
		name           string
		input          []byte
		encryptionType int
		errorExpected  bool
	}{
		{
			name:           "base case encrypted with ECB",
			input:          []byte("WBLjgxlE+JXs3uh7I86B71gS44MZRPiV7N7oeyPOge9g+jZwfkX0mdug8luSIwGl"),
			encryptionType: ECB,
		}, {
			name:           "base case encrypted with CBC",
			input:          []byte("/HeASZHem/P1e9K5LLoZKymlEDJl2Im8i4HxRhrA37nKMjoysgM3cuBnspXgBnpS"),
			encryptionType: CBC,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			decodedInput, err := base64.Decode(test.input)
			if err != nil {
				t.Errorf("failed to decode test input: %s", err)
			}

			res := DetectMode(decodedInput)
			if res != test.encryptionType {
				t.Errorf("expected: %d, got: %d", test.encryptionType, res)
			}
		})
	}
}
