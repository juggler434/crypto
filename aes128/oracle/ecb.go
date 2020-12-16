package oracle

import (
	"github.com/juggler434/crypto/aes128/ecb"
	"math/rand"
	"time"
)

type ECBOracle struct {
	key         []byte
	secretStuff []byte
}

func NewECBOracle(secretStuff []byte) Encrypter {
	key, _ := GenerateKey()
	return &ECBOracle{key: key, secretStuff: secretStuff}
}

func (o *ECBOracle) Encrypt(userInput []byte) ([]byte, error) {
	return ecb.Encrypt(append(userInput, o.secretStuff...), o.key)
}

type AdvancedECBOracle struct {
	key    []byte
	prefix []byte
	secret []byte
}

func NewAdvancedECBOracle(secret []byte) Encrypter {
	rand.Seed(time.Now().UnixNano()) //seed our random generator
	pl := rand.Intn(32)

	prefix := make([]byte, pl)
	rand.Read(prefix)

	key, _ := GenerateKey()
	return &AdvancedECBOracle{
		key:    key,
		prefix: prefix,
		secret: secret,
	}
}

func (o *AdvancedECBOracle) Encrypt(userInput []byte) ([]byte, error) {
	b := append(o.prefix, userInput...)
	return ecb.Encrypt(append(b, o.secret...), o.key)
}
