package cryptopals

import (
	"github.com/juggler434/crypto/encoding/hex"
	"github.com/juggler434/crypto/xor"
	"io/ioutil"
	"strings"
)

func FindXorCipherString(dest string) ([]byte, error) {
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
		st, sc:= xor.SingleCharDecode(dl)

		if sc > score {
			res = st
			score = sc
		}
	}
	return res, nil
}