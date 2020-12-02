package padding

func PKCS7(input []byte, blockSize int) []byte {
	r := len(input) % blockSize
	var pl int
	if r != 0 {
		pl = blockSize - r // This makes it so we won't pad input with the correct length
	}
	for i := 0; i < pl; i++ {
		input = append(input, byte(pl))
	}

	return input
}
