package aes128

import (
	"crypto/rand"
	"errors"
	"github.com/juggler434/crypto/aes128/oracle"
	"github.com/juggler434/crypto/encoding/base64"
	pkcs7 "github.com/juggler434/crypto/padding"
)

var InvalidReturn = errors.New("invalid data format")

func PaddingAttack(paddingOracle *oracle.CBCOracle) ([]byte, error) {
	iv := make([]byte, 16)
	_, err := rand.Read(iv)
	if err != nil {
		return nil, err
	}

	eb, err := paddingOracle.Encrypt(iv)
	if err != nil {
		return nil, err
	}

	db, err := base64.Decode(eb)
	if err != nil {
		return nil, err
	}

	if len(db) < 17 {
		return nil, InvalidReturn
	}

	blankIV := []byte("0000000000000000")
	et := db[17:]
	res := make([]byte, 16)

	// For each block
	// Make blank previous block
	// For last byte:
	// Loop through all possible values until padding one is reached
	for j := 0; j < 16; j++ {
		index := 15 - j
		for i := 0; i < 256; i++ {
			t := blankIV
			t[index] = byte(i)

			t = append(t, ':')
			t = append(t, et...)

			v := base64.Encode(t)
			_, err := paddingOracle.Decrypt(v)

			if err != nil && err != pkcs7.InvalidPaddingError {
				return res, err
			}

			if err == nil {
				res[index] = byte(i) ^ byte(j+1) ^ iv[index]
				for i, b := range res {
					blankIV[i] = b ^ (byte(j + 2)) ^ iv[i]
				}
				break
			}
		}
	}
	return res, nil
}
