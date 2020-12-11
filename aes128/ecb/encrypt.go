package ecb

import (
	"crypto/aes"
	"github.com/juggler434/crypto/padding"
)

func Encrypt(plaintext, key []byte) ([]byte, error) {
	bs := len(key)
	pi := padding.PKCS7(plaintext, bs)

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ret := make([]byte, 0)

	for eb, db := 0, 16; eb < len(pi); eb, db = eb+len(key), db+len(key) {
		b := make([]byte, bs)
		cipher.Encrypt(b, pi[eb:db])
		ret = append(ret, b...)
	}
	return ret, nil
}
