package xor

//EncryptWithRepeatingKey XOR encrypts a given string with a given key.  Returns a hex encoded string.
func EncryptWithRepeatingKey(input, key []byte) []byte {
	eb := make([]byte, len(input))

	for i := 0; i < len(input); i++ {
		eb[i] = input[i] ^ key[i%len(key)]
	}
	return eb
}
