package cbc

import (
	"crypto/aes"
	pkcs7 "github.com/juggler434/crypto/padding"
	"github.com/juggler434/crypto/xor"
	"math"
)

func Decrypt(input, key, initializationVector []byte) ([]byte, error) {
	lb := initializationVector
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	res := make([]byte, 0)

	for eb, db := 0, 16; eb < len(input); eb, db = eb+16, db+16 {
		b := make([]byte, 16)
		ev := math.Min(float64(len(input)), float64(db))
		cipher.Decrypt(b, input[eb:int(ev)])

		b, err := xor.Fixed(b, lb)
		if err != nil {
			return nil, err
		}

		res = append(res, b...)
		lb = input[eb:db]
	}

	res, err = pkcs7.Unpad(res)

	return res, err
}
