package cryptopals

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"io/ioutil"
)

func DecryptAES128Ecb(file string, key []byte)([]byte, error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	ueb := make([]byte, base64.StdEncoding.DecodedLen(len(f)))
	_, err = base64.StdEncoding.Decode(ueb, f)
	if err != nil {
		return nil, err
	}

	plainText, err := decryptEcb(ueb, key)
	if err != nil {
		return nil, err
	}

	plainText = bytes.Trim(plainText, "\x04")
	return plainText, nil
}

func decryptEcb(ueb []byte,  key []byte) ([]byte,  error) {
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