package ecb

import (
	"bytes"
	"crypto/aes"
	pkcs7 "github.com/juggler434/crypto/padding"
)

func Decrypt(cypherText, key []byte) ([]byte, error) {
	plainText, err := decryptEcb(cypherText, key)
	if err != nil {
		return nil, err
	}

	plainText, err = pkcs7.Unpad(plainText)
	if err != nil {
		return nil, err
	}
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
