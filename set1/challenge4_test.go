package cryptopals

import (
	"testing"
)


const TestFile = "./files/xor_test.txt"
const TestString = "test case"
func TestFindXorCipherString(t *testing.T) {
	t.Run("with valid params", func(t *testing.T) {
		res, err := FindXorCipherString(TestFile)
		if err != nil {
			t.Errorf("Expected: error to be nil, got: %s", err)
		}

		if string(res) != TestString {
			t.Errorf("Exptected: %s, got: %s", TestString, res)
		}
	})

	t.Run("with invalid file", func(t *testing.T) {
		res, err := FindXorCipherString("nothing/there.txt")
		if res != nil {
			t.Errorf("Expected: result to be nil, got: %s", res)
		}

		if err == nil {
			t.Error("Expected; an error to be thrown, got: nil")
		}
	})
}