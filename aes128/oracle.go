package aes128

import "github.com/juggler434/crypto/aes128/ecb"

const (
	ECB = iota
	CBC
)

// Oracle takes in encrypted text, and returns whether it is ECB or CBC encoded
func Oracle(encryptedText []byte) int {
	rl := ecb.FindRepeats(encryptedText)
	if rl > 0 {
		return ECB
	}
	return CBC
}
