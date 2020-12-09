package oracle

import (
	"crypto/aes"
	"github.com/juggler434/crypto/aes128/cbc"
	"math/rand"
	"time"
)

const (
	ECB = iota
	CBC
)

type Random struct {
	key []byte
}

func NewRandom() *Random {
	key, err := GenerateKey()
	if err != nil {
		panic(err)
	}
	return &Random{key: key}
}

// RandomEncrypt will generate a random key, prepend and append random bytes to the text, and encrypt under either ecb or cbc
func (o *Random) Encrypt(plainText []byte) ([]byte, int, error) {
	rand.Seed(time.Now().UnixNano())
	ppLen := rand.Intn(6) + 5
	apLen := rand.Intn(6) + 5

	pp := make([]byte, ppLen)
	rand.Read(pp)

	ft := append(pp, plainText...)

	ap := make([]byte, apLen)
	ft = append(ft, ap...)

	et := rand.Intn(2)
	switch et {
	case ECB:
		//TODO Make an encrypt function for this
		cipher, err := aes.NewCipher(o.key)
		if err != nil {
			return nil, -1, err
		}

		ret := make([]byte, len(ft))
		cipher.Encrypt(ret, ft)
		return ft, ECB, nil
	case CBC:
		iv, err := GenerateKey() //Initialization Vector is pretty similar to a key
		if err != nil {
			return nil, -1, err
		}
		et, err := cbc.Encrypt(ft, o.key, iv)
		if err != nil {
			return nil, -1, err
		}
		return et, CBC, nil
	}

	return nil, -1, nil

}
