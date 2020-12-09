package aes128

import (
	"bytes"
	"github.com/juggler434/crypto/aes128/oracle"
)

func BreakECBSimple(secretText []byte) ([]byte, error) {
	o := oracle.NewECBOracle()

	bs, err := findBlockSize(o, secretText)
	if err != nil {
		return nil, err
	}

	if !detectECB(o, secretText, bs) { //This will always be true, just an exercise
		return nil, nil
	}

	pad := make([]byte, bs-1)
	for i := 0; i < len(pad); i++ {
		pad[i] = byte('A')
	}

	res := make([]byte, 0)

	for _, b := range secretText {
		cb := createInputDictionary(o, pad)
		op, _ := o.Encrypt(pad, []byte{b})
		for k, v := range cb {
			if bytes.Equal(v, op) {
				res = append(res, k)
				break
			}
		}
	}

	return res, nil
}

func findBlockSize(encrypter oracle.Encrypter, secret []byte) (int, error) {
	bs, _ := encrypter.Encrypt([]byte(""), secret)
	buf := make([]byte, 0)
	for {
		buf = append(buf, byte('A'))
		ct, err := encrypter.Encrypt(buf, secret)
		if err != nil {
			return -1, err
		}
		if len(ct) > len(bs) {
			return len(ct) - len(bs), nil
		}
	}
}

func detectECB(encrypter oracle.Encrypter, secretText []byte, blockSize int) bool {
	pad := make([]byte, blockSize*2)
	for i := 0; i < blockSize*2; i++ {
		pad[i] = 'A'
	}

	res, _ := encrypter.Encrypt(pad, secretText)
	return DetectMode(res) == ECB
}

func createInputDictionary(encrypter oracle.Encrypter, pad []byte) map[byte][]byte {
	cb := make(map[byte][]byte)
	for i := 0; i < 256; i++ {
		op, _ := encrypter.Encrypt(pad, []byte{byte(i)})
		cb[byte(i)] = op
	}
	return cb
}
