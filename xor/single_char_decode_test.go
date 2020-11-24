package xor

import (
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

		res, _ := SingleCharDecode(eb)

		if string(res) != string(ts) {
			t.Errorf("Expected: %s, got: %s", ts, res)
		}
	})
}
