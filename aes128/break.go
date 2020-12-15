package aes128

import (
	"bytes"
	"errors"
	"github.com/juggler434/crypto/aes128/oracle"
)

func BreakECBSimple(oracle oracle.Encrypter) ([]byte, error) {
	bs, tl, err := findBlockSize(oracle)
	if err != nil {
		return nil, err
	}

	if !detectECB(oracle, bs) {
		return nil, errors.New("not ecb mode")
	}

	secretText, _ := oracle.Encrypt([]byte(""))
	numberOfBlocks := len(secretText) / bs

	padding := make([]byte, (bs*numberOfBlocks)-1)
	for i := 0; i < len(padding); i++ {
		padding[i] = 'A'
	}

	dictionarySeed := padding

	res := make([]byte, 0)
	for i := 0; i < tl; i++ {
		cb := createInputDictionary(oracle, dictionarySeed, bs)
		op, _ := oracle.Encrypt(padding)
		hc := op[(numberOfBlocks-1)*bs : (numberOfBlocks * bs)]

		for k, v := range cb {
			if bytes.Equal(v, hc) {
				res = append(res, k)
				dictionarySeed = dictionarySeed[1:]
				dictionarySeed = append(dictionarySeed, k)
				break
			}
		}
		padding = padding[1:]
	}

	return res, nil
}

func BreakECBAdvanced(oracle oracle.Encrypter) ([]byte, error) {
	return nil, nil
}

func findBlockSize(encrypter oracle.Encrypter) (int, int, error) {
	bs, _ := encrypter.Encrypt([]byte(""))
	buf := make([]byte, 0)
	for {
		buf = append(buf, byte('A'))
		ct, err := encrypter.Encrypt(buf)
		if err != nil {
			return -1, -1, err
		}
		if len(ct) > len(bs) {
			blockSize := len(ct) - len(bs)
			cts := blockSize - len(buf)
			return blockSize, cts, nil
		}
	}
}

func detectECB(encrypter oracle.Encrypter, blockSize int) bool {
	pad := make([]byte, blockSize*2)
	for i := 0; i < blockSize*2; i++ {
		pad[i] = 'A'
	}

	res, _ := encrypter.Encrypt(pad)
	return DetectMode(res) == ECB
}

func createInputDictionary(encrypter oracle.Encrypter, input []byte, blockSize int) map[byte][]byte {
	id := make(map[byte][]byte)

	for i := 0; i < 256; i++ {
		ui := append(input, byte(i))
		res, _ := encrypter.Encrypt(ui)
		id[byte(i)] = res[0:blockSize]
	}

	return id
}

//func findPrefixLen(encrypter oracle.Encrypter) int {
//	i1, _ := encrypter.Encrypt([]byte("A"))
//	return 0
//}
