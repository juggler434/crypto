package ecb

import (
	"bytes"
	"crypto/aes"
)

func Decrypt(cypherText, key []byte) ([]byte, error) {
	plainText, err := decryptEcb(cypherText, key)
	if err != nil {
		return nil, err
	}

	plainText = bytes.Trim(plainText, "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f")
	return plainText, nil
}

func decryptEcb(ueb []byte, key []byte) ([]byte, error) {
	ueb = bytes.Trim(ueb, "\x00")

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plainText := make([]byte, len(ueb))
	for eb, db := 0, 16; eb < len(ueb); eb, db = eb+16, db+16 {
		cipher.Decrypt(plainText[eb:db], ueb[eb:db])
	}
	return plainText, nil
}
