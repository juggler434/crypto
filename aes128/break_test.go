package aes128

import (
	"bytes"
	"github.com/juggler434/crypto/aes128/oracle"
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
			oracle := oracle.NewECBOracle(test.input)
			ret, err := BreakECBSimple(oracle)
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

//func TestBreakECBAdvanced(t *testing.T) {
//	tests := []struct {
//		name      string
//		input     []byte
//		shouldErr bool
//	}{
//		{
//			name:      "base case",
//			input:     []byte("this is a test"),
//			shouldErr: false,
//		},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			o := oracle.NewAdvancedECBOracle(test.input)
//			ret, err := BreakECBAdvanced(o)
//			if test.shouldErr {
//				if err == nil {
//					t.Error("expected: err, got: nil")
//				}
//			} else {
//				if err != nil {
//					t.Errorf("expected: nil error, got: %s", err)
//				}
//			}
//
//			if !bytes.Equal(ret, test.input) {
//				t.Errorf("expected: %s, got: %s", test.input, ret)
//			}
//		})
//	}
//}
