package aes128

import (
	"bytes"
	"github.com/juggler434/crypto/aes128/oracle"
	"github.com/juggler434/crypto/encoding/base64"
	"testing"
)

func TestPaddingAttack(t *testing.T) {
	tests := []struct {
		name           string
		expectedOutput []byte
	}{
		{
			name:           "base case",
			expectedOutput: []byte("this is a test"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input := base64.Encode(test.expectedOutput)

			o, err := oracle.NewCBCOracle([][]byte{input})
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}

			v := o.(*oracle.CBCOracle)

			res, err := PaddingAttack(v)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}

			if !bytes.Equal(res, test.expectedOutput) {
				t.Errorf("expected: %s, got: %s", test.expectedOutput, res)
			}
		})
	}
}
