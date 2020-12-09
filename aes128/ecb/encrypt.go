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
	cipher.Encrypt(ret, pi)
	return ret, nil
}
