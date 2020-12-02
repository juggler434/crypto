package ecb

import (
	"bytes"
	"fmt"
)

func Detect(encryptedLines [][]byte) int {
	res := 0
	dups := 0
	var ct []byte

	for i, ln := range encryptedLines {
		dia := 0
		chunks := make([][]byte, 0)
		for j := 0; j < len(ln); j += 16 {
			batch := ln[j:min(j+15, len(ln))]
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
			res = i + 1
		}
	}
	fmt.Println(ct)
	fmt.Printf("Line Number: %d \n", res)
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
