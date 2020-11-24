package xor

func DetectSingleCharEncryption(input [][]byte) []byte {
	var res []byte
	var score int

	for _, l := range input {
		st, sc:= SingleCharDecode(l)

		if sc > score {
			res = st
			score = sc
		}
	}
	return res
}