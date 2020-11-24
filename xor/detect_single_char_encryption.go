package xor

import (
	"github.com/juggler434/crypto/encoding/hex"
	"io/ioutil"
	"strings"
)

func DetectSingleCharEncryption(dest string) ([]byte, error) {
	f, err := ioutil.ReadFile(dest)
	if err != nil {
		return nil, err
	}

	lns := strings.Split(string(f), "\n")

	var res []byte
	var score int

	for _, l := range lns {
		dl, err := hex.Decode([]byte(l))
		if err != nil {
			return nil, err
		}
		st, sc:= SingleCharDecode(dl)

		if sc > score {
			res = st
			score = sc
		}
	}
	return res, nil
}