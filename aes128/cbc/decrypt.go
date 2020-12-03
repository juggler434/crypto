package cbc

import (
	"crypto/aes"
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

		for _, by := range b {
			if by > byte(16) { // 0 - 15 will be padding characters
				res = append(res, by)
			}
		}
		lb = input[eb:db]
	}

	return res, nil
}
