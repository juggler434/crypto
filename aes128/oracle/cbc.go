package oracle

import (
	"errors"
	"github.com/juggler434/crypto/aes128/cbc"
	"github.com/juggler434/crypto/encoding/base64"
	"math/rand"
	"time"
)

type CBCOracle struct {
	Plaintext [][]byte
	Key       []byte
}

func NewCBCOracle(plaintext [][]byte) (Encrypter, error) {
	key, err := GenerateKey()
	return &CBCOracle{Plaintext: plaintext, Key: key}, err
}

func (c *CBCOracle) Encrypt(initializationVector []byte) ([]byte, error) {
	rand.Seed(time.Now().UnixNano())

	uel := c.Plaintext[rand.Intn(len(c.Plaintext))]

	eel, err := base64.Decode(uel)
	if err != nil {
		return nil, err
	}

	r, err := cbc.Encrypt(eel, c.Key, initializationVector)
	if err != nil {
		return nil, err
	}

	ret := append(initializationVector, ':')
	ret = append(ret, r...)

	res := base64.Encode(ret)

	return res, nil
}

func (c *CBCOracle) Decrypt(encryptedText []byte) ([]byte, error) {
	et, err := base64.Decode(encryptedText)
	if err != nil {
		return nil, err
	}

	if len(et) < 17 {
		return nil, errors.New("input too short")
	}

	if et[16] != ':' {
		return nil, errors.New("malformed initialization vector")
	}

	return cbc.Decrypt(et[17:], c.Key, et[:16])
}
