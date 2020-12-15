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

	ret := make([]byte, len(pi))

	for eb, db := 0, 16; eb < len(pi); eb, db = eb+16, db+16 {
		cipher.Encrypt(ret[eb:db], pi[eb:db])
	}
	return ret, nil
}
