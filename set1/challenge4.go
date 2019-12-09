package cryptopals

import (
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
		st, sc, err := SingleXorCipher([]byte(l))
		if err != nil {
			return nil, err
		}
		if sc > score {
			res = st
			score = sc
		}
	}
	return res, nil
}