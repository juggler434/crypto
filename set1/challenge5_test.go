package cryptopals

import "testing"

const RepeatingXorTestString = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"

const RepeatingXorTestResult = "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

func TestEncryptWithRepeatingXor(t *testing.T) {
	t.Run("with valid input", func(t *testing.T) {
		input := []byte(RepeatingXorTestString)
		key := []byte("ICE")
		expected := RepeatingXorTestResult
		res, err := EncryptWithRepeatingXor(input, key)
		if err != nil {
			t.Errorf("Expected: error to be nil, got: %s", err)
		}

		if string(res) != expected {
			t.Errorf("Expected: %s, got: %s", expected, res)
		}
	})
}
