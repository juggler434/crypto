package xor

import (
	"sort"
)

type keyLength struct {
	length   int
	strength float64
}

func getHammingDistance(input1, input2 []byte) int {
	var d int

	for i := 0; i < len(input1); i++ {
		for j := 0; j < 8; j++ {
			if input1[i]&(1<<uint(j)) != input2[i]&(1<<uint(j)) {
				d++
			}
		}
	}
	return d
}

func getKeyLengths(input []byte) []keyLength {
	kl := make([]keyLength, 0)
	for len := 2; len < 40; len++ {
		var hd float64
		for i := 0; i < 4; i++ {
			si := len * i
			hd += float64(getHammingDistance(input[si:si+len], input[si+len:si+(len*2)]))
		}
		hd = hd / float64(len)
		kl = append(kl, keyLength{strength: hd, length: len})
	}
	sort.Slice(kl, func(i, j int) bool { return kl[i].strength < kl[j].strength })
	return kl
}

func Decrypt(encryptedBytes []byte) ([]byte, error) {

	kls := getKeyLengths(encryptedBytes)

	var ret []byte
	var retstre int
	for kli := 0; kli < 5; kli++ {
		kl := kls[kli].length
		chunks := makeByteChunks(kl, encryptedBytes)

		unencChunks := applyXorCipher(kl, chunks)
		r := rebuildString(unencChunks)
		var score int
		for _, b := range r {
			score += GetCharWeight(b)
		}
		if score > retstre {
			ret = r
			retstre = score
		}

	}
	return ret, nil
}

func rebuildString(unencChunks [][]byte) []byte {
	r := make([]byte, 0)
	for i := 0; i < len(unencChunks[0]); i++ {
		for _, sl := range unencChunks {
			if i < len(sl) {
				r = append(r, sl[i])
			}
		}
	}
	return r
}

func applyXorCipher(kl int, chunks [][]byte) [][]byte {
	unencChunks := make([][]byte, kl)
	for i, c := range chunks {
		s, _ := SingleCharDecode(c)
		unencChunks[i] = s
	}
	return unencChunks
}

func makeByteChunks(kl int, ueb []byte) [][]byte {
	chunks := make([][]byte, kl)
	for i, c := range ueb {
		chunks[i%kl] = append(chunks[i%kl], c)
	}
	return chunks
}
