package cryptopals

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/juggler434/crypto/encoding/hex"
	"io/ioutil"
)

func DetectECBEncryption(file string) (int, error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return 0, err
	}
	res := 0
	scanner := bufio.NewScanner(bytes.NewReader(f))
	dups := 0
	var ct []byte

	ln := 0
	for scanner.Scan() {
		ln += 1
		l, err := hex.Decode(scanner.Bytes())
		if err != nil {
			return 0, err
		}
		dia := 0
		chunks := make([][]byte, 0)
		for i := 0; i < len(l); i += 16 {
			batch := l[i:min(i+15, len(l))]
			for _, c := range chunks {
				if bytes.Equal(c, batch) {
					dia += 1
					break
				}
			}
			chunks = append(chunks, batch)
		}
		if dia > dups {
			dups = dia
			res = ln
			ct = hex.Encode(l)
		}
	}
	fmt.Println(ct)
	fmt.Printf("Line Number: %d \n", res)
	return res, nil
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
