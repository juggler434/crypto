package cryptopals

import "testing"

const (
	ECBEncryptedLine = 2
)
func TestDetectECBEncryption (t *testing.T) {
	res, err := DetectECBEncryption("./files/ecb_detection_test.txt")

	if err != nil {
		t.Errorf("Expected nil error, got %s", err)
	}

	if res != ECBEncryptedLine {
		t.Errorf("Expected %d, got %d", ECBEncryptedLine, res)
	}
}
