package cryptopals

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"math"
)

func getHammingDistance(input1, input2 []byte) int {
	var d int

	for i := 0; i < len(input1); i ++ {
		for j := 0; j < 8; j++ {
			if input1[i] & (1<<uint(j)) != input2[i] & (1<<uint(j)) {
				d++
			}
		}
	}
	return d
}

func getKeyLength(input []byte) int {
	var kl int
	shd := math.MaxFloat64
	for len := 2; len < 22; len++ {
		var hd float64
		for i := 0; i < 4; i++ {
			si := len * i
			hd += float64(getHammingDistance(input[si:si+len], input[si+len:si+(len*2)]))
		}
		hd = hd / float64(len)
		if hd < shd {
			shd = hd
			kl = len
		}
	}
	return kl
}

func XorDecryptFile(file string) ([]byte, error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	ueb := make([]byte, base64.StdEncoding.DecodedLen(len(f)))
	_, err = base64.StdEncoding.Decode(ueb, f)
	if err != nil {
		return nil, err
	}

	ueb = bytes.Trim(ueb, "\x00")

	kl := getKeyLength(ueb)
	chunks := make([][]byte, kl)
	for i, c := range ueb {
		chunks[i%kl] = append(chunks[i%kl], c)
	}

	unencChunks := make([][]byte, kl)

	for i, c := range chunks {
		hb := make([]byte, hex.EncodedLen(len(c)))
		hex.Encode(hb, c)
		s, _, err := SingleXorCipher(hb)
		if err != nil {
			return nil, err
		}
		unencChunks[i] = s
	}

	r := make([]byte, 0)
	for i := 0; i < len(unencChunks[0]); i++ {
		for _, sl := range unencChunks {
			if i < len(sl) {
				r = append(r, sl[i])
			}
		}
	}
	return r, nil
}