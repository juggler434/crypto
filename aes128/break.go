package aes128

import (
	"bytes"
	"errors"
	"github.com/juggler434/crypto/aes128/oracle"
)

func BreakECBSimple(oracle oracle.Encrypter) ([]byte, error) {
	bs, tl, err := findBlockSize(oracle, []byte(""))
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
		cb := createInputDictionary(oracle, dictionarySeed, bs, 0)
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
	bs, _, _ := findBlockSize(oracle, []byte(""))
	pl := findPrefixLen(oracle, bs)
	prefixRemainder := pl % bs
	prefixRemainder = bs - prefixRemainder

	align := make([]byte, prefixRemainder)
	for i := 0; i < len(align); i++ {
		align[i] = 'A'
	}

	psl, _ := oracle.Encrypt(align)
	secretLength := len(psl) - (pl + prefixRemainder)
	numberOfBlocks := secretLength / bs

	padding := make([]byte, (bs*numberOfBlocks)-1)
	for i := 0; i < len(padding); i++ {
		padding[i] = 'A'
	}

	dictionarySeed := append(align, padding...)

	_, etl, _ := findBlockSize(oracle, align)
	res := make([]byte, 0)

	for i := 0; i < etl; i++ {
		cb := createInputDictionary(oracle, dictionarySeed, bs, len(align)+pl)
		op, _ := oracle.Encrypt(append(align, padding...))
		prefixPad := pl + prefixRemainder
		startIndex := ((numberOfBlocks - 1) * bs) + prefixPad
		endIndex := (numberOfBlocks * bs) + prefixPad
		hc := op[startIndex:endIndex]
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

func findBlockSize(encrypter oracle.Encrypter, buf []byte) (int, int, error) {
	bs, _ := encrypter.Encrypt([]byte(""))
	bil := len(buf)
	for {
		buf = append(buf, byte('A'))
		ct, err := encrypter.Encrypt(buf)
		if err != nil {
			return -1, -1, err
		}
		if len(ct) > len(bs) {
			blockSize := len(ct) - len(bs)
			cts := blockSize - (len(buf) - bil)
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

func createInputDictionary(encrypter oracle.Encrypter, input []byte, blockSize, prefixLength int) map[byte][]byte {
	id := make(map[byte][]byte)

	for i := 0; i < 256; i++ {
		ui := append(input, byte(i))
		res, _ := encrypter.Encrypt(ui)
		id[byte(i)] = res[prefixLength : blockSize+prefixLength]
	}

	return id
}

func findPrefixLen(encrypter oracle.Encrypter, blockSize int) int {
	baseEncrypted, _ := encrypter.Encrypt([]byte(""))
	//numberOfBlocks := len(baseEncrypted) / blockSize

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

	// get base encryption with no input
	// if block 1 changes when adding a byte, then we know prefix is less than block size
	// if block 1 doesn't change
	return 10
}
