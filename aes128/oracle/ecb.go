package oracle

import (
	"github.com/juggler434/crypto/aes128/ecb"
)

type ECBOracle struct {
	key []byte
}

func NewECBOracle() Encrypter {
	key, _ := GenerateKey()
	return &ECBOracle{key: key}
}

func (o *ECBOracle) Encrypt(userInput, secretStuff []byte) ([]byte, error) {
	return ecb.Encrypt(append(userInput, secretStuff...), o.key)

}
