package cbc

import (
	"crypto/aes"
	"github.com/juggler434/crypto/padding"
	"github.com/juggler434/crypto/xor"
)

func Encrypt(plaintext, key, initializationVector []byte) ([]byte, error) {
	lb := initializationVector
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	res := make([]byte, 0)
	pi := pkcs7.Pad(plaintext, len(key))

	for eb, db := 0, 16; eb < len(pi); eb, db = eb+len(key), db+len(key) {
		b, err := xor.Fixed(lb, pi[eb:db])
		if err != nil {
			return nil, err
		}

		cipher.Encrypt(b, b)
		res = append(res, b...)
		lb = b
	}

	return res, nil
}
