package cryptopals

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math"
)

func DetectECBEncryption(file string) ([]byte, error) {
	// 1. Read file line by line
	// 2. Hex decode line
	// 3. AES Decrypt the line
	// 4. Encrpyt line again
	// 5. Check to see if the encrypted line matches the unecrypted line
	// 6. If it does, return the unencrypted line
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	lhd := math.MaxInt32
	var res []byte

	scanner := bufio.NewScanner(bytes.NewReader(f))
	for scanner.Scan() {
		l, err := decodeHexBytes(scanner.Bytes())
		if err != nil {
			return nil, err
		}
		key := []byte("YELLOW SUBMARINE") // TODO MAKE THIS A CONSTANT BECAUSE IT AINT CHANGING
		ueb, err := decryptEcb(l, key)
		if err != nil {
			return nil, err
		}

		hd := getHammingDistance(ueb[:16], ueb[16:32])
		if hd < lhd {
			lhd = hd
			m := make([]byte, hex.EncodedLen(len(l)))
			hex.Encode(m, l)
			res = m
		}
	}
	fmt.Println(lhd)
	fmt.Printf("%s\n", res)
	return f, nil
}
