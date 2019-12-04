package cryptopals

import "testing"

func TestFixedXor(t *testing.T) {
	t.Run("with valid inputs", func(t *testing.T) {
		i1 := []byte("1c0111001f010100061a024b53535009181c")
		i2 := []byte("686974207468652062756c6c277320657965")
		expected := "746865206b696420646f6e277420706c6179"

		ret, err := FixedXor(i1, i2)
		if err != nil {
			t.Errorf("expected: error to be nil, got: %s", err)
		}

		if string(ret) != expected {
			t.Errorf("expteced: %s, got: %s", expected, ret)
		}
	})

	t.Run("without valid inputs", func(t *testing.T) {
		i1 := []byte("Non valid input")
		i2 := []byte("686974207468652062756c6c277320657965")

		ret, err := FixedXor(i1, i2)
		if ret != nil {
			t.Errorf("expected: nil, got: %s", ret)
		}

		if err == nil {
			t.Error("expected: an error to be thrown, got: nil")
		}
	})
}
