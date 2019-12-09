package cryptopals

//SingleXorCipher takes an encoded message and decodes it into a human readable string
func SingleXorCipher(codedMessage []byte) ([]byte, int, error) {
	b, err := decodeHexBytes(codedMessage)
	if err != nil {
		return nil, 0,  err
	}
	var answer []byte
	var score int
	for i := 0; i < 256; i++ {
		r := make([]byte, len(b))
		var s int
		for j := 0; j < len(b); j++ {
			c := b[j] ^ byte(i)
			s += getCharWeight(c)
			r[j] = c
		}
		if s > score {
			answer = r
			score = s
		}

		s = 0
	}
	return answer, score, nil
}

func getCharWeight(char byte) int {
	wm := map[byte]int{
		byte('U'): 1,
		byte('u'): 1,
		byte('L'): 2,
		byte('l'): 2,
		byte('D'): 3,
		byte('d'): 3,
		byte('R'): 4,
		byte('r'): 4,
		byte('H'): 5,
		byte('h'): 5,
		byte('S'): 6,
		byte('s'): 6,
		byte(' '): 7,
		byte('N'): 8,
		byte('n'): 8,
		byte('I'): 9,
		byte('i'): 9,
		byte('O'): 10,
		byte('o'): 10,
		byte('A'): 11,
		byte('a'): 11,
		byte('T'): 12,
		byte('t'): 12,
		byte('E'): 13,
		byte('e'): 13,
	}
	return wm[char]
}
