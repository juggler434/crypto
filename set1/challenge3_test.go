package cryptopals

import (
	"encoding/hex"
	"testing"
)

func TestSingleXorCipher(t *testing.T) {

	t.Run("with valid input", func(t *testing.T) {
		ts := []byte("I should get this")
		k := byte('e')

		eb := make([]byte, len(ts))
		for i := 0; i < len(ts); i++ {
			eb[i] = ts[i] ^ k
		}

		hb := make([]byte, hex.EncodedLen(len(eb)))
		hex.Encode(hb, eb)

		res, err := SingleXorCipher(hb)
		if err != nil {
			t.Errorf("Expected: error to be nil, got: %s", err)
		}

		if string(res) != string(ts) {
			t.Errorf("Expected: %s, got: %s", ts, res)
		}
	})

	t.Run("with invalid input", func(t *testing.T) {
		ts := []byte("Not a valid input")

		res, err := SingleXorCipher(ts)
		if res != nil {
			t.Errorf("Expected nil, got: %s", res)
		}

		if err == nil {
			t.Error("Expected: an error, got: nil")
		}
	})
}
