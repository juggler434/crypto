package aes128

import (
	"bytes"
	"errors"
	"github.com/juggler434/crypto/aes128/oracle"
)

func BreakECBSimple(oracle oracle.Encrypter) ([]byte, error) {
	bs, tl, err := findBlockSize(oracle, []byte(""), 0)
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

	dictionarySeed := padding[(numberOfBlocks-1)*bs : (numberOfBlocks*bs)-1]

	res := make([]byte, 0)
	for i := 0; i < tl; i++ {
		cb := createInputDictionary(oracle, dictionarySeed, 0, 0)
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
	bs, _, err := findBlockSize(oracle, []byte(""), 0)
	if err != nil {
		return nil, err
	}
	pl := findPrefixLen(oracle, bs)

	prefixRemainder := bs - (pl % bs)

	align := makeAlign(prefixRemainder)
	noOfSecretBlocks := getNoOfSecretBlocks(oracle, align, pl, prefixRemainder, bs)

	padding := makePadding(bs, noOfSecretBlocks)
	prefixPad := pl + prefixRemainder

	dictionarySeed := append(align, padding...)

	_, etl, _ := findBlockSize(oracle, align, pl)
	res := make([]byte, 0)

	for i := 0; i < etl; i++ {
		cb := createInputDictionary(oracle, dictionarySeed, pl, prefixRemainder)
		op, _ := oracle.Encrypt(append(align, padding...))
		endIndex := prefixPad + (noOfSecretBlocks * bs)
		hc := op[prefixPad:endIndex]
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

func makePadding(bs int, noOfSecretBlocks int) []byte {
	padding := make([]byte, (bs*noOfSecretBlocks)-1)
	for i := 0; i < len(padding); i++ {
		padding[i] = 'A'
	}
	return padding
}

func getNoOfSecretBlocks(encrypter oracle.Encrypter, align []byte, pl int, prefixRemainder int, bs int) int {
	psl, _ := encrypter.Encrypt(align)
	sl := len(psl) - (pl + prefixRemainder)
	nob := sl / bs
	return nob
}

func makeAlign(prefixRemainder int) []byte {
	align := make([]byte, prefixRemainder)
	for i := 0; i < len(align); i++ {
		align[i] = 'A'
	}
	return align
}

func findBlockSize(encrypter oracle.Encrypter, buf []byte, padLength int) (int, int, error) {
	et, _ := encrypter.Encrypt(buf)
	for {
		buf = append(buf, byte('A'))
		ct, err := encrypter.Encrypt(buf)
		if err != nil {
			return -1, -1, err
		}
		if len(ct) > len(et) {
			blockSize := len(ct) - len(et)
			noOfBlocks := len(et) / blockSize
			cts := (noOfBlocks * blockSize) - len(buf) - padLength
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

func createInputDictionary(encrypter oracle.Encrypter, input []byte, prefixLength, padRemainder int) map[byte][]byte {
	id := make(map[byte][]byte)

	for i := 0; i < 256; i++ {
		ui := append(input, byte(i))
		res, _ := encrypter.Encrypt(ui)
		id[byte(i)] = res[prefixLength+padRemainder : prefixLength+len(input)+1]
	}

	return id
}

func findPrefixLen(encrypter oracle.Encrypter, blockSize int) int {
	baseEncrypted, _ := encrypter.Encrypt([]byte(""))

	var pl int
	ui := make([]byte, 0)
	ui = append(ui, byte('A'))
	e, _ := encrypter.Encrypt(ui)

	for i := 0; i < len(baseEncrypted); i += blockSize {
		if !bytes.Equal(e[i:i+blockSize], baseEncrypted[i:i+blockSize]) {
			eb := e[i : i+blockSize]
			for {
				ui = append(ui, 'A')
				ne, _ := encrypter.Encrypt(ui)
				if bytes.Equal(eb, ne[i:i+blockSize]) {
					pl = (i + blockSize) - (len(ui) - 1)
					return pl
				}
				eb = ne[i : i+blockSize]
			}
		}
	}
	return -1
}
